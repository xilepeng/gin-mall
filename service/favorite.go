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

type FavoriteService struct {
	ProductId  uint `json:"product_id" form:"product_id"`
	BossId     uint `json:"boss_id" form:"boss_id"`
	FavoriteId uint `json:"favorite_id" form:"favorite_id"`
	model.BasePage
}

// List 获取商品列表
func (service *FavoriteService) List(ctx context.Context, uId uint) serializer.Response {
	favoriteDao := dao.NewFavoriteDao(ctx)
	code := e.SUCCESS
	favorite, err := favoriteDao.ListFavorite(uId)
	if err != nil {
		code = e.ERROR
		util.LogrusObj.Infoln(err)
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	return serializer.BuildListResponse(serializer.BuildFavorites(ctx, favorite), uint(len(favorite)))
}

func (service *FavoriteService) Create(ctx context.Context, uId uint) serializer.Response {
	code := e.SUCCESS
	favoriteDao := dao.NewFavoriteDao(ctx)
	exist, _ := favoriteDao.FavoriteExistOrNot(service.ProductId, uId)
	if exist {
		code = e.ErrorExistFavorite
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	userDao := dao.NewUserDao(ctx)
	user, err := userDao.GetUserById(uId)
	if err != nil {
		code = e.ErrorUserNotFound
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	bossDao := dao.NewUserDao(ctx)
	boss, err := bossDao.GetUserById(service.BossId)
	if err != nil {
		code = e.ErrorUserNotFound
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	productDao := dao.NewProductDao(ctx)
	product, err := productDao.GetProductById(service.ProductId)
	if err != nil {
		code = e.ErrorNotExistProduct
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	favorite := &model.Favorite{
		User:      *user,
		UserId:    uId,
		Product:   *product,
		ProductId: service.ProductId,
		Boss:      *boss,
		BossId:    service.BossId,
	}

	err = favoriteDao.CreateFavorite(favorite)
	if err != nil {
		code = e.ErrorExistFavorite
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}

// Delete 获取商品列表
func (service *FavoriteService) Delete(ctx context.Context, uId uint, fId string) serializer.Response {
	favoriteDao := dao.NewFavoriteDao(ctx)
	favoriteId, _ := strconv.Atoi(fId)
	code := e.SUCCESS
	err := favoriteDao.DeleteFavorite(uId, uint(favoriteId))
	if err != nil {
		code = e.ERROR
		util.LogrusObj.Infoln(err)
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
