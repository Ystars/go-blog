package service

import (
	"github.com/casbin/casbin/v2"
	casbinUtils "goblog/utils/casbin"
)

type Authorize struct {
	*casbin.Enforcer
}

func NewAuthorize() *Authorize {
	enforcer, _ := casbinUtils.NewConnect().GormConnect()
	return &Authorize{enforcer}
}
