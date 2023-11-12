package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xilepeng/gin-mall/consts"
	util "github.com/xilepeng/gin-mall/pkg/utils"
	"github.com/xilepeng/gin-mall/service"
)

// CreateProduct 创建商品
func CreateProduct(c *gin.Context) {
	form, _ := c.MultipartForm()
	files := form.File["file"]
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	createProductService := service.ProductService{}
	if err := c.ShouldBind(&createProductService); err == nil {
		res := createProductService.Create(c.Request.Context(), claim.ID, files)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln(err)
	}
}

// ShowProduct 商品详情
func ShowProduct(c *gin.Context) {
	showProductService := service.ProductService{}
	res := showProductService.Show(c.Request.Context(), c.Param("id"))
	c.JSON(http.StatusOK, res)
}

// 删除商品
func DeleteProduct(c *gin.Context) {
	DeleteProductService := service.ProductService{}
	res := DeleteProductService.Delete(c.Request.Context(), c.Param("id"))
	c.JSON(consts.StatusOK, res)
}

// 更新商品
func UpdateProduct(c *gin.Context) {
	updateProductService := service.ProductService{}
	if err := c.ShouldBind(&updateProductService); err == nil {
		res := updateProductService.Update(c, c.Param("id"))
		c.JSON(consts.StatusOK, res)
	} else {
		c.JSON(consts.IlleageRequest, ErrorResponse(err))
		util.LogrusObj.Info(err)
	}
}

// SearchProduct 搜索商品
func SearchProduct(c *gin.Context) {
	searchProductService := service.ProductService{}
	if err := c.ShouldBind(&searchProductService); err == nil {
		res := searchProductService.Search(c.Request.Context())
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln(err)
	}
}

// ListProduct 商品列表
func ListProduct(c *gin.Context) {
	listProductService := service.ProductService{}
	if err := c.ShouldBind(&listProductService); err == nil {
		res := listProductService.List(c.Request.Context())
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln(err)
	}
}
