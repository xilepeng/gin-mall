package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	util "github.com/xilepeng/gin-mall/pkg/utils"
	"github.com/xilepeng/gin-mall/service"
)

// CreateCart 创建购物车
func CreateCart(c *gin.Context) {
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	createCartsService := service.CartService{}
	if err := c.ShouldBind(&createCartsService); err == nil {
		res := createCartsService.Create(c.Request.Context(), claim.ID)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln(err)
	}
}

// DeleteCart 删除购物车
func DeleteCart(c *gin.Context) {
	deleteCartService := service.CartService{}
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&deleteCartService); err == nil {
		res := deleteCartService.Delete(c.Request.Context(), claim.ID, c.Param("id"))
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln(err)
	}
}

// UpdateCart 修改购物车
func UpdateCart(c *gin.Context) {
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	updateCartsService := service.CartService{}
	if err := c.ShouldBind(&updateCartsService); err == nil {
		res := updateCartsService.Update(c.Request.Context(), claim.ID, c.Param("id"))
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln(err)
	}
}

// ListCart 展示购物车信息
func ListCart(c *gin.Context) {
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	listCartService := service.CartService{}
	if err := c.ShouldBind(&listCartService); err == nil {
		res := listCartService.Show(c.Request.Context(), claim.ID)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln(err)
	}
}
