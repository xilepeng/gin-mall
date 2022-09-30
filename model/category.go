package model

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	CategoryName string 
}