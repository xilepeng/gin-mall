package serializer

import (
	"context"

	"github.com/xilepeng/gin-mall/conf"
	"github.com/xilepeng/gin-mall/repository/db/dao"
	"github.com/xilepeng/gin-mall/repository/db/model"
)

type Favorite struct {
	UserID        uint   `json:"user_id"`
	ProductID     uint   `json:"product_id"`
	CreateAt      int64  `json:"create_at"`
	Name          string `json:"name"`
	CategoryID    uint   `json:"category_id"`
	Title         string `json:"title"`
	Info          string `json:"info"`
	ImgPath       string `json:"img_path"`
	Price         string `json:"price"`
	DiscountPrice string `json:"discount_price"`
	BossID        uint   `json:"boss_id"`
	Num           int    `json:"num"`
	OnSale        bool   `json:"on_sale"`
}

func BuildFavorite(favorite *model.Favorite, product *model.Product, boss *model.User) Favorite {
	return Favorite{
		UserID:        favorite.ID,
		ProductID:     favorite.ProductID,
		CreateAt:      favorite.CreatedAt.Unix(),
		Name:          product.Name,
		CategoryID:    product.CategoryID,
		Title:         product.Title,
		Info:          product.Info,
		ImgPath:       conf.Host + conf.HttpPort + conf.ProductPath + product.ImgPath,
		Price:         product.Price,
		DiscountPrice: product.DiscountPrice,
		BossID:        boss.ID,
		Num:           product.Num,
		OnSale:        product.OnSale,
	}
}

func BuildFavorites(ctx context.Context, items []*model.Favorite) (favorites []Favorite) {
	productDao := dao.NewProductDao(ctx)
	bossDao := dao.NewUserDao(ctx)
	for _, item := range items {
		product, err := productDao.GetProductById(item.ProductID)
		if err != nil {
			continue
		}
		boss, err := bossDao.GetUserById(item.UserID)
		if err != nil {
			continue
		}
		favorite := BuildFavorite(item, product, boss)
		favorites = append(favorites, favorite)

	}
	return
}
