package service

import (
	"context"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/xilepeng/gin-mall/dao"
	"github.com/xilepeng/gin-mall/model"
	"github.com/xilepeng/gin-mall/pkg/e"
	util "github.com/xilepeng/gin-mall/pkg/utils"
	"github.com/xilepeng/gin-mall/serializer"
)

type OrderService struct {
	ProductId uint    `json:"product_id" form:"product_id"`
	Num       int     `json:"num" form:"num"`
	AddressId uint    `json:"address_id" form:"address_id"`
	Money     float64 `json:"money" form:"money"`
	BossId    uint    `json:"boss_id" form:"boss_id"`
	UserId    uint    `json:"user_id" form:"user_id"`
	OrderNum  uint    `json:"order_num" form:"order_num"`
	Type      int     `json:"type" form:"type"`
	model.BasePage
}

func (service *OrderService) Create(ctx context.Context, uId uint) serializer.Response {
	code := e.SUCCESS

	orderDao := dao.NewOrderDao(ctx)
	order := &model.Order{
		UserId:    uId,
		ProductId: service.ProductId,
		BossId:    service.BossId,
		Num:       service.Num,
		Money:     service.Money,
		Type:      1, // 默认未支付订单
	}
	// 检验地址是否存在
	addressDao := dao.NewAddressDao(ctx)
	address, err := addressDao.GetAddressByAid(service.AddressId)
	if err != nil {
		util.LogrusObj.Infoln("err", err)
		code = e.ERROR
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	order.AddressId = address.ID
	// 订单号创建，自动自动生成的随机 number + 唯一标识的 productId + 用户的 id
	number := fmt.Sprintf("09%v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(100000000)) // 	"math/rand"
	productNum := strconv.Itoa(int(service.ProductId))
	userNum := strconv.Itoa(int(service.UserId))
	number = number + productNum + userNum
	orderNum, _ := strconv.ParseUint(number, 10, 64)
	order.OrderNum = orderNum

	err = orderDao.CreateOrder(order)
	if err != nil {
		util.LogrusObj.Infoln("err", err)
		code = e.ERROR
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}

func (service *OrderService) Show(ctx context.Context, uId uint, oId string) serializer.Response {

	code := e.SUCCESS
	orderId, _ := strconv.Atoi(oId)
	orderDao := dao.NewOrderDao(ctx)
	order, err := orderDao.GetOrderById(uint(orderId), uId)
	if err != nil {
		util.LogrusObj.Infoln("err", err)
		code = e.ERROR
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	// 获取地址信息
	addressDao := dao.NewAddressDao(ctx)
	address, err := addressDao.GetAddressByAid(service.AddressId)
	if err != nil {
		util.LogrusObj.Infoln("err", err)
		code = e.ERROR
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	// 获取商品信息
	productDao := dao.NewProductDao(ctx)
	product, err := productDao.GetProductById(service.ProductId)
	if err != nil {
		util.LogrusObj.Infoln("err", err)
		code = e.ERROR
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildOrder(order, product, address),
	}
}

func (service *OrderService) List(ctx context.Context, uId uint) serializer.Response {
	var orders []*model.Order
	var total int64
	code := e.SUCCESS
	if service.PageSize == 0 {
		service.PageSize = 5
	}

	orderDao := dao.NewOrderDao(ctx)
	condition := make(map[string]interface{})
	condition["user_id"] = uId

	if service.Type == 0 {
		condition["type"] = 0
	} else {
		condition["type"] = service.Type
	}
	orders, total, err := orderDao.ListOrderByCondition(condition, service.BasePage)
	if err != nil {
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	return serializer.BuildListResponse(serializer.BuildOrders(ctx, orders), uint(total))
}

func (service *OrderService) Delete(ctx context.Context, uId uint, aId string) serializer.Response {
	OrderId, _ := strconv.Atoi(aId)
	code := e.SUCCESS
	OrderDao := dao.NewOrderDao(ctx)
	err := OrderDao.DeleteOrderByOrderId(uint(OrderId), uId)
	if err != nil {
		util.LogrusObj.Infoln("err", err)
		code = e.ERROR
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}
