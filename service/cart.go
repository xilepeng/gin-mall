package service

import (
	"context"
	"strconv"

	"github.com/xilepeng/gin-mall/dao"
	"github.com/xilepeng/gin-mall/model"
	"github.com/xilepeng/gin-mall/pkg/e"
	util "github.com/xilepeng/gin-mall/pkg/utils"
	"github.com/xilepeng/gin-mall/serializer"
)

type CartService struct {
	Id        uint `json:"id" form:"id"`
	BossId    uint `json:"boss_id" form:"boss_id"`
	ProductId uint `json:"product_id" form:"product_id"`
	Num       int  `json:"num" form:"num"`
}

func (service *CartService) Create(ctx context.Context, uId uint) serializer.Response {
	var cart *model.Cart
	code := e.SUCCESS
	// 判断有没有这个商品
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
	// 创建购物车
	cartDao := dao.NewCartDao(ctx)
	cart = &model.Cart{
		UserId:    uId,
		ProductId: service.ProductId,
		BossId:    service.BossId,
		Num:       uint(service.Num),
	}
	err = cartDao.CreateCart(cart)
	if err != nil {
		util.LogrusObj.Infoln("err", err)
		code = e.ERROR
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	userDao := dao.NewUserDao(ctx)
	boss, err := userDao.GetUserById(service.BossId)
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
		Data:   serializer.BuildCart(cart, product, boss),
	}
}

// Show 购物车
func (service *CartService) Show(ctx context.Context, uId uint) serializer.Response {
	code := e.SUCCESS
	cartDao := dao.NewCartDao(ctx)
	carts, err := cartDao.ListCartByUserId(uId)
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
		Data:   serializer.BuildCarts(ctx, carts),
	}
}

// Update 修改购物车信息
func (service *CartService) Update(ctx context.Context, uId uint, cId string) serializer.Response {
	code := e.SUCCESS
	cartDao := dao.NewCartDao(ctx)
	cartId, _ := strconv.Atoi(cId)
	err := cartDao.UpdateCartNumById(uint(cartId), service.Num)
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

// Delete 删除购物车
func (service *CartService) Delete(ctx context.Context, uId uint, cId string) serializer.Response {
	cartId, _ := strconv.Atoi(cId)
	code := e.SUCCESS
	cartDao := dao.NewCartDao(ctx)
	err := cartDao.DeleteCartByCartId(uint(cartId), uId)
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
