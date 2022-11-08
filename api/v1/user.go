package v1

import (
	"net/http"

	"gin-mall/service"
	"github.com/gin-gonic/gin"
)

func UserRegister(c *gin.Context) {
	var userRegister service.UserService
	if err := c.ShouldBind(&userRegister); err == nil {
		res := userRegister.Register(c.Request.Context())
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, err)
	}
}

func UserLogin(c *gin.Context) {
	var userLogin service.UserService
	if err := c.ShouldBind(&userLogin); err == nil {

		//res := userLogin.Register(c.Request.Context())  // 导致如下错误
		//{
		//	"status": 500,
		//	"data": "密钥长度不足~",
		//	"msg": "fail",
		//	"error": ""
		//}
		res := userLogin.Login(c.Request.Context())
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, err)
	}
}
