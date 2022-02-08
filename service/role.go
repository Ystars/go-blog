package service

import (
	"github.com/casbin/casbin/v2"
	casbinUtils "goblog/utils/casbin"
	"goblog/validate"
)

type RoleService struct {
	*casbin.Enforcer
}

// Add 添加用户角色
func (t *RoleService) Add(roleValidate validate.RoleValidate) error {
	if err, ok := roleValidate.Check(&roleValidate); !ok {
		return err
	}
	_, err := NewAuthorize().AddRoleForUser(roleValidate.User, roleValidate.Role)
	if err != nil {
		return err
	}
	return nil
}

// Delete 删除用户角色
func (t *RoleService) Delete(roleValidate validate.RoleValidate) error {
	if err, ok := roleValidate.Check(&roleValidate); !ok {
		return err
	}
	_, err := NewAuthorize().DeleteRoleForUser(roleValidate.User, roleValidate.Role)
	if err != nil {
		return err
	}
	return nil
}

// NewRoleService 默认初始化
func NewRoleService() *RoleService {
	enforcer, _ := casbinUtils.NewConnect().GormConnect()
	return &RoleService{enforcer}
}
