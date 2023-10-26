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

type AddressService struct {
	Name    string `json:"name" form:"name"`
	Phone   string `json:"phone" form:"phone"`
	Address string `json:"address" form:"address"`
}

func (service *AddressService) Create(ctx context.Context, uId uint) serializer.Response {
	var address *model.Address
	code := e.SUCCESS
	addressDao := dao.NewAddressDao(ctx)
	address = &model.Address{
		UserID:  uId,
		Name:    service.Name,
		Phone:   service.Phone,
		Address: service.Address,
	}
	err := addressDao.CreateAddress(address)
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

func (service *AddressService) Show(ctx context.Context, aId string) serializer.Response {
	addressId, _ := strconv.Atoi(aId)
	code := e.SUCCESS
	addressDao := dao.NewAddressDao(ctx)
	address, err := addressDao.GetAddressByAid(uint(addressId))
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
		Data:   serializer.BuildAddress(address),
	}
}

func (service *AddressService) List(ctx context.Context, uId uint) serializer.Response {
	code := e.SUCCESS
	addressDao := dao.NewAddressDao(ctx)
	addressList, err := addressDao.ListAddressByUserId(uId)
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
		Data:   serializer.BuildAddresses(addressList),
	}
}

func (service *AddressService) Update(ctx context.Context, uId uint, aId string) serializer.Response {
	var address *model.Address
	code := e.SUCCESS
	addressDao := dao.NewAddressDao(ctx)
	address = &model.Address{
		UserID:  uId,
		Name:    service.Name,
		Phone:   service.Phone,
		Address: service.Address,
	}
	addressId, _ := strconv.Atoi(aId)
	err := addressDao.UpdateAddressByUserId(uint(addressId), address)
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

func (service *AddressService) Delete(ctx context.Context, uId uint, aId string) serializer.Response {
	var address *model.Address
	addressId, _ := strconv.Atoi(aId)
	code := e.SUCCESS
	addressDao := dao.NewAddressDao(ctx)
	err := addressDao.DeleteAddressByAddressId(uint(addressId), uId)
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
		Data:   serializer.BuildAddress(address),
	}
}
