package serializer

import (
	"github.com/xilepeng/gin-mall/conf"
	"github.com/xilepeng/gin-mall/model"
)

type ProductImg struct {
	ProductId uint   `json:"product_id"`
	ImgPath   string `json:"img_path"`
}

func BuildProductImg(item *model.ProductImg) ProductImg {
	return ProductImg{
		ProductId: item.ProductId,
		ImgPath:   conf.Host + conf.HttpPort + conf.ProductPath + item.ImgPath,
	}
}

func BuildProductImgs(items []*model.ProductImg) (productImg []ProductImg) {
	for _, item := range items {
		product := BuildProductImg(item)
		productImg = append(productImg, product)

	}
	return
}
