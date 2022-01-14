package config

import (
	"gopkg.in/ini.v1"
)

var (
	DbType     string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassWord string
	DbName     string
)

// LoadDatabase 加载数据库配置
func LoadDatabase(file *ini.File) {
	DbType = file.Section("database").Key("DbType").MustString("mysql")
	DbHost = file.Section("database").Key("DbHost").MustString("localhost")
	DbPort = file.Section("database").Key("DbPort").MustString("3306")
	DbUser = file.Section("database").Key("DbUser").MustString("root")
	DbPassWord = file.Section("database").Key("DbPassWord").MustString("root")
	DbName = file.Section("database").Key("DbName").MustString("goblog")
}
