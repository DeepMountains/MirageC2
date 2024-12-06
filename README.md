# MirageC2
一款基于golang开发的C2平台，在后续的开发中，将陆续完善其功能。

# 安装方式
1. 在服务端，您的PC需要具备[golang1.20+](https://go.dev)的运行环境，用于生成木马。
2. go环境安装完毕后，运行run.bat或者run.sh会自动帮您安装服务端所需的库环境。

# 使用方式
本程序中使用到的所有的模块，例如创建后门、创建ssh、创建socks等都在您的电脑上进行，所以第一次运行这些指令的时候需要花上一定的时间，请耐心等待。

1. 运行可执行程序打开MirageC2服务器端。
```
> help:	输出帮助信息。
> create: 	通过Create参数来生成指定的Shell。例子: create <windows/linux/darwin> <C2IP:C2Port> <Save Path>
            [!] 注意如果使用create功能，您必须保证Server端拥有golang的运行环境。
> kill:		杀死指定的job或者session。例子：kill job <id>
> listen:  	在本地设置侦听的端口。例子: listen 127.0.0.1:8888
> show: 	[jobs,sessions]查看侦听的端口以及能够命令控制的会话。例如：show sessions
> session:	进入到指定会话中进行命令执行。例子: session <id>
```

2. Session会话中支持的指令。
```
> help:        输出帮助信息。
> show:        查看当前session中命令执行的列表。
> upload:      上传文件至目标服务器，例子：upload <LocalPath> <RemotePath>
> download:    下载目标服务器文件至本地，例子：download <RemotePath> <LocalPath>
> socks:       1. 部署代理功能，例子：socks <ReverseServerIP> <ServerPort> <ForwardPort>
               2. 控制目标主机连接代理服务器，例子：socks connect <ReverseServerIP>:<ServerPort>
> ssh:         1. 部署SSH功能，例子：ssh install
               2. 连接目标主机的SSH：ssh connect <SSHServerIP>
```

3. 使用案例
```
Mirage Control > listen 127.0.0.1:8888
[+] Jobs Set: listen ip 127.0.0.1:8888

Mirage Control > create darwin 127.0.0.1:8888 mirage
[+] 已经成功输出后门程序至指定的目录！
[!] 如果您还未进行侦听，请运行：listen 127.0.0.1:8888

Mirage Control >  
[+] Sessions 1 Online: 127.0.0.1:52056 -> 127.0.0.1:8888

Mirage Control > show sessions
[*] Job ID: "1", Client IP: "127.0.0.1:52057", Shell User: "deepmountains", Platform: "darwin", ProcessName: "mirage"

Mirage Control > session 1

Mirage Control -> Session1 > whoami
[+] Command ID: 0 
[-+] Result: --------------- 
deepmountains
[-+] Result: ---------------
Mirage Control -> Session1 > 
```

# 免责声明
本工具仅供合法用途使用，旨在帮助用户测试和评估其自身系统的安全性，以发现潜在漏洞并采取适当的补救措施。使用本工具前，您需仔细阅读并同意以下免责声明：
1. 合法使用：本工具应仅在您拥有合法授权的系统或网络上使用。任何在未经授权的情况下使用本工具对他人系统进行测试的行为，均被视为违法，并可能导致刑事或民事责任。
2. 责任限制：使用本工具所产生的一切后果由用户自行承担。开发者不对因使用本工具造成的任何直接或间接损失负责，包括但不限于系统崩溃、数据丢失或服务中断。
3. 禁止恶意用途：严禁将本工具用于任何形式的恶意活动，包括但不限于黑客攻击、数据窃取、破坏系统或绕过安全措施等。
4. 知法守法：用户需确保其使用行为符合所在国家和地区的法律法规。若因用户的非法使用行为引发任何法律纠纷，开发者概不负责。
5. 技术支持：本工具为“按现状”提供，开发者不保证其适用于任何特定用途或场景，也不承诺提供技术支持或更新。
