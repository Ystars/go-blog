package model

import (
	"gorm.io/gorm"
)

type Base struct {
	gorm.Model
}

func (e *Base) GetDb() *gorm.DB {
	return db
}
