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

		// 商品操作
		v1.GET("carousels", api.ListCarousel)   // 轮播图
		v1.GET("products", api.ListProduct)     // 获取商品列表
		v1.GET("products/:id", api.ShowProduct) // 获取商品展示信息
		v1.GET("imgs/:id", api.ListProductImg)  // 获取商品图片
		v1.GET("categories", api.ListCategory)  // 商品分类

		// 需要登录保护
		authed := v1.Group("/")
		authed.Use(middleware.JWT())
		{
			// 用户操作
			authed.PUT("user/update", api.UserUpdate)
			authed.POST("user/updateAvatar", api.UpdateAvatar)
			authed.POST("user/sending-email", api.SendEmail) // 绑定邮箱
			authed.POST("user/valid-email", api.ValidEmail)  // 验证邮箱

			// 商品操作
			authed.POST("product", api.CreateProduct)        // 创建商品
			authed.POST("search_product", api.SearchProduct) // 搜索商品

			// 收藏夹操作
			authed.GET("favorites", api.ListFavorite)          // 展示收藏夹
			authed.POST("favorites", api.CreateFavorite)       // 创建收藏夹
			authed.DELETE("favorites/:id", api.DeleteFavorite) // 删除收藏夹

			// 地址操作
			authed.POST("addresses", api.CreateAddress)       // 创建地址
			authed.GET("addresses/:id", api.GetAddress)       // 获取详细地址
			authed.GET("addresses", api.ListAddress)          // 查看所有地址
			authed.PUT("addresses/:id", api.UpdateAddress)    // 更新地址
			authed.DELETE("addresses/:id", api.DeleteAddress) // 删除地址
		}
	}
	return r
}
