package V1

import (
	"github.com/gin-gonic/gin"
	"goblog/service"
	"goblog/utils/response"
	"goblog/validate"
)

type AdminApi struct{}

// Login 后台登录
func (t *AdminApi) Login(c *gin.Context) {
	var validate validate.AdminCommon

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
	var validate validate.AdminCommon
	if err := validate.BindJsonValidate(&validate, c); err != nil {
		response.OkWithErrorMessage(err, c)
		return
	}

	var adminService service.AdminService

	add := adminService.Add(validate)

	if add != nil {
		response.FailWithError(add, c)
		return
	}
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
	/*var validate validate.AdminId

	_ = c.ShouldBindJSON(&validate)
	var adminService service.AdminService

	admin, err := adminService.Detail(validate.ID)

	if err != nil {
		response.FailWithError(err, c)
		return
	}
	response.OkWithData(response.R{"username": admin.Username, "id": admin.ID}, c)*/
}
