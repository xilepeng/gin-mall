package model

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name          string
	Category      uint
	Title         string
	Info          string
	ImgPath       string
	Price         string
	DiscountPrice string // 折后价
	OnSale        bool   `gorm:"default:false"`
	Num           int
	BossId        uint
	BossName      string
	BossAvatar    string
}
