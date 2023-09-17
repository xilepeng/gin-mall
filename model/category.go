package model

import "gorm.io/gorm"

// 商品的分类
type Category struct {
	gorm.Model

	CategoryName string // 分类名
}
