package dao

import (
	"context"

	"github.com/xilepeng/gin-mall/model"
	"gorm.io/gorm"
)

type CarouselDao struct {
	*gorm.DB
}

func NewCarouselDao(ctx context.Context) *CarouselDao {
	return &CarouselDao{NewDBClient(ctx)}
}

func NewCarouselDaoByDB(db *gorm.DB) *CarouselDao {
	return &CarouselDao{db}
}

// GetUserById 根据 id 获取 user
func (dao *CarouselDao) ListCarousel() (carousel []model.Carousel, err error) { // nitice 拼写错误
	err = dao.DB.Model(&model.Carousel{}).Find(&carousel).Error
	return
}

// CreateCarousel 创建Carousel
func (dao *CarouselDao) CreateCarousel(carousel *model.Carousel) error {
	return dao.DB.Model(&model.Carousel{}).Create(&carousel).Error
}
