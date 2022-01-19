package middleware

import (
	"github.com/gin-gonic/gin"
	"goblog/utils/enum"
	"goblog/utils/jwt"
	"net/http"
	"strings"
)

// JwtToken jwt中间件
func JwtToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		tokenHeader := c.Request.Header.Get("Authorization")
		if tokenHeader == "" {
			code = enum.ERROR_TOKEN_EXIST
			c.JSON(http.StatusOK, gin.H{
				"status":  code,
				"message": enum.GetCodeMsg(code),
			})
			c.Abort()
			return
		}

		checkToken := strings.Split(tokenHeader, " ")
		if len(checkToken) == 0 {
			c.AbortWithStatusJSON(http.StatusOK, gin.H{"status": code, "message": enum.GetCodeMsg(code), "data": nil})
			return
		}

		if len(checkToken) != 2 || checkToken[0] != "Bearer" {
			c.AbortWithStatusJSON(http.StatusOK, gin.H{"status": code, "message": enum.GetCodeMsg(code), "data": nil})
			return
		}

		j := jwt.NewJWT()
		// 解析token
		claims, err := j.ParserToken(checkToken[1])
		if err != nil {
			if err == jwt.TokenExpired {
				code = enum.ERROR_TOKEN_RUNTIME
				c.AbortWithStatusJSON(http.StatusOK, gin.H{"status": enum.ERROR, "message": enum.GetCodeMsg(code), "data": nil})
				return
			}
			// 其他错误
			c.AbortWithStatusJSON(http.StatusOK, gin.H{"status": enum.ERROR, "message": err.Error(), "data": nil})
			return
		}

		c.Set("username", claims.Username)
		c.Next()
	}
}
