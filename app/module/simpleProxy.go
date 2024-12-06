package module

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

var logger *log.Logger

func init() {
	logger = initLogger()
}

func initLogger() *log.Logger {
	logFile, err := os.OpenFile("proxy.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("[-] Could not open log file: %v\n", err)
	}
	return log.New(logFile, "", log.LstdFlags|log.Lshortfile)
}

func SocksStart(serverPort, forwardPort string) {
	userConnChan := make(chan net.Conn, 10)
	agentConnChan := make(chan net.Conn, 10)

	go ListenService(agentConnChan, "0.0.0.0:"+serverPort)
	go ListenService(userConnChan, "0.0.0.0:"+forwardPort)

	for agentConn := range agentConnChan {
		userConn := <-userConnChan

		// 打印配对成功信息
		logger.Printf("[+] Proxy connection established: User [%s] <-> Agent [%s]\n", userConn.RemoteAddr(), agentConn.RemoteAddr())

		go copyConn(userConn, agentConn)
	}
}

func ListenService(c chan net.Conn, ListenAddress string) {
	listener, err := net.Listen("tcp", ListenAddress)
	if err != nil {
		logger.Printf("[-] Listen Error: %v\n", err)
		return
	}
	defer listener.Close()

	fmt.Printf("[*] Start Listening on %s\n", ListenAddress)
	for {
		conn, err := listener.Accept()
		if err != nil {
			logger.Printf("[-] Accept Error: %v\n", err)
			continue
		}

		select {
		case c <- conn:
		default:
			logger.Println("[-] Connection queue is full.")
		}
	}
}

func copyConn(srcConn, dstConn net.Conn) {
	defer srcConn.Close()
	defer dstConn.Close()

	errChan := make(chan error, 2)

	// 双向数据传输
	go func() {
		_, err := io.Copy(dstConn, srcConn)
		errChan <- err
	}()
	go func() {
		_, err := io.Copy(srcConn, dstConn)
		errChan <- err
	}()

	// 只等待数据传输错误
	err := <-errChan
	if err != nil && err != io.EOF {
		logger.Printf("[-] Data transfer error: %v\n", err)
	}
}
