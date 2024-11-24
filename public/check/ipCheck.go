package check

import (
	"regexp"
	"strconv"
)

// IPCheck 检查字符串是否符合 "IP:port" 格式
func IPCheck(input string) bool {
	// 正则表达式检查 IP 和端口格式
	re := regexp.MustCompile(`^((25[0-5]|2[0-4][0-9]|[0-1]?[0-9]{1,2})\.){3}(25[0-5]|2[0-4][0-9]|[0-1]?[0-9]{1,2}):(\d{1,5})$`)

	// 检查是否匹配并验证端口号范围（0-65535）
	if re.MatchString(input) {
		// 提取端口号并检查范围
		portStr := re.FindStringSubmatch(input)[4]
		port, err := strconv.Atoi(portStr)
		return err == nil && port >= 0 && port <= 65535
	}
	return false
}
