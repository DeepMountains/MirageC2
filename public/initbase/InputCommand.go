package initbase

import (
	inputbase2 "MirageC2_V1.0/public/initbase/inputbase"
	"fmt"
	"github.com/chzyer/readline"
	"os"
	"regexp"
	"strings"
	"time"
)

func GetInput() {
	rl, err := readline.New("Mirage Control > ")
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
		handleCommand(strings.TrimSpace(line))
		time.Sleep(10 * time.Millisecond)
	}
}

func handleCommand(input string) {
	input = strings.TrimSpace(input)
	if input == "" {
		return
	}

	// 正则分割命令和参数
	re := regexp.MustCompile(`\s+`)
	parts := re.Split(input, -1)

	// 主命令解析逻辑
	switch parts[0] {
	case "exit":
		fmt.Println("[*] 退出控制台")
		os.Exit(0)
	case "help":
		inputbase2.PrintHelp()
	case "listen", "l":
		inputbase2.Listen(parts)
	case "show":
		inputbase2.Show(parts)
	case "create":
		inputbase2.Create(parts)
	case "session", "s":
		inputbase2.Sessions(parts)
	case "kill":
		inputbase2.Kill(parts)
	default:
		fmt.Println("[-] 未知命令。输入 'help' 查看可用命令。")
	}
}

/*
func sleep(parts []string) {
	if len(parts) != 3 {
		fmt.Println("[-] 输入的参数存在错误，例子: sleep 1 3")
		return
	}
	intID, _ := strconv.Atoi(parts[1])
	intTime, _ := strconv.Atoi(parts[2])
	jobs.SleepTimeChange(intID, intTime)
}
*/
