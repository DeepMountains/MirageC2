package sessionbase

import "fmt"

func Help() {
	helpstring := `
> help:        输出帮助信息。
> show:        查看当前session中命令执行的列表。
> upload:      上传文件至目标服务器，例子：upload <LocalPath> <RemotePath>
> download:    下载目标服务器文件至本地，例子：download <RemotePath> <LocalPath>
> socks:       1. 部署代理功能，例子：socks <ReverseServerIP> <ServerPort> <ForwardPort>
               2. 控制目标主机连接代理服务器，例子：socks connect <ReverseServerIP>:<ServerPort>
> ssh:         1. 部署SSH功能，例子：ssh install
               2. 连接目标主机的SSH：ssh connect <SSHServerIP>
`
	fmt.Println(helpstring)
}
