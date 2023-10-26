package service

import (
	"context"

	"github.com/xilepeng/gin-mall/dao"
	"github.com/xilepeng/gin-mall/pkg/e"
	util "github.com/xilepeng/gin-mall/pkg/utils"
	"github.com/xilepeng/gin-mall/serializer"
)

type CategoryService struct {
}

func (service *CategoryService) List(ctx context.Context) serializer.Response {
	categoryDao := dao.NewCategoryDao(ctx)
	code := e.SUCCESS
	category, err := categoryDao.ListCategory()
	if err != nil {
		util.LogrusObj.Infoln("err", err)
		code = e.ERROR
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.BuildListResponse(serializer.BuildCategorys(category), uint(len(category)))
}
