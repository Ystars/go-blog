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

	admin := r.Group("/api/v1/admin")
	{
		admin.POST("login", V1.Entry.Admin.Login)

		admin.Use(middleware.JwtToken())
		{
			// 后台casbin授权接口
			admin.Use(middleware.Authorization())
			{
				admin.POST("add", V1.Entry.Admin.Add)
				admin.PUT("edit", V1.Entry.Admin.Edit)
				admin.DELETE("edit", V1.Entry.Admin.Delete)
				admin.GET("detail", V1.Entry.Admin.Detail)
				admin.POST("permission/add", V1.Entry.Permission.Add)
				admin.DELETE("permission/delete", V1.Entry.Permission.Delete)
				admin.POST("role/add", V1.Entry.Role.Add)
				admin.DELETE("role/delete", V1.Entry.Role.Delete)
			}
		}
	}

	return r

}
