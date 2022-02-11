package config

import (
	"fmt"
	_ "github.com/joho/godotenv/autoload"
	"goblog/model"
	"goblog/validate"
)

func Init() {
	// 加载国际化校验
	if err := validate.TransInit("zh"); err != nil {
		fmt.Printf("init trans failed, err:%v\n", err)
		return
	}

	// 连接数据库
	model.InitDb()
}
