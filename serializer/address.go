package serializer

import "github.com/xilepeng/gin-mall/model"

type Address struct {
	Id       uint   `json:"id"`
	UserId   uint   `json:"user_id"`
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
	CreateAt int64  `json:"create_at"`
}

func BuildAddress(item *model.Address) Address {
	return Address{
		Id:       item.ID,
		UserId:   item.UserID,
		Name:     item.Name,
		Phone:    item.Phone,
		Address:  item.Address,
		CreateAt: item.CreatedAt.Unix(),
	}
}

func BuildAddresses(items []*model.Address) (addresses []Address) {
	for _, item := range items {
		address := BuildAddress(item)
		addresses = append(addresses, address)
	}
	return
}
