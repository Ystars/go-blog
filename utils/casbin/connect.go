package casbin

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/persist"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"goblog/enum"
	"goblog/model"
)

type Connect struct{}

func NewConnect() *Connect {
	return &Connect{}
}

// Common casbin执行入口文件
func (e *Connect) Common(adapter persist.Adapter) (*casbin.Enforcer, error) {
	enforcer, err := casbin.NewEnforcer("config/rbac_model.conf", adapter)
	if err != nil {
		return nil, fmt.Errorf(enum.GetCodeMsg(enum.ERROR_CASBIN_LOCAL)+": %w", err)
	}
	// Load policies from DB dynamically
	err = enforcer.LoadPolicy()
	if err != nil {
		return nil, fmt.Errorf(enum.GetCodeMsg(enum.ERROR_CASBIN_DYNAMICALLY)+": %w", err)
	}
	return enforcer, nil
}

// GormConnect gorm执行入口文件
func (e *Connect) GormConnect() (*casbin.Enforcer, error) {
	var base model.Base
	adapter, _ := gormadapter.NewAdapterByDBWithCustomTable(base.GetDb(), &model.CasbinRule{})
	enforcer, err := e.Common(adapter)
	return enforcer, err
}
