package model

import (
	"strconv"

	"github.com/xilepeng/gin-mall/repository/cache"
	"gorm.io/gorm"
)

// 定义mysql的模型

type Product struct {
	gorm.Model
	Name          string
	CategoryID    uint
	Title         string
	Info          string
	ImgPath       string
	Price         string
	DiscountPrice string // 折后价
	OnSale        bool   `gorm:"default:false"`
	Num           int
	BossID        uint
	BossName      string
	BossAvatar    string
}

func (product *Product) View() uint64 {
	countStr, _ := cache.RedisClient.Get(cache.ProductViewKey(product.ID)).Result()
	count, _ := strconv.ParseUint(countStr, 10, 64)
	return count
}

func (product *Product) AddView() {
	// 增加商品点击数
	cache.RedisClient.Incr(cache.ProductViewKey(product.ID))
	cache.RedisClient.ZIncrBy(cache.RankKey, 1, strconv.Itoa(int(product.ID)))

}
