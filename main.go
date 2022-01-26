package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goblog/config"
	"goblog/model"
	"goblog/routes"
	"goblog/validate"
)

func main() {
	if err := validate.TransInit("zh"); err != nil {
		fmt.Printf("init trans failed, err:%v\n", err)
		return
	}

	config.AutoLoad()

	model.InitDb()

	gin.SetMode(config.AppMode)

	r := gin.New()

	routes.Load(r)

	_ = r.Run(config.HttpPort)
}
