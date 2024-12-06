package inputbase

import (
	"MirageC2_V1.0/app/jobs"
	"fmt"
)

func Show(parts []string) {
	if len(parts) < 2 {
		fmt.Println("[-] 您输入的Show命令存在问题，请仔细检查输入的参数，或使用help命令查看帮助信息。")
		return
	}

	if len(parts) == 3 {
		if (parts[1] == "jobs" || parts[1] == "j") && parts[2] == "-a" {
			jobs.ShowAllJobs()
			return
		}
	}

	if parts[1] == "jobs" || parts[1] == "j" {
		jobs.ShowJobs()
		return
	}

	if parts[1] == "sessions" || parts[1] == "s" {
		jobs.ShowSessions()
		return
	}
	fmt.Println("[-] 您输入的Show命令存在问题，请仔细检查输入的参数，或使用help命令查看帮助信息。")
}
