package model

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type Admin struct {
	gorm.Model
	UserName       string
	PasswordDigest string
	Avatar         string `gorm:"size:1000"`
}

// 设置密码
func (Admin *Admin) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PassWordCost)
	if err != nil {
		return err
	}
	Admin.PasswordDigest = string(bytes)
	return nil
}

// 校验密码
func (Admin *Admin) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(Admin.PasswordDigest), []byte(password))
	return err == nil
}

// 封面地址
func (Admin *Admin) AvatarURL() string {
	signedGetURL := "http://qs4jac5zs.hn-bkt.clouddn.com/avatar1.png"
	return signedGetURL
}
