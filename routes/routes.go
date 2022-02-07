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
		// 后台casbin授权接口
		auth.Use(middleware.Authorization())
		{
			auth.GET("users", V1.Entry.Admin.Show)
			auth.POST("permission/add", V1.Entry.Permission.Add)
			auth.DELETE("permission/delete", V1.Entry.Permission.Delete)
			auth.POST("role/add", V1.Entry.Role.Add)
			auth.DELETE("role/delete", V1.Entry.Role.Delete)
		}
	}

	return r

}
