package service

import (
	"fmt"
	"goblog/enum"
	"goblog/model"
	"goblog/utils/jwt"
	"goblog/utils/scrypt"
	"goblog/validate"
)

type AdminService struct{}

// Add 添加管理员
func (t *AdminService) Add(data validate.AdminCommon) error {
	var adminModel model.Admin
	return adminModel.Add(data)
}

// Delete 删除管理员
func (t *AdminService) Delete(data model.Admin) error {
	return model.DB.Delete(&data).Error
}

// Edit 修改管理员
func (t *AdminService) Edit(data model.Admin) error {
	return model.DB.Save(&data).Error
}

// Detail 查看管理员
func (t *AdminService) Detail(id int) (model.Admin, error) {
	var adminModel model.Admin

	if err := model.DB.First(&adminModel, id).Error; err != nil {
		return adminModel, fmt.Errorf(enum.GetCodeMsg(enum.ERROR_USER_NOT_EXIST))
	}

	return adminModel, nil
}

//  setJwtToken 设置jwt登录信息
func (t *AdminService) setJwtToken(m model.Admin) (string, error) {
	return jwt.NewJWT().SetToken(jwt.SetTokenData{
		Username: m.Username,
		ID:       int(m.ID),
		Issuer:   "goblog",
	})
}

// Login 登录
func (t *AdminService) Login(username string, password string) (string, error) {

	var adminModel model.Admin

	if err := model.DB.Where("username = ?", username).First(&adminModel).Error; err != nil {
		return "", fmt.Errorf(enum.GetCodeMsg(enum.ERROR_USER_NOT_EXIST))
	}
	if scrypt.NewScrypt().ComparePassword(adminModel.Password, password) == false {
		return "", fmt.Errorf(enum.GetCodeMsg(enum.ERROR_PASSWORD_WRONG))
	}

	return t.setJwtToken(adminModel)
}
