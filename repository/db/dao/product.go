package dao

import (
	"context"

	"github.com/xilepeng/gin-mall/repository/db/model"
	"gorm.io/gorm"
)

//  dao层，对db进行操作

type ProductDao struct {
	*gorm.DB
}

func NewProductDao(ctx context.Context) *ProductDao {
	return &ProductDao{NewDBClient(ctx)}
}

func NewProductDaoByDB(db *gorm.DB) *ProductDao {
	return &ProductDao{db}
}

// CreateProduct 创建商品
func (dao *ProductDao) CreateProduct(product *model.Product) (err error) {
	return dao.DB.Model(&model.Product{}).Create(&product).Error
}

// CountProductByCondition 根据情况获取商品的数量
func (dao *ProductDao) CountProductByCondition(condition map[string]interface{}) (total int64, err error) {
	err = dao.DB.Model(&model.Product{}).Where(condition).Count(&total).Error
	return total, err
}

// ListProductByCondition 获取商品列表
func (dao *ProductDao) ListProductByCondition(condition map[string]interface{}, page model.BasePage) (products []*model.Product, err error) {
	err = dao.DB.Preload("Category").Where(condition).
		Offset((page.PageSize - 1) * (page.PageNum)).
		Limit(page.PageSize).Find(&products).Error
	return
}

// SearchProduct 搜索商品
func (dao *ProductDao) SearchProduct(info string, page model.BasePage) (products []*model.Product, count int64, err error) {
	err = dao.DB.Model(&model.Product{}).
		Where("title LIKE ? OR info LIKE ?", "%"+info+"%", "%"+info+"%").Count(&count).Error
	if err != nil {
		return
	}
	err = dao.DB.Model(&model.Product{}).
		Where("title LIKE ? OR info LIKE ?", "%"+info+"%", "%"+info+"%").
		Offset((page.PageSize - 1) * (page.PageNum)).
		Limit(page.PageSize).Find(&products).Error
	return
}

// GetProductById 通过 id 获取product
func (dao *ProductDao) GetProductById(id uint) (product *model.Product, err error) {
	err = dao.DB.Model(&model.Product{}).Where("id=?", id).
		First(&product).Error
	return
}

// UpdateProduct 更新商品
func (dao *ProductDao) UpdateProduct(id uint, product *model.Product) error {
	return dao.DB.Model(&model.Product{}).Where("id=?", id).Updates(product).Error
}

// DeleteProduct 删除商品
func (dao *ProductDao) DeleteProduct(pId uint) error {
	return dao.DB.Model(&model.Product{}).Delete(&model.Product{}).Error
}
