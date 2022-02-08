package service

import (
	"fmt"
	"goblog/enum"
	"goblog/model"
	"goblog/utils/jwt"
	"goblog/utils/scrypt"
)

type AdminService struct{}

// Add 添加管理员
func (t *AdminService) Add() {}

// Delete 添加管理员
func (t *AdminService) Delete() {}

// Edit 添加管理员
func (t *AdminService) Edit() {}

// Detail 查看管理员
func (t *AdminService) Detail() {}

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
