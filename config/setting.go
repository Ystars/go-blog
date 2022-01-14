package config

import (
	"fmt"
	"gopkg.in/ini.v1"
)

var (
	AppMode  string
	HttpPort string
	JwtKey   string
)

// AutoLoad 全局自动加载配置
func AutoLoad() {
	file, err := ini.Load("config/config.ini")
	if err != nil {
		fmt.Println("配置文件读取错误，请检查文件路径:", err)
	}
	LoadServer(file)
	LoadDatabase(file)
}

// LoadServer 加载配置文件
func LoadServer(file *ini.File) {
	AppMode = file.Section("server").Key("AppMode").MustString("debug")
	HttpPort = file.Section("server").Key("HttpPort").MustString(":3030")
	JwtKey = file.Section("server").Key("JwtKey").MustString("JweRaYTHIK")
}
