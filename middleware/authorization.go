package middleware

import (
	"github.com/gin-gonic/gin"
	"goblog/utils/casbin"
	"goblog/utils/enum"
)

func Authorization(obj string, act string) gin.HandlerFunc {
	return func(c *gin.Context) {

		val, existed := c.Get("username")
		if !existed {
			code := enum.ERROR_USER_MIDDLEWARE_AUTH
			c.AbortWithStatusJSON(code, gin.H{"message": enum.GetCodeMsg(code)})
			return
		}

		if obj == "" {
			obj = c.Request.URL.RequestURI()
		}

		if act == "" {
			act = c.Request.Method
		}

		enforcer, err := casbin.NewConnect().GormConnect()
		if err != nil {
			code := enum.ERROR_USER_AUTHORIZING
			c.AbortWithStatusJSON(code, gin.H{"message": enum.GetCodeMsg(code)})
			return
		}

		ok, _ := enforcer.Enforce(val, obj, act)
		if !ok {
			code := enum.FORBIDDEN
			c.AbortWithStatusJSON(code, gin.H{"message": enum.GetCodeMsg(code)})
			return
		}
		c.Next()
	}
}
