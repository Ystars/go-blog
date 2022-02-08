package V1

import (
	"github.com/gin-gonic/gin"
	"goblog/enum"
	"goblog/service"
	"goblog/utils/response"
	"goblog/validate"
	"net/http"
)

type AdminApi struct{}

// Login 后台登录
func (t *AdminApi) Login(c *gin.Context) {
	var validate validate.AdminLogin

	_ = c.ShouldBindJSON(&validate)

	var adminService service.AdminService

	token, err := adminService.Login(validate.Username, validate.Password)

	if err != nil {
		response.FailWithError(err, c)
		return
	}
	response.OkWithData(response.R{"token": token}, c)

}

// Add 添加管理员
func (t *AdminApi) Add(c *gin.Context) {

	response.Ok(c)
}

// Delete 删除管理员
func (t *AdminApi) Delete(c *gin.Context) {

	response.Ok(c)
}

// Edit 修改管理员
func (t *AdminApi) Edit(c *gin.Context) {

	response.Ok(c)
}

// Detail 查看单个管理员
func (t *AdminApi) Detail(c *gin.Context) {
	user, _ := c.Get("username")
	c.JSON(http.StatusOK, gin.H{
		"status":  enum.SUCCESS,
		"aaa":     112,
		"message": user,
	})
}
