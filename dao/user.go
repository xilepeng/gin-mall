package dao

import (
	"context"

	"github.com/xilepeng/gin-mall/model"
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
func (dao *UserDao) ExistOrNotByUserName(userName string) (user *model.User, exist bool, err error) {
	err = dao.DB.Model(&model.User{}).Where("user_name=?", userName).Find(&user).Error
	if user == nil || err == gorm.ErrDryRunModeUnsupported {
		return nil, false, err
	}
	return user, true, nil
}

// 创建用户
func (dao *UserDao) CreateUser(user model.User) (err error) {
	
}
