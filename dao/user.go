package dao

import (
	"context"

	"gin-mall/model"
	"gorm.io/gorm"
)

type UserDao struct {
	*gorm.DB
}

func NewUserDao(ctx context.Context) *UserDao {
	return &UserDao{NewDBClient(ctx)}
}

func NewUserDaoByDB(db *gorm.DB) *UserDao {
	return &UserDao{db}
}

// 根据 ExistOrNotByUserName 判断是否存在该名字
// func (dao *UserDao) ExistOrNotByUserName(userName string) (user *model.User, exist bool, err error) {
// 	var count int64

// 	err = dao.DB.Model(&model.User{}).Where("user_name=?", userName).Count(&count).Error
// 	if count == 0 {
// 		return nil, false, err
// 	}
// 	// err = dao.DB.Model(&model.User{}).Where("user_name=?", userName).
// 	// 	First(&user).Error
// 	// if err != nil {
// 	// 	return nil, false, err
// 	// }
// 	return user, true, nil
// }

// ExistOrNotByUserName 根据username判断是否存在该名字
func (dao *UserDao) ExistOrNotByUserName(userName string) (user *model.User, exist bool, err error) {
	var count int64
	err = dao.DB.Model(&model.User{}).Where("user_name=?", userName).
		Count(&count).Error
	if count == 0 {
		return nil, false, err
	}
	err = dao.DB.Model(&model.User{}).Where("user_name=?", userName).
		First(&user).Error
	if err != nil {
		return nil, false, err
	}
	return user, true, nil
}

// CreateUser 创建用户
func (dao *UserDao) CreateUser(user *model.User) (err error) {
	err = dao.DB.Model(&model.User{}).Create(&user).Error
	return
}
