package template

import (
	"os"
	"os/exec"
)

func CreateModule(app, platform, arch string) {
	switch app {
	case "socks":
		socksapp(platform, arch)
	case "ssh":
		sshapp(platform, arch)
	}

}

func socksapp(platform, arch string) {
	var cmd *exec.Cmd
	switch arch {
	case "amd64":
		if platform == "linux" {
			cmd = exec.Command("go", "build", "-o", "app/module/release/amd64_linux_proxyclient", "-ldflags=-s -w", "app/module/template/simpleProxyClient/simpleProxyClient.go")
			cmd.Env = append(os.Environ(), "GOOS=linux", "GOARCH=amd64")
		} else if platform == "darwin" { // 适配 macOS
			cmd = exec.Command("go", "build", "-o", "app/module/release/amd64_darwin_proxyclient", "-ldflags=-s -w", "app/module/template/simpleProxyClient/simpleProxyClient.go")
			cmd.Env = append(os.Environ(), "GOOS=darwin", "GOARCH=amd64")
		} else {
			cmd = exec.Command("go", "build", "-o", "app/module/release/amd64_windows_proxyclient.exe", "-ldflags", "-s -w -H=windowsgui", "app/module/template/simpleProxyClient/simpleProxyClient.go")
			cmd.Env = append(os.Environ(), "GOOS=windows", "GOARCH=amd64")
		}
	case "arm64":
		if platform == "linux" {
			cmd = exec.Command("go", "build", "-o", "app/module/release/arm64_linux_proxyclient", "-ldflags=-s -w", "app/module/template/simpleProxyClient/simpleProxyClient.go")
			cmd.Env = append(os.Environ(), "GOOS=linux", "GOARCH=arm64")
		} else if platform == "darwin" { // 适配 macOS ARM64
			cmd = exec.Command("go", "build", "-o", "app/module/release/arm64_darwin_proxyclient", "-ldflags=-s -w", "app/module/template/simpleProxyClient/simpleProxyClient.go")
			cmd.Env = append(os.Environ(), "GOOS=darwin", "GOARCH=arm64")
		} else {
			cmd = exec.Command("go", "build", "-o", "app/module/release/arm64_windows_proxyclient.exe", "-ldflags", "-s -w -H=windowsgui", "app/module/template/simpleProxyClient/simpleProxyClient.go")
			cmd.Env = append(os.Environ(), "GOOS=windows", "GOARCH=arm64")
		}
	}

	cmd.Run()
}

func sshapp(platform, arch string) {
	var cmd *exec.Cmd
	switch arch {
	case "amd64":
		if platform == "linux" {
			cmd = exec.Command("go", "build", "-o", "app/module/release/amd64_linux_ssh", "-ldflags=-s -w", "app/module/template/sshServer/sshServer.go")
			cmd.Env = append(os.Environ(), "GOOS=linux", "GOARCH=amd64")
		} else if platform == "darwin" { // 适配 macOS
			cmd = exec.Command("go", "build", "-o", "app/module/release/amd64_darwin_ssh", "-ldflags=-s -w", "app/module/template/sshServer/sshServer.go")
			cmd.Env = append(os.Environ(), "GOOS=darwin", "GOARCH=amd64")
		} else {
			cmd = exec.Command("go", "build", "-o", "app/module/release/amd64_windows_ssh.exe", "-ldflags", "-s -w -H=windowsgui", "app/module/template/sshServer/sshServer.go")
			cmd.Env = append(os.Environ(), "GOOS=windows", "GOARCH=amd64")
		}
	case "arm64":
		if platform == "linux" {
			cmd = exec.Command("go", "build", "-o", "app/module/release/arm64_linux_ssh", "-ldflags=-s -w", "app/module/template/sshServer/sshServer.go")
			cmd.Env = append(os.Environ(), "GOOS=linux", "GOARCH=arm64")
		} else if platform == "darwin" { // 适配 macOS ARM64
			cmd = exec.Command("go", "build", "-o", "app/module/release/arm64_darwin_ssh", "-ldflags=-s -w", "app/module/template/sshServer/sshServer.go")
			cmd.Env = append(os.Environ(), "GOOS=darwin", "GOARCH=arm64")
		} else {
			cmd = exec.Command("go", "build", "-o", "app/module/release/arm64_windows_ssh.exe", "-ldflags", "-s -w -H=windowsgui", "app/module/template/sshServer/sshServer.go")
			cmd.Env = append(os.Environ(), "GOOS=windows", "GOARCH=arm64")
		}
	}

	cmd.Run()
}
