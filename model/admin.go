package model

import "gorm.io/gorm"

type Admin struct {
	gorm.Model
	UserName string
	PasswordDigest string
	Avatar string
}