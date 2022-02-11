package V1

import (
	"github.com/gin-gonic/gin"
	"goblog/service"
	"goblog/utils/response"
	"goblog/validate"
)

type PermissionApi struct{}

// Add 添加角色权限
func (t *PermissionApi) Add(c *gin.Context) {
	var permissionValidate validate.PermissionValidate
	if err := permissionValidate.BindJsonValidate(&permissionValidate, c); err != nil {
		response.OkWithErrorMessage(err, c)
		return
	}
	permissionService := service.NewPermissionService()
	err := permissionService.Add(permissionValidate)
	if err != nil {
		response.OkWithErrorMessage(err, c)
		return
	}
	response.Ok(c)
}

// Delete 删除角色权限
func (t *PermissionApi) Delete(c *gin.Context) {
	var permissionValidate validate.PermissionValidate
	_ = c.ShouldBindJSON(&permissionValidate)
	permissionService := service.NewPermissionService()
	err := permissionService.Delete(permissionValidate)
	if err != nil {
		response.OkWithErrorMessage(err, c)
		return
	}
	response.Ok(c)
}
