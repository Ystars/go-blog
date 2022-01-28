package service

import (
	"github.com/casbin/casbin/v2"
	casbinUtils "goblog/utils/casbin"
)

type Role struct {
	enforcer *casbin.Enforcer
}

func NewRole() *Authorize {
	enforcer, _ := casbinUtils.NewConnect().GormConnect()
	return &Authorize{enforcer}
}

// AddRoleForUser 为用户添加角色
func (e *Authorize) AddRoleForUser(user string, role string, domain ...string) (bool, error) {
	return e.enforcer.AddRoleForUser(user, role, domain...)
}
