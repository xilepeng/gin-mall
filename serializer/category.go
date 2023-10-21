package serializer

import "github.com/xilepeng/gin-mall/model"

type Category struct {
	Id           uint   `json:"id"`
	CategoryName string `json:"category_name"`
	CreateAt     int64  `json:"create_at"`
}

func BuildCategory(item *model.Category) Category {
	return Category{
		Id:           item.ID,
		CategoryName: item.CategoryName,
		CreateAt:     item.CreatedAt.Unix(),
	}
}

func BuildCategorys(items []model.Category) (categorys []Category) {
	for _, item := range items {
		category := BuildCategory(&item)
		categorys = append(categorys, category)
	}
	return categorys
}
