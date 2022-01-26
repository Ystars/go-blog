package response

import (
	"github.com/gin-gonic/gin"
	"goblog/utils/enum"
	"goblog/utils/maps"
)

type R map[string]interface{}

func Result(code int, data map[string]interface{}, msgCode int, c *gin.Context) {
	// 开始时间
	data = maps.NewMaps().MergeStringInterface(map[string]interface{}{
		"code":    msgCode,
		"message": enum.GetCodeMsg(msgCode),
	}, data)
	c.JSON(code, data)
}

func AbortWithStatusJSON(code int, data map[string]interface{}, msgCode int, c *gin.Context) {
	Result(code, data, msgCode, c)
	c.Abort()
}

func Ok(c *gin.Context) {
	code := enum.SUCCESS
	Result(code, map[string]interface{}{}, code, c)
}

func OkWithMessage(msgCode int, c *gin.Context) {
	Result(enum.SUCCESS, map[string]interface{}{}, msgCode, c)
}

func OkWithData(data map[string]interface{}, c *gin.Context) {
	code := enum.SUCCESS
	Result(code, data, code, c)
}

func OkWithDetailed(data map[string]interface{}, msgCode int, c *gin.Context) {
	Result(enum.SUCCESS, data, msgCode, c)
}

func Fail(c *gin.Context) {
	code := enum.ERROR
	Result(code, map[string]interface{}{}, code, c)
}

func FailWithMessage(msgCode int, c *gin.Context) {
	Result(enum.ERROR, map[string]interface{}{}, msgCode, c)
}

func FailWithDetailed(data map[string]interface{}, msgCode int, c *gin.Context) {
	Result(enum.ERROR, data, msgCode, c)
}
