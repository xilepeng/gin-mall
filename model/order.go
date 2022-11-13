package model

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	UserId    uint `gorm:"not null"`
	ProductId uint `gorm:"not null"`
	BossId    uint `gorm:"not null"`
	AddressId uint `gorm:"not null"`
	Num       int
	OrderNum  uint64  // 订单数
	Type      uint    // 订单类型：1 未支付 2 已支付
	Money     float64 // 这个订单多少钱
}
