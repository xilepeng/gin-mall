package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	api "github.com/xilepeng/gin-mall/api/v1"
	"github.com/xilepeng/gin-mall/middleware"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.Cors())
	r.StaticFS("/static", http.Dir("./static"))
	v1 := r.Group("api/v1")
	{
		v1.GET("ping", func(c *gin.Context) { c.JSON(200, "success") })
		// 用户操作
		v1.POST("user/register", api.UserRegisterHandler())
		v1.POST("user/login", api.UserLoginHandler())

		authed := v1.Group("/") // 需要登录保护
		authed.Use(middleware.JWT())
		{
			// 用户操作
			authed.PUT("user", api.UserUpdateHandler()) // 测试
		}
	}
	return r
}
