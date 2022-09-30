package dao

import (
	"fmt"

	"github.com/xilepeng/gin-mall/model"
)

func migration() {
	err := _db.Set("gorm:table_options", "charset=utf8mb4").AutoMigrate(
		&model.User{},
		&model.Address{},
		&model.Admin{},
		&model.Carousel{},
		&model.Favorite{},
		&model.Notice{},
		&model.Order{},
		&model.Product{},
		&model.ProductImg{},
		&model.Category{},
		&model.Cart{},
	)
	if err != nil {
		fmt.Println(err)
	}
	return
}
