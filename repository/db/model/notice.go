package model

import "gorm.io/gorm"

// 公告
type Notice struct {
	gorm.Model

	Text string `gorm:"type:text"`
}
