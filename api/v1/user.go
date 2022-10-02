package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xilepeng/gin-mall/service"
)

func UserRegister(c *gin.Context) {
	var userRegister service.UserService
	if err := c.ShouldBind(&userRegister); err == nil {
		res := userRegister.Register(c.Request.Context())
		c.JSON(http.StatusOK, res)
	}
}
