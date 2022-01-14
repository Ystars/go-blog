package main

import (
	"github.com/gin-gonic/gin"
	"goblog/config"
	"goblog/model"
	"goblog/routes"
)

func main() {
	config.AutoLoad()

	model.InitDb()

	gin.SetMode(config.AppMode)

	r := gin.New()

	routes.Load(r)

	_ = r.Run(config.HttpPort)
}
