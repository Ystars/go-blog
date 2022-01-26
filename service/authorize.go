package service

import (
	"github.com/casbin/casbin/v2"
	casbinUtils "goblog/utils/casbin"
)

type Authorize struct {
	enforcer *casbin.Enforcer
}

func NewAuthorize() *Authorize {
	enforcer, _ := casbinUtils.NewConnect().GormConnect()
	return &Authorize{enforcer}
}

// Enforce 校验
func (e *Authorize) Enforce(val ...interface{}) (bool, error) {
	return e.enforcer.Enforce(val...)
}

// AddPermissionForUser 为用户或角色添加权限
func (e *Authorize) AddPermissionForUser(user string, permission ...string) (bool, error) {
	return e.enforcer.AddPermissionForUser(user, permission...)
}
