package service

import (
	"context"

	"github.com/xilepeng/gin-mall/dao"
	"github.com/xilepeng/gin-mall/pkg/e"
	"github.com/xilepeng/gin-mall/serializer"
)

type CarouselService struct {
}

func (service *CarouselService) List(ctx context.Context) serializer.Response {
	carouselDao := dao.NewCarouselDao(ctx)
	code := e.SUCCESS
	carousels, err := carouselDao.ListCarousel()
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.BuildListResponse(serializer.BuildCarousels(carousels), uint(len(carousels)))
}
