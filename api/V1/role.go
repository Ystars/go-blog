package V1

import (
	"github.com/gin-gonic/gin"
	"goblog/service"
	"goblog/utils/response"
)

type RoleApi struct{}

// Add 添加用户角色
func (t *RoleApi) Add(c *gin.Context) {
	roleService := service.NewRoleService()
	err := roleService.Add(c)
	if err != nil {
		response.OkWithErrorMessage(err, c)
		return
	}
	response.Ok(c)
}

// Delete 删除角色权限
func (t *RoleApi) Delete(c *gin.Context) {
	roleService := service.NewRoleService()
	err := roleService.Delete(c)
	if err != nil {
		response.OkWithErrorMessage(err, c)
		return
	}
	response.Ok(c)
}
