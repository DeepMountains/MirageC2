package common

import "os"

func SuppressOutput() func() {
	// 保存原始输出
	stdout := os.Stdout
	stderr := os.Stderr

	// 打开空设备
	nullFile, _ := os.OpenFile(os.DevNull, os.O_WRONLY, 0666)

	// 重定向输出到空设备
	os.Stdout = nullFile
	os.Stderr = nullFile

	// 返回一个恢复原始输出的函数
	return func() {
		os.Stdout = stdout
		os.Stderr = stderr
		nullFile.Close()
	}
}
