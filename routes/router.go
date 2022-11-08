package routes

import (
	"net/http"

	api "gin-mall/api/v1"
	"gin-mall/middleware"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	r.Use(middleware.Cors())

	r.StaticFS("/static", http.Dir("./static"))

	v1 := r.Group("api/v1")
	{
		//v1.GET("ping", func(c *gin.Context) { c.JSON(200, "success") })
		// 用户操作
		v1.POST("user/register", api.UserRegister)
		v1.POST("user/login", api.UserLogin)
	}
	return r
}
