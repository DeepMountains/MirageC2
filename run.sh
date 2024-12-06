#!/bin/bash

# 程序名称
program="MirageC2_linux_amd"

# 检查程序是否存在
if [ ! -f "$program" ]; then
    echo "[Error] Program \"$program\" not found. Exiting."
    exit 1
fi

# 检查 Go 是否安装
if ! command -v go &> /dev/null; then
    echo "[Error] Go is not installed. Please install Go from https://go.dev/ before proceeding."
    exit 1
fi

go env -w GOPROXY=https://goproxy.cn,direct
# 需要的 Go 库
libraries=(
    "github.com/armon/go-socks5"
    "github.com/chzyer/readline"
    "golang.org/x/net"
    "github.com/shirou/gopsutil"
    "golang.org/x/crypto"
    "golang.org/x/text"
)

# 检查并安装必要的库
for lib in "${libraries[@]}"; do
    echo "Checking library $lib..."
    if ! go list "$lib" &> /dev/null; then
        echo "Installing $lib..."
        if ! go get "$lib"; then
            echo "[Error] Failed to install $lib. Exiting."
            exit 1
        fi
    fi
done

# 运行主程序
echo "Starting application: $program"
chmod +x "$program"
./"$program"

