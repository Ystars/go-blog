package adapter

import (
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"goblog/model"
)

// GormAdapter gorm适配器
func GormAdapter() (*gormadapter.Adapter, error) {
	return gormadapter.NewAdapterByDBWithCustomTable(model.Db, &model.CasbinRule{})
}
