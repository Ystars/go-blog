package service

import (
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	casbinUtils "goblog/utils/casbin"
	"goblog/validate"
)

type PermissionService struct {
	*casbin.Enforcer
}

// Add 添加角色权限
func (t *PermissionService) Add(c *gin.Context) error {
	var permissionValidate validate.PermissionValidate
	if err, ok := permissionValidate.Check(c, &permissionValidate); !ok {
		return err
	}
	_, err := NewAuthorize().AddPermissionForUser(permissionValidate.User, permissionValidate.Rule, permissionValidate.Method)
	if err != nil {
		return err
	}
	return nil
}

// Delete 删除角色权限
func (t *PermissionService) Delete(c *gin.Context) error {
	var permissionValidate validate.PermissionValidate
	if err, ok := permissionValidate.Check(c, &permissionValidate); !ok {
		return err
	}
	_, err := NewAuthorize().DeletePermissionForUser(permissionValidate.User, permissionValidate.Rule, permissionValidate.Method)
	if err != nil {
		return err
	}
	return nil
}

// NewPermissionService 默认初始化
func NewPermissionService() *PermissionService {
	enforcer, _ := casbinUtils.NewConnect().GormConnect()
	return &PermissionService{enforcer}
}
