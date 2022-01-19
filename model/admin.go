package model

import (
	"goblog/utils/enum"
	"golang.org/x/crypto/bcrypt"
)

// Admin 后台用户
type Admin struct {
	Init
	Username string `gorm:"type:varchar(20);not null;comment:用户名" json:"username" validate:"required,min=4,max=12" label:"用户名"`
	Password string `gorm:"type:varchar(500);not null;comment:密码" json:"password" validate:"required,min=6,max=120" label:"密码"`
}

// CheckLogin 校验登录
func (t *Admin) CheckLogin(username string, password string) int {
	var PasswordErr error

	Db.Where("username = ?", username).First(&t)

	PasswordErr = bcrypt.CompareHashAndPassword([]byte(t.Password), []byte(password))

	if t.ID == 0 {
		return enum.ERROR_USER_NOT_EXIST
	}
	if PasswordErr != nil {
		return enum.ERROR_PASSWORD_WRONG
	}

	return enum.SUCCESS
}
