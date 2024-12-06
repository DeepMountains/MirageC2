#!/bin/bash

# 目标平台和架构组合
platforms=("darwin/arm64" "darwin/amd64" "linux/arm64" "linux/amd64" "windows/arm64" "windows/amd64")
output_names=("MirageC2_macOS_arm" "MirageC2_macOS_amd" "MirageC2_linux_arm" "MirageC2_linux_amd" "MirageC2_windows_arm" "MirageC2_windows_amd")

# 当前目录的必要文件和目录
config_file="config.json"
client_dir="MirageClient"
module_dir="app/module"
mod_file="go.mod"
sum_file="go.sum"
run_sh="run.sh"
run_bat="run.bat"

# 检查必需文件和目录是否存在
required_files=("$config_file" "$mod_file" "$sum_file" "$run_sh" "$run_bat")
required_dirs=("$client_dir" "$module_dir")

for file in "${required_files[@]}"; do
    if [[ ! -f "$file" ]]; then
        echo "[-] Error: $file 文件不存在！"
        exit 1
    fi
done

for dir in "${required_dirs[@]}"; do
    if [[ ! -d "$dir" ]]; then
        echo "[-] Error: $dir 目录不存在！"
        exit 1
    fi
done

# 循环构建和打包
for i in "${!platforms[@]}"; do
    platform="${platforms[$i]}"
    output_name="${output_names[$i]}"

    # 分割平台和架构
    OS="${platform%%/*}"       # 提取操作系统（例如 darwin）
    ARCH="${platform##*/}"     # 提取架构（例如 arm64）

    # 设置目标文件名
    output_file="${output_name}"

    # 针对 Windows 生成 .exe 后缀
    if [[ "$OS" == "windows" ]]; then
        output_file="${output_file}.exe"
    fi

    # 编译二进制文件
    echo "[*] 正在为 $platform 构建 $output_file ..."
    GOOS=$OS GOARCH=$ARCH go build -o "$output_file" main.go
    if [[ $? -ne 0 ]]; then
        echo "[-] Error: 构建 $output_file 失败！"
        continue
    fi

    # 创建打包目录
    package_dir="${output_name}_package"
    mkdir -p "$package_dir"

    # 将生成的文件、config.json、go.mod、go.sum 和 MirageClient 复制到打包目录
    cp "$output_file" "$package_dir/"
    cp "$config_file" "$package_dir/"
    cp "$mod_file" "$package_dir/"
    cp "$sum_file" "$package_dir/"
    cp -r "$client_dir" "$package_dir/"
    cp "$run_sh" "$package_dir/"       # 添加 run.sh
    cp "$run_bat" "$package_dir/"      # 添加 run.bat

    # 仅复制 app/module 目录到打包目录下
    mkdir -p "$package_dir/app"
    cp -r "$module_dir" "$package_dir/app/"

    # 打包为 tar.gz
    tarball_name="${output_name}.tar.gz"
    echo "[*] 正在打包 $tarball_name ..."
    tar -czf "$tarball_name" -C "$package_dir" .  # 使用 . 确保打包目录结构正确

    # 清理临时打包目录和生成的文件
    rm -rf "$package_dir"
    rm -f "$output_file"

    echo "[+] $tarball_name 生成成功！"
done

echo "[*] 所有任务完成！"
