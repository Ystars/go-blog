package middleware

import (
	"github.com/gin-gonic/gin"
	"goblog/utils/enum"
	"goblog/utils/jwt"
	"goblog/utils/response"
	"strings"
)

// JwtToken jwt中间件
func JwtToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		tokenHeader := c.Request.Header.Get("Authorization")
		if tokenHeader == "" {
			code = enum.ERROR_TOKEN_EXIST
			response.AbortWithStatusJSON(enum.SUCCESS, response.R{}, code, c)
			return
		}

		checkToken := strings.Split(tokenHeader, " ")
		if len(checkToken) == 0 {
			response.AbortWithStatusJSON(enum.SUCCESS, response.R{}, code, c)
			return
		}

		if len(checkToken) != 2 || checkToken[0] != "Bearer" {
			response.AbortWithStatusJSON(enum.SUCCESS, response.R{}, code, c)
			return
		}

		j := jwt.NewJWT()
		// 解析token
		claims, err := j.ParserToken(checkToken[1])
		if err != nil {
			if err == jwt.TokenExpired {
				code = enum.ERROR_TOKEN_RUNTIME
				response.AbortWithStatusJSON(enum.SUCCESS, response.R{}, code, c)
				return
			}
			// 其他错误
			response.AbortWithStatusJSON(enum.SUCCESS, response.R{"message": err.Error()}, enum.ERROR, c)
			return
		}

		c.Set("username", claims.Username)
		c.Next()
	}
}
