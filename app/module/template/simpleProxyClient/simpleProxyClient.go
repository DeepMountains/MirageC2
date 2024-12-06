package main

import (
	"flag"
	"log"
	"net"
	"os"
	"os/signal"
	"time"

	"github.com/armon/go-socks5"
)

var (
	server *socks5.Server
	logger *log.Logger
)

func init() {
	// 初始化日志
	logger = log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)
}

func main() {
	// 定义一个 flag 用于接收目标地址
	target := flag.String("target", "", "Target address to connect (e.g., 124.221.70.199:20001)")
	flag.Parse()

	if *target == "" {
		logger.Println("[+] Usage: program -target=<ip:port>")
		os.Exit(1)
	}

	// 检查 target 格式是否合法
	if _, _, err := net.SplitHostPort(*target); err != nil {
		logger.Fatalf("[-] Invalid target address: %v\n", err)
	}

	// 起一个简单的 socks5 服务
	var err error
	server, err = socks5.New(&socks5.Config{})
	if err != nil {
		logger.Fatalf("[-] Failed to start SOCKS5 server: %v\n", err)
	}

	// 捕获系统中断信号，优雅关闭
	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, os.Interrupt)
	go func() {
		<-stopChan
		logger.Println("[*] Received interrupt signal. Shutting down...")
		os.Exit(0)
	}()

	// 使用 5 个管道管理并发连接
	const maxPipes = 5
	pipes := make(chan struct{}, maxPipes)

	// 创建 worker 协程处理管道
	for i := 0; i < maxPipes; i++ {
		go connectionWorker(pipes, *target)
	}

	// 向管道中填充空闲标记，表示可以开始处理
	for i := 0; i < maxPipes; i++ {
		pipes <- struct{}{}
	}

	// 主线程阻塞，等待中断信号
	select {}
}

func connectionWorker(pipes chan struct{}, target string) {
	for {
		// 等待管道空闲信号
		<-pipes

		// 尝试建立连接
		conn, err := net.DialTimeout("tcp", target, 5*time.Second)
		if err != nil {
			logger.Printf("[-] Dial error: %v\n", err)
			// 放回信号，表示当前管道空闲
			pipes <- struct{}{}
			time.Sleep(2 * time.Second) // 避免频繁重试
			continue
		}

		logger.Printf("[+] Successfully connected to %s\n", target)

		// 处理连接
		go func() {
			handleSocks5(conn)
			// 放回信号，表示当前管道空闲
			pipes <- struct{}{}
		}()
	}
}

func handleSocks5(conn net.Conn) {
	defer conn.Close()

	// 取消默认的读写超时
	_ = conn.SetDeadline(time.Time{})

	// 使用该 socks5 库提供的 ServeConn 方法
	err := server.ServeConn(conn)
	if err != nil {
		logger.Printf("[-] Error handling connection: %v\n", err)
	}
}
