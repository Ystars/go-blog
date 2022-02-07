package V1

import (
	"github.com/gin-gonic/gin"
	"goblog/service"
	"goblog/utils/response"
)

type PermissionApi struct{}

// Add 添加角色权限
func (t *PermissionApi) Add(c *gin.Context) {
	permissionService := service.NewPermissionService()
	err := permissionService.Add(c)
	if err != nil {
		response.OkWithErrorMessage(err, c)
		return
	}
	response.Ok(c)
}

// Delete 删除角色权限
func (t *PermissionApi) Delete(c *gin.Context) {
	permissionService := service.NewPermissionService()
	err := permissionService.Delete(c)
	if err != nil {
		response.OkWithErrorMessage(err, c)
		return
	}
	response.Ok(c)
}
