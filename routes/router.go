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
		v1.GET("ping", func(c *gin.Context) {
			c.JSON(200, "success")
		})

		// 用户操作
		v1.POST("user/register", api.UserRegister)
		v1.POST("user/login", api.UserLogin)

		// 需要登录保护
		authed := v1.Group("/")
		authed.Use(middleware.JWT())
		{
			// 用户操作
			authed.PUT("user/update", api.UserUpdate)
			authed.POST("user/updateAvatar", api.UpdateAvatar)
			authed.POST("user/sending-email", api.SendEmail) // 绑定邮箱
			authed.POST("user/valid-email", api.ValidEmail)  // 验证邮箱
		}
	}
	return r
}
