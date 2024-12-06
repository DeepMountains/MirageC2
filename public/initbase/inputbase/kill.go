package inputbase

import (
	"MirageC2_V1.0/public/check"
	"MirageC2_V1.0/public/common"
	"fmt"
	"strconv"
)

func Kill(parts []string) {
	if len(parts) != 3 {
		fmt.Println("[-] 您输入的Kill命令存在问题，请仔细检查输入的参数，或使用help命令查看帮助信息！")
		return
	}
	intid, _ := strconv.Atoi(parts[2])
	job := &common.Jobs[intid-1]
	if check.JobsExistsCheck(intid) == false {
		fmt.Println("[-] 输入的Job不存在")
		return
	}
	if parts[1] == "job" || parts[1] == "j" {
		job.Cancel()
		job.Status = false
		return
	}

	if parts[1] == "session" || parts[1] == "s" {
		command := "kill -9 " + job.Processid
		if job.Platform == "windows" {
			command = "taskkill /IM " + job.Processname + " /F"
		}
		job.Tasks[99998] = command
	} else {
		fmt.Println("[-] 您输入的Kill命令存在问题，请仔细检查输入的参数，或使用help命令查看帮助信息！")
	}
}
