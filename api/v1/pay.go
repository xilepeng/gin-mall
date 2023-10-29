package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	util "github.com/xilepeng/gin-mall/pkg/utils"
	"github.com/xilepeng/gin-mall/service"
)

func OrderPay(c *gin.Context) {
	orderPay := service.OrderPay{}
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&orderPay); err == nil {
		res := orderPay.OrderPay(c.Request.Context(), claim.ID)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, err)
	}
}
