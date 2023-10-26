package service

import (
	"context"
	"mime/multipart"
	"strconv"
	"sync"

	"github.com/xilepeng/gin-mall/dao"
	"github.com/xilepeng/gin-mall/model"
	"github.com/xilepeng/gin-mall/pkg/e"
	util "github.com/xilepeng/gin-mall/pkg/utils"
	"github.com/xilepeng/gin-mall/serializer"
)

// 接口函数的实现

type ProductService struct {
	Id            uint   `json:"id" form:"id"`
	Name          string `json:"name" form:"name"`
	Num           int    `json:"num"`
	CategoryId    uint   `json:"category_id" form:"category_id"`
	Title         string `json:"title" form:"title"`
	Info          string `json:"info" form:"info"`
	ImgPath       string `json:"img_path" form:"img_path"`
	Price         string `json:"price" form:"price"`
	DiscountPrice string `json:"discount_price" form:"discount_price"`
	OnSale        bool   `json:"on_sale" form:"on_sale"`
	model.BasePage
}

func (service *ProductService) Create(ctx context.Context, uId uint, files []*multipart.FileHeader) serializer.Response {
	var boss *model.User
	var err error
	code := e.SUCCESS
	userDao := dao.NewUserDao(ctx)
	boss, _ = userDao.GetUserById(uId)
	// 以第1张作为封面图
	tmp, _ := files[0].Open()
	path, err := UploadProductToLocalStatic(tmp, uId, service.Name)
	if err != nil {
		code = e.ErrorNotExistProduct
		util.LogrusObj.Infoln(err)
		return serializer.Response{
			Status: code,
			Data:   e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	product := &model.Product{
		Name:          service.Name,
		CategoryId:    service.CategoryId,
		Title:         service.Title,
		Info:          service.Info,
		ImgPath:       path,
		Price:         service.Price,
		DiscountPrice: service.DiscountPrice, // 折后价
		OnSale:        true,
		Num:           service.Num,
		BossId:        uId,
		BossName:      boss.UserName,
		BossAvatar:    boss.Avatar,
	}
	productDao := dao.NewProductDao(ctx)
	err = productDao.CreateProduct(product)
	if err != nil {
		code = e.ERROR
		util.LogrusObj.Infoln(err)
		return serializer.Response{
			Status: code,
			Data:   e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	wg := new(sync.WaitGroup)
	wg.Add(len(files))
	for index, file := range files {
		num := strconv.Itoa(index)
		productImgDao := dao.NewProductImgDaoByDB(productDao.DB)
		tmp, _ = file.Open()
		path, err = UploadProductToLocalStatic(tmp, uId, service.Name+num)
		if err != nil {
			code = e.ErrorNotExistProduct
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
				Error:  err.Error(),
			}
		}
		productImg := model.ProductImg{
			ProductId: product.ID,
			ImgPath:   path,
		}
		err = productImgDao.CreateProductImg(&productImg)
		if err != nil {
			code = e.ERROR
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
				Error:  err.Error(),
			}
		}
		wg.Done()
	}
	wg.Wait()
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		// Error:  err.Error(), // panic 500
		Data: serializer.BuildProduct(product),
	}
}

// List 获取商品列表
func (service *ProductService) List(ctx context.Context) serializer.Response {
	var products []*model.Product
	var err error
	code := e.SUCCESS
	if service.PageSize == 0 {
		service.PageSize = 15
	}
	condition := make(map[string]interface{})
	if service.CategoryId != 0 {
		condition["category_id"] = service.CategoryId
	}
	productDao := dao.NewProductDao(ctx)
	total, err := productDao.CountProductByCondition(condition)
	if err != nil {
		code = e.ERROR
		util.LogrusObj.Infoln(err)
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func() {
		productDao = dao.NewProductDaoByDB(productDao.DB)
		products, _ = productDao.ListProductByCondition(condition, service.BasePage)
		wg.Done()
	}()
	wg.Wait()

	return serializer.BuildListResponse(serializer.BuildProducts(products), uint(total))
}

// Search 搜索商品
func (service *ProductService) Search(ctx context.Context) serializer.Response {
	code := e.SUCCESS
	if service.PageSize == 0 {
		service.PageSize = 15
	}
	productDao := dao.NewProductDao(ctx)
	products, count, err := productDao.SearchProduct(service.Info, service.BasePage)
	if err != nil {
		code = e.ERROR
		util.LogrusObj.Infoln(err)
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.BuildListResponse(serializer.BuildProducts(products), uint(count))
}

// Show 获取商品展示信息服务
func (service *ProductService) Show(ctx context.Context, id string) serializer.Response {
	code := e.SUCCESS
	pId, _ := strconv.Atoi(id)
	productDao := dao.NewProductDao(ctx)
	product, err := productDao.GetProductById(uint(pId))
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
		Data:   serializer.BuildProduct(product),
	}
}
