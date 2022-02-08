package response

import (
	"github.com/gin-gonic/gin"
	"goblog/enum"
	"goblog/utils/maps"
)

type R map[string]interface{}

func Result(code int, data R, msgCode int, c *gin.Context) {
	// 开始时间
	data = maps.NewMaps().MergeStringInterface(R{
		"code":    msgCode,
		"message": enum.GetCodeMsg(msgCode),
	}, data)
	c.JSON(code, data)
}

func AbortWithStatusJSON(code int, data R, msgCode int, c *gin.Context) {
	Result(code, data, msgCode, c)
	c.Abort()
}

func Ok(c *gin.Context) {
	code := enum.SUCCESS
	Result(code, R{}, code, c)
}

func OkWithErrorMessage(msg error, c *gin.Context) {
	var transfer maps.Transfer
	code := enum.SUCCESS
	Result(code, R{
		"status":  enum.ERROR,
		"message": transfer.ErrorJsonToMap(msg),
	}, code, c)
}

func OkWithData(data R, c *gin.Context) {
	code := enum.SUCCESS
	Result(code, data, code, c)
}

func OkWithDetailed(data R, msgCode int, c *gin.Context) {
	Result(enum.SUCCESS, data, msgCode, c)
}

func Fail(c *gin.Context) {
	code := enum.ERROR
	Result(code, R{}, code, c)
}

func FailWithMessage(msgCode int, c *gin.Context) {
	Result(enum.ERROR, R{}, msgCode, c)
}

func FailWithError(err error, c *gin.Context) {
	code := enum.ERROR
	r := R{
		"message": err.Error(),
	}
	Result(enum.ERROR, r, code, c)
}

func FailWithDetailed(data R, msgCode int, c *gin.Context) {
	Result(enum.ERROR, data, msgCode, c)
}
