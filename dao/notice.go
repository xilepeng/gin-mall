package dao

import (
	"context"

	"github.com/xilepeng/gin-mall/model"
	"gorm.io/gorm"
)

type NoticeDao struct {
	*gorm.DB
}

func NewNoticeDao(ctx context.Context) *NoticeDao {
	return &NoticeDao{NewDBClient(ctx)}
}

func NewNoticeDaoByDB(db *gorm.DB) *NoticeDao {
	return &NoticeDao{db}
}

// GetNoticeById 根据 id 获取 notice
func (dao *NoticeDao) GetNoticeById(id uint) (notice *model.Notice, err error) { // nitice 拼写错误
	err = dao.DB.Model(&model.Notice{}).Where("id=?", id).First(&notice).Error
	return
}

// CreateNotice 创建notice
func (dao *NoticeDao) CreateNotice(notice *model.Notice) error {
	return dao.DB.Model(&model.Notice{}).Create(&notice).Error
}
