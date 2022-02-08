package V1

import (
	"github.com/gin-gonic/gin"
	"goblog/service"
	"goblog/utils/response"
	"goblog/validate"
)

type RoleApi struct{}

// Add 添加用户角色
func (t *RoleApi) Add(c *gin.Context) {
	var roleValidate validate.RoleValidate
	_ = c.ShouldBindJSON(&roleValidate)
	roleService := service.NewRoleService()
	err := roleService.Add(roleValidate)
	if err != nil {
		response.OkWithErrorMessage(err, c)
		return
	}
	response.Ok(c)
}

// Delete 删除角色权限
func (t *RoleApi) Delete(c *gin.Context) {
	var roleValidate validate.RoleValidate
	_ = c.ShouldBindJSON(&roleValidate)
	roleService := service.NewRoleService()
	err := roleService.Delete(roleValidate)
	if err != nil {
		response.OkWithErrorMessage(err, c)
		return
	}
	response.Ok(c)
}
