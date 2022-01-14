package routes

import (
	"github.com/gin-gonic/gin"
	"goblog/api/V1"
	"goblog/middleware"
)

func Load(r *gin.Engine) {

	r.Use(middleware.Log())
	r.Use(gin.Recovery())
	r.Use(middleware.Cors())

	auth := r.Group("admin")
	{
		admin := V1.AdminApi{}
		auth.POST("login", admin.Login)
	}

	auth.Use(middleware.JwtToken())
	{
		admin := V1.AdminApi{}
		// 用户模块的路由接口
		auth.GET("users", admin.Show)
	}

}
