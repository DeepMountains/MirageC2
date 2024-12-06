package inputbase

import (
	createC2 "MirageC2_V1.0/app/create"
	"fmt"
)

func Create(parts []string) {
	if len(parts) != 4 {
		fmt.Println("[-] 您输入的Create命令存在问题，请仔细检查输入的参数，或使用help命令查看帮助信息！")
		return
	}
	err := createC2.ClientGenerate(parts[1], parts[2], parts[3])
	if err == nil {
		fmt.Println("[+] 已经成功输出后门程序至指定的目录！")
		fmt.Println("[!] 如果您还未进行侦听，请运行：listen", parts[2])
	} else {
		fmt.Println("[-] 后门程序生成失败！")
	}
}
