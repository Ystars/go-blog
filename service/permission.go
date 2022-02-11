package service

import (
	"github.com/casbin/casbin/v2"
	casbinUtils "goblog/utils/casbin"
	"goblog/validate"
)

type PermissionService struct {
	*casbin.Enforcer
}

// Add 添加角色权限
func (t *PermissionService) Add(permissionValidate validate.PermissionValidate) error {
	_, err := NewAuthorize().AddPermissionForUser(permissionValidate.User, permissionValidate.Rule, permissionValidate.Method)
	if err != nil {
		return err
	}
	return nil
}

// Delete 删除角色权限
func (t *PermissionService) Delete(permissionValidate validate.PermissionValidate) error {
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
