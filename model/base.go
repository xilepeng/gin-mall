package model

// 分页
type BasePage struct {
	PageNum int `form:"pageNum"`
	PageSize int `form:"pageSize"`
}