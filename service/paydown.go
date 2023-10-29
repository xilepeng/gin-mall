package service

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/xilepeng/gin-mall/conf"
	"github.com/xilepeng/gin-mall/dao"
	"github.com/xilepeng/gin-mall/model"
	"github.com/xilepeng/gin-mall/pkg/e"
	util "github.com/xilepeng/gin-mall/pkg/utils"
	"github.com/xilepeng/gin-mall/serializer"
)

type OrderPay struct {
	OrderId   uint    `json:"order_id" form:"order_id"`
	Money     float64 `json:"money" form:"money"`
	OrderNo   string  `json:"order_no" form:"order_no"`
	ProductId uint    `json:"product_id" form:"product_id"`
	PayTime   string  `json:"pay_time" form:"pay_time"`
	Sign      string  `json:"sign" form:"sign"`
	BossId    uint    `json:"boss_id" form:"boss_id"`
	BossName  string  `json:"boss_name" form:"boss_name"`
	Num       int     `json:"num" form:"num"`
	Key       string  `json:"key" form:"key"` // Key 支付金额
}

func (service *OrderPay) OrderPay(ctx context.Context, uId uint) serializer.Response {
	util.Encrypt.SetKey(service.Key)
	code := e.SUCCESS
	orderDao := dao.NewOrderDao(ctx)
	tx := orderDao.Begin()
	order, err := orderDao.GetOrderById(service.OrderId, uId)
	if err != nil {
		util.LogrusObj.Infoln("err", err)
		code = e.ERROR
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	money := order.Money
	num := order.Num
	money = money * float64(num)

	// 用户扣钱
	userDao := dao.NewUserDao(ctx)
	user, err := userDao.GetUserById(uId)
	if err != nil {
		util.LogrusObj.Infoln("err", err)
		code = e.ERROR
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	// 对钱解密，减去订单，再加密保存
	moneyStr := util.Encrypt.AesDecoding(user.Money)
	moneyFloat, _ := strconv.ParseFloat(moneyStr, 64)
	// 金额不足进行回滚
	if moneyFloat-money < 0.0 {
		tx.Rollback()
		util.LogrusObj.Infoln(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  errors.New("金额不足").Error(),
		}
	}

	finMoney := fmt.Sprintf("%f", moneyFloat-money)
	user.Money = util.Encrypt.AesEncoding(finMoney)

	userDao = dao.NewUserDaoByDB(userDao.DB)
	// 更新用户金额失败，回滚
	err = userDao.UpdateUserById(uId, user)
	if err != nil {
		tx.Rollback()
		util.LogrusObj.Infoln(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	// 商家加钱
	var boss *model.User
	boss, err = userDao.GetUserById(service.BossId)

	moneyStr = util.Encrypt.AesDecoding(boss.Money) // 解密
	moneyFloat, _ = strconv.ParseFloat(moneyStr, 64)
	finMoney = fmt.Sprintf("%f", moneyFloat+money)
	boss.Money = util.Encrypt.AesEncoding(finMoney)

	// err = userDao.UpdateUserById(boss.ID, boss)
	err = userDao.UpdateUserById(service.BossId, boss)

	// 更新boss金额失败，回滚
	if err != nil {
		tx.Rollback()
		util.LogrusObj.Infoln(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	// 对应的商品数量 -1
	var product *model.Product
	productDao := dao.NewProductDao(ctx)
	product, err = productDao.GetProductById(service.ProductId)
	product.Num -= num
	if err != nil {
		tx.Rollback()
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	err = productDao.UpdateProduct(service.ProductId, product)
	if err != nil {
		tx.Rollback()
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	// 订单删除
	err = orderDao.DeleteOrderByOrderId(service.OrderId, uId)
	if err != nil {
		tx.Rollback()
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	// 自己的商品 + 1
	productUser := &model.Product{
		Name:          product.Name,
		CategoryId:    product.CategoryId,
		Title:         product.Title,
		Info:          product.Info,
		ImgPath:       conf.Host + conf.HttpPort + conf.ProductPath + product.ImgPath,
		Price:         product.Price,
		DiscountPrice: product.DiscountPrice,
		OnSale:        false,
		// Num:           1,  错误❌
		Num:        num,
		BossId:     uId,
		BossName:   user.UserName,
		BossAvatar: user.Avatar,
	}
	err = productDao.CreateProduct(productUser)
	if err != nil {
		tx.Rollback()
		util.LogrusObj.Infoln(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	tx.Commit()
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}
