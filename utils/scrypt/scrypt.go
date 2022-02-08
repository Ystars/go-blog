package scrypt

import (
	"golang.org/x/crypto/bcrypt"
	"log"
)

type Scrypt struct{}

// GeneratePassword 生成密码
func (t *Scrypt) GeneratePassword(password string) string {

	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)

	if err != nil {
		log.Fatal(err)
	}

	return string(hash)
}

// ComparePassword 对比密码是否正确
func (t *Scrypt) ComparePassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

// NewScrypt 初始化
func NewScrypt() *Scrypt {
	return &Scrypt{}
}
