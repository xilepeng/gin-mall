package serializer

import (
	"context"

	"github.com/xilepeng/gin-mall/conf"
	"github.com/xilepeng/gin-mall/repository/db/dao"
	"github.com/xilepeng/gin-mall/repository/db/model"
)

type Order struct {
	ID            uint   `json:"id"`
	OrderNum      uint64 `json:"order_num"`
	CreatedAt     int64  `json:"created_at"`
	UpdatedAt     int64  `json:"updated_at"`
	UserID        uint   `json:"user_id"`
	ProductID     uint   `json:"product_id"`
	BossID        uint   `json:"boss_id"`
	Num           int    `json:"num"`
	AddressName   string `json:"address_name"`
	AddressPhone  string `json:"address_phone"`
	Address       string `json:"address"`
	Type          uint   `json:"type"`
	ProductName   string `json:"product_name"`
	ImgPath       string `json:"img_path"`
	DiscountPrice string `json:"discount_price"`
	// Money        string `json:"money"`
}

func BuildOrder(order *model.Order, product *model.Product, address *model.Address) Order {
	return Order{
		ID:            order.ID,
		OrderNum:      order.OrderNum,
		CreatedAt:     order.CreatedAt.Unix(),
		UpdatedAt:     order.CreatedAt.Unix(),
		UserID:        order.UserID,
		ProductID:     order.ProductID,
		BossID:        order.BossID,
		Num:           order.Num,
		AddressName:   address.Address,
		AddressPhone:  address.Phone,
		Address:       address.Address,
		Type:          order.Type,
		ProductName:   product.Name,
		ImgPath:       conf.Host + conf.HttpPort + conf.ProductPath + product.ImgPath,
		DiscountPrice: product.DiscountPrice,
	}
}

func BuildOrders(ctx context.Context, items []*model.Order) (orders []Order) {
	productDao := dao.NewProductDao(ctx)
	addressDao := dao.NewAddressDao(ctx)
	for _, item := range items {
		product, err := productDao.GetProductById(item.ProductID)
		if err != nil {
			continue
		}
		address, err := addressDao.GetAddressByAid(item.AddressID)
		if err != nil {
			continue
		}
		order := BuildOrder(item, product, address)
		orders = append(orders, order)

	}
	return orders
}
