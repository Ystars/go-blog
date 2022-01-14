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
				"message": enum.GetErrMsg(code),
			})
			c.Abort()
			return
		}

		checkToken := strings.Split(tokenHeader, " ")
		if len(checkToken) == 0 {
			c.JSON(http.StatusOK, gin.H{
				"status":  code,
				"message": enum.GetErrMsg(code),
			})
			c.Abort()
			return
		}

		if len(checkToken) != 2 || checkToken[0] != "Bearer" {
			c.JSON(http.StatusOK, gin.H{
				"status":  code,
				"message": enum.GetErrMsg(code),
			})
			c.Abort()
			return
		}

		j := jwt.NewJWT()
		// 解析token
		claims, err := j.ParserToken(checkToken[1])
		if err != nil {
			if err == jwt.TokenExpired {
				code = enum.ERROR_TOKEN_RUNTIME
				c.JSON(http.StatusOK, gin.H{
					"status":  enum.ERROR,
					"message": enum.GetErrMsg(code),
					"data":    nil,
				})
				c.Abort()
				return
			}
			// 其他错误
			c.JSON(http.StatusOK, gin.H{
				"status":  enum.ERROR,
				"message": err.Error(),
				"data":    nil,
			})
			c.Abort()
			return
		}

		c.Set("username", claims)
		c.Next()
	}
}
