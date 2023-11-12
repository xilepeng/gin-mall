package v1

import (
	"encoding/json"
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/xilepeng/gin-mall/conf"
	"github.com/xilepeng/gin-mall/consts"
	"github.com/xilepeng/gin-mall/serializer"
)

func ErrorResponse(err error) serializer.Response {
	if ve, ok := err.(validator.ValidationErrors); ok {
		for _, e := range ve {
			field := conf.T(fmt.Sprintf("Field.%s", e.Field()))
			tag := conf.T(fmt.Sprintf("Tag.Valid.%s", e.Tag()))
			return serializer.Response{
				Status: consts.IlleageRequest,
				Msg:    fmt.Sprintf("%s%s", field, tag),
				Error:  fmt.Sprint(err),
			}
		}
	}

	if _, ok := err.(*json.UnmarshalTypeError); ok {
		return serializer.Response{
			Status: 400,
			Msg:    "JSON 类型不匹配",
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: 400,
		Msg:    "参数错误❌ ",
		Error:  err.Error(),
	}
}
