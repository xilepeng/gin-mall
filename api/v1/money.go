package v1

import (
	"github.com/gin-gonic/gin"
	util "github.com/xilepeng/gin-mall/pkg/utils"
	"github.com/xilepeng/gin-mall/service"
)

func ShowMoney(c *gin.Context) {
	showMoneyService := service.ShowMoneyService{}
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&showMoneyService); err == nil {
		res := showMoneyService.Show(c.Request.Context(), claim.ID)
		c.JSON(200, res)
	} else {
		c.JSON(400, err)
	}
}
