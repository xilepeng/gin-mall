package dao

import (
	"context"

	"github.com/xilepeng/gin-mall/model"
	"gorm.io/gorm"
)

type CartDao struct {
	*gorm.DB
}

func NewCartDao(ctx context.Context) *CartDao {
	return &CartDao{NewDBClient(ctx)}
}

func (dao *CartDao) CreateCart(in *model.Cart) error {
	return dao.DB.Model(&model.Cart{}).Create(&in).Error
}

func (dao *CartDao) GetCartByAid(aId uint) (cart *model.Cart, err error) {
	err = dao.DB.Model(&model.Cart{}).Where("id=?", aId).First(&cart).Error
	return
}

func (dao *CartDao) ListCartByUserId(uId uint) (cartes []*model.Cart, err error) {
	err = dao.DB.Model(&model.Cart{}).
		Where("user_id=?", uId).
		Find(&cartes).Error
	return
}

// func (dao *CartDao) UpdateCartById(cId uint, cart *model.Cart) error {
// 	return dao.DB.Model(&model.Cart{}).Where("id=?", cId).Updates(&cart).Error
// }

func (dao *CartDao) DeleteCartByCartId(cId, uId uint) error {
	return dao.DB.Model(&model.Cart{}).
		Where("id=? AND user_id=?", cId, uId).
		Delete(&model.Cart{}).Error
}

func (dao *CartDao) UpdateCartNumById(cId uint, num int) error {
	return dao.DB.Model(&model.Cart{}).
		Where("id=?", cId).
		Update("num", num).Error
}
