package jwt

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"goblog/config"
	"goblog/utils/enum"
	"net/http"
	"time"
)

// 定义错误
var (
	TokenExpired     error = errors.New(enum.GetErrMsg(enum.ERROR_TOKEN_RUNTIME))
	TokenNotValidYet error = errors.New(enum.GetErrMsg(enum.ERROR_TOKEN_RUNTIME))
	TokenMalformed   error = errors.New(enum.GetErrMsg(enum.ERROR_TOKEN_WRONG))
	TokenInvalid     error = errors.New(enum.GetErrMsg(enum.ERROR_TOKEN_TYPE_WRONG))
)

type SetJwtInfo struct {
	Username string
	ID       int
	Issuer   string
}

type JWT struct {
	JwtKey []byte
}

type MyClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// CreateToken 生成token
func (j *JWT) CreateToken(claims MyClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.JwtKey)
}

// ParserToken 解析token
func (j *JWT) ParserToken(tokenString string) (*MyClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.JwtKey, nil
	})

	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}

	if token != nil {
		if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, TokenInvalid
	}

	return nil, TokenInvalid
}

// SetToken 设置token
func (j *JWT) SetToken(c *gin.Context, info SetJwtInfo) {
	claims := MyClaims{
		Username: info.Username,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 100,
			ExpiresAt: time.Now().Unix() + 604800,
			Issuer:    info.Issuer,
		},
	}

	token, err := j.CreateToken(claims)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  enum.ERROR,
			"message": enum.GetErrMsg(enum.ERROR),
			"token":   token,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"data":    info.Username,
		"id":      info.ID,
		"message": enum.GetErrMsg(200),
		"token":   token,
	})
	return
}

// NewJWT 初始化
func NewJWT() *JWT {
	return &JWT{
		[]byte(config.JwtKey),
	}
}
