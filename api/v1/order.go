package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	util "github.com/xilepeng/gin-mall/pkg/utils"
	"github.com/xilepeng/gin-mall/service"
)

// CreateOrder 新建收藏
func CreateOrder(c *gin.Context) {
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	createOrderService := service.OrderService{}
	if err := c.ShouldBind(&createOrderService); err == nil {
		res := createOrderService.Create(c.Request.Context(), claim.ID)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln(err)
	}
}

// DeleteOrder 删除收藏夹
func DeleteOrder(c *gin.Context) {
	deleteOrderService := service.OrderService{}
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&deleteOrderService); err == nil {
		res := deleteOrderService.Delete(c.Request.Context(), claim.ID, c.Param("id"))
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln(err)
	}
}

// ShowOrder 新建收藏
func ShowOrder(c *gin.Context) {
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	showOrdersService := service.OrderService{}
	if err := c.ShouldBind(&showOrdersService); err == nil {
		res := showOrdersService.Show(c.Request.Context(), claim.ID, c.Param("id"))
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln(err)
	}
}

// ListOrder 获取商品展示信息
func ListOrder(c *gin.Context) {
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	listOrderService := service.OrderService{}
	if err := c.ShouldBind(&listOrderService); err == nil {
		res := listOrderService.List(c.Request.Context(), claim.ID)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln(err)
	}
}
