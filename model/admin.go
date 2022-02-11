package model

// Admin 后台用户
type Admin struct {
	Base
	Username string `gorm:"type:varchar(20);not null;comment:用户名"`
	Password string `gorm:"type:varchar(500);not null;comment:密码"`
}

// Add 添加管理员
func (t *Admin) Add(data interface{}) error {
	return DB.Model(&t).Create(data).Error
}

// Delete 删除管理员
func (t *Admin) Delete() error {
	return DB.Delete(&t).Error
}

// Edit 修改管理员
func (t *Admin) Edit() error {
	return DB.Save(&t).Error
}
