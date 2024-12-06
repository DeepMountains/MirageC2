package inputbase

import (
	"MirageC2_V1.0/app/jobs"
	"MirageC2_V1.0/public/check"
	"fmt"
)

func Listen(parts []string) {
	if len(parts) != 2 {
		fmt.Println("[-] 输入的参数存在错误，例子: listen 127.0.0.1:8888")
		return
	}
	if check.IPCheck(parts[1]) {
		jobs.CreateJobs(parts[1])
	} else {
		fmt.Println("[-] 请输入正确的IP地址！例子: listen 127.0.0.1:8888")
	}
	return
}
