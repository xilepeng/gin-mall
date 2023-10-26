package dao

import (
	"context"

	"github.com/xilepeng/gin-mall/model"
	"gorm.io/gorm"
)

type OrderDao struct {
	*gorm.DB
}

func NewOrderDao(ctx context.Context) *OrderDao {
	return &OrderDao{NewDBClient(ctx)}
}

func (dao *OrderDao) CreateOrder(in *model.Order) error {
	return dao.DB.Model(&model.Order{}).Create(&in).Error
}

func (dao *OrderDao) GetOrderById(id, uId uint) (order *model.Order, err error) {
	err = dao.DB.Model(&model.Order{}).Where("id=? AND user_id=?", id, uId).First(&order).Error
	return
}

func (dao *OrderDao) ListOrderByUserId(uId uint) (orderes []*model.Order, err error) {
	err = dao.DB.Model(&model.Order{}).Where(" user_id=?", uId).Find(&orderes).Error
	return
}

func (dao *OrderDao) UpdateOrderByUserId(aId uint, order *model.Order) error {
	return dao.DB.Model(&model.Order{}).Where("id=?", aId).Updates(&order).Error
}

func (dao *OrderDao) DeleteOrderByOrderId(aId, uId uint) error {
	return dao.DB.Model(&model.Order{}).Where("id=? AND user_id=?", aId, uId).Delete(&model.Order{}).Error
}

// ListOrderByCondition 获取订单List
func (dao *OrderDao) ListOrderByCondition(condition map[string]interface{}, page model.BasePage) (orders []*model.Order, total int64, err error) {
	err = dao.DB.Model(&model.Order{}).Where(condition).
		Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = dao.DB.Model(&model.Order{}).Where(condition).
		Offset((page.PageNum - 1) * page.PageSize).
		Limit(page.PageSize).Order("created_at desc").Find(&orders).Error
	return
}
