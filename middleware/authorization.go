package middleware

import (
	"github.com/gin-gonic/gin"
	"goblog/service"
	"goblog/utils/enum"
)

func Authorization() gin.HandlerFunc {
	return func(c *gin.Context) {

		val, existed := c.Get("username")
		if !existed {
			code := enum.ERROR_USER_MIDDLEWARE_AUTH
			c.AbortWithStatusJSON(code, gin.H{"message": enum.GetCodeMsg(code)})
			return
		}

		ok, _ := service.NewAuthorize().Enforce(val, c.Request.URL.RequestURI(), c.Request.Method)
		if !ok {
			code := enum.FORBIDDEN
			c.AbortWithStatusJSON(code, gin.H{"message": enum.GetCodeMsg(code)})
			return
		}
		c.Next()
	}
}
