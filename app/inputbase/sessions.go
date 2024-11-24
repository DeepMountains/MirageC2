package inputbase

import (
	"MirageC2_V1.0/app/jobs"
	"MirageC2_V1.0/public/check"
	"fmt"
	"github.com/chzyer/readline"
	"os"
	"strconv"
	"time"
)

func Sessions(parts []string) {
	if len(parts) != 2 {
		fmt.Println("[-] 您输入的Session命令存在问题，请仔细检查输入的参数，或使用help命令查看帮助信息！")
		return
	}
	intID, err := strconv.Atoi(parts[1])
	if check.SessionsExistsCheck(intID) == false || err != nil {
		fmt.Println("[-] 输入的Sessions不存在，或者并不处于Active状态。")
		return
	}
	rl, err := readline.New("Mirage Control -> Session" + parts[1] + " > ")
	if err != nil {
		fmt.Println("[-] 读取输入出错:", err)
		os.Exit(1)
	}
	defer rl.Close()
	for {
		line, err := rl.Readline()
		if err != nil {
			break // 处理 EOF 或其他读取错误
		}
		if line == "exit" {
			fmt.Printf("[*] Sessions %d 将被挂起\n", intID)
			os.Stdout.Sync()
			break
		}
		if line == "show" {
			jobs.ShowTasksPool(intID)
			continue
		}
		if line == "" {
			continue
		}
		jobs.TaskPoolAdd(intID, line)
		time.Sleep(10 * time.Millisecond)
	}
}
