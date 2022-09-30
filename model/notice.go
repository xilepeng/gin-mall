package model

import "gorm.io/gorm"

type Notice struct {
	gorm.Model
	Text string `gorm:"type:text"`
}
