package routes

import (
	"github.com/gin-gonic/gin"
	"goblog/api/V1"
	"goblog/middleware"
)

func Load() *gin.Engine {

	r := gin.New()

	r.Use(middleware.Log())
	r.Use(gin.Recovery())
	r.Use(middleware.Cors())

	auth := r.Group("admin")
	{
		auth.POST("login", V1.Entry.Admin.Login)
	}

	auth.Use(middleware.JwtToken())
	{
		auth.Use(middleware.Authorization())
		{
			// 用户模块的路由接口
			auth.GET("users", V1.Entry.Admin.Show)
			auth.GET("permission/add", V1.Entry.Permission.Add)
		}
	}

	return r

}
