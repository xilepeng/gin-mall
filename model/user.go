package model

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	UserName       string `gorm:"unique"`
	Email          string
	PasswordDigest string
	NiceName       string
	Status         string
	Avatar         string
	Money          string
}

const (
	PasswordCost        = 12       // 密码加密难度
	Active       string = "active" // 激活用户
)

// SetPassword 密码加密
func (user *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PasswordCost)
	if err != nil {
		return err
	}
	user.PasswordDigest = string(bytes)
	return nil
}

func (user *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordDigest), []byte(password))
	return err == nil
}
