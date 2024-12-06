@echo off

:: 设置窗口大小
set rows=50
set cols=200
mode con: cols=%cols% lines=%rows%

:: 设置程序名
set "program=MirageC2_windows_amd.exe"

:: 检查程序是否存在
if not exist "%program%" (
    echo [Error] Program "%program%" not found. Exiting.
    exit /b 1
)

:: 检查 Go 是否安装
go version >nul 2>&1
if errorlevel 1 (
    echo [Error] Go is not installed. Please install Go from https://go.dev/ before proceeding.
    exit /b 1
)

:: 检查并安装所需的 Go 库
go env -w GOPROXY=https://goproxy.cn,direct
for %%G in (
    "github.com/armon/go-socks5"
    "github.com/chzyer/readline"
    "golang.org/x/net"
    "github.com/shirou/gopsutil"
    "golang.org/x/crypto"
    "golang.org/x/text"
) do (
    echo Checking library %%G...
    go list %%G >nul 2>&1
    if errorlevel 1 (
        echo Installing %%G...
        go get %%G
        if errorlevel 1 (
            echo [Error] Failed to install %%G. Exiting.
            exit /b 1
        )
    )
)

:: 运行应用程序
echo Starting application: %program%
"%program%"
