package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xilepeng/gin-mall/service"
)

func ListCategory(c *gin.Context) {
	var ListCategory service.CategoryService
	if err := c.ShouldBind(&ListCategory); err == nil {
		res := ListCategory.List(c.Request.Context())
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, err)
	}
}
