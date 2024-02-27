package model

import "gorm.io/gorm"

type Cart struct {
	gorm.Model

	UserID    uint `gorm:"not null"`
	ProductID uint `gorm:"not null"`
	BossID    uint `gorm:"not null"`
	Num       uint `gorm:"not null"`
	MaxNum    uint `gorm:"not null"`
	Check     bool // 是否支付
}
