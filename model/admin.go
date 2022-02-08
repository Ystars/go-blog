package model

// Admin 后台用户
type Admin struct {
	Base
	Username string `gorm:"type:varchar(20);not null;comment:用户名"`
	Password string `gorm:"type:varchar(500);not null;comment:密码"`
}
