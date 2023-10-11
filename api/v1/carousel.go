package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xilepeng/gin-mall/service"
)

func ListCarousel(c *gin.Context) {
	var listCarousel service.CarouselService
	if err := c.ShouldBind(&listCarousel); err == nil {
		res := listCarousel.List(c.Request.Context())
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, err)
	}
}
