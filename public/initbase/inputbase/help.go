package inputbase

import "fmt"

func PrintHelp() {
	helpstring := `
> help:		输出帮助信息。
> create: 	通过Create参数来生成指定的Shell。例子: create <windows/linux/darwin> <C2IP:C2Port> <Save Path>
            [!] 注意如果使用create功能，您必须保证Server端拥有golang的运行环境。
> kill:		杀死指定的job或者session。例子：kill job <id>
> listen:  	在本地设置侦听的端口。例子: listen 127.0.0.1:8888
> show: 	[jobs,sessions]查看侦听的端口以及能够命令控制的会话。例如：show sessions
> session:	进入到指定会话中进行命令执行。例子: session <id>
`
	fmt.Println(helpstring)
}
