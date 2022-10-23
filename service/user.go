package service

import (
	"context"
	logging "github.com/sirupsen/logrus"
	"github.com/xilepeng/gin-mall/dao"
	"github.com/xilepeng/gin-mall/model"
	"github.com/xilepeng/gin-mall/pkg/e"
	util "github.com/xilepeng/gin-mall/pkg/utils"
	"github.com/xilepeng/gin-mall/serializer"
)

type UserService struct {
	NickName string `json:"nick_name" form:"nick_name"`
	UserName string `json:"user_name" form:"user_name"`
	Password string `json:"password" form:"password"`
	Key      string `json:"key" form:"key"` // 前端验证
}

func (service UserService) Register(ctx context.Context) serializer.Response {
	code := e.SUCCESS
	if service.Key == "" || len(service.Key) != 16 {
		code = e.ERROR
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Data:   "密钥长度不足",
		}
	}
	util.Encrypt.SetKey(service.Key)
	userDao := dao.NewUserDao(ctx)
	_, exist, err := userDao.ExistOrNotByUserName(service.UserName)
	if err != nil {
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	if exist {
		code = e.ErrorExistUser
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	user := &model.User{
		NickName: service.NickName,
		UserName: service.UserName,
		Status:   model.Active,
		Money:    util.Encrypt.AesEncoding("10000"), // 初始金额
	}
	//加密密码
	if err = user.SetPassword(service.Password); err != nil {
		logging.Info(err)
		code = e.ErrorFailEncryption
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	user.Avatar = "http://q1.qlogo.cn/g?b=qq&nk=294350394&s=640"
	//创建用户
	err = userDao.CreateUser(user)
	if err != nil {
		logging.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}

// Login 用户登陆函数
func (service UserService) Login(ctx context.Context) serializer.Response {
	code := e.SUCCESS
	userDao := dao.NewUserDao(ctx)
	user, exist, err := userDao.ExistOrNotByUserName(service.UserName)
	if !exist { //如果查询不到，返回相应的错误
		logging.Info(err)
		code = e.ErrorUserNotFound
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	if user.CheckPassword(service.Password) == false {
		code = e.ErrorNotCompare
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	token, err := util.GenerateToken(user.ID, service.UserName, 0)
	if err != nil {
		logging.Info(err)
		code = e.ErrorAuthToken
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	return serializer.Response{
		Status: code,
		Data:   serializer.TokenData{User: serializer.BuildUser(user), Token: token},
		Msg:    e.GetMsg(code),
	}
}

// // 用户注册
// func (service *UserService) Register(ctx context.Context) serializer.Response {
// 	var user model.User
// 	code := e.SUCCESS
// 	if service.Key == "" || len(service.Key) != 16 {
// 		code = e.Error
// 		return serializer.Response{
// 			Status: code,
// 			Msg:    e.GetMsg(code),
// 			Error:  "秘钥长度不足",
// 		}
// 	}

// 	// 10000 ----> 密文存储，对称加密操作
// 	utils.Encrypt.SetKey(service.Key)

// 	userDao := dao.NewUserDao(ctx)
// 	_, exist, err := userDao.ExistOrNotByUserName(service.UserName)
// 	if err != nil {
// 		code = e.Error
// 		return serializer.Response{
// 			Status: code,
// 			Msg:    e.GetMsg(code),
// 		}
// 	}
// 	if exist {
// 		code = e.ErrorExistUser
// 		return serializer.Response{
// 			Status: code,
// 			Msg:    e.GetMsg(code),
// 		}
// 	}

// 	user = model.User{

// 		UserName:       service.UserName,
// 		Email:          user.Email,
// 		PasswordDigest: "",
// 		NickName:       service.NickName,
// 		Status:         model.Active,
// 		Avatar:         "avatar.png",
// 		Money:          utils.Encrypt.AesEncoding("10000"),
// 	}

// 	// 密码加密
// 	if err = user.SetPassword(service.Password); err != nil {
// 		code = e.ErrorFailEncryption
// 		return serializer.Response{
// 			Status: code,
// 			Msg:    e.GetMsg(code),
// 		}
// 	}

// 	// 创建用户
// 	err = userDao.CreateUser(user)
// 	if err != nil {
// 		code = e.Error
// 		return serializer.Response{
// 			Status: code,
// 			Msg:    e.GetMsg(code),
// 		}
// 	}
// 	return serializer.Response{
// 		Status: code,
// 		Msg:    e.GetMsg(code),
// 	}
// }

// // 用户登录
// func (service *UserService) Login(ctx context.Context) serializer.Response {
// 	var user *model.User
// 	code := e.Success
// 	userDao := dao.NewUserDao(ctx)

// 	// 判断用户是否存在
// 	user, exist, err := userDao.ExistOrNotByUserName(service.UserName)
// 	if !exist || err != nil {
// 		code = e.ErrorExistUserNotFound
// 		return serializer.Response{
// 			Status: code,
// 			Msg:    e.GetMsg(code),
// 			Data:   "用户不存在，请先注册！",
// 		}
// 	}

// 	// 校验密码
// 	if user.CheckPassword(service.Password) == false {
// 		code = e.ErrorNotCompare
// 		return serializer.Response{
// 			Status: code,
// 			Msg:    e.GetMsg(code),
// 		}
// 	}

// 	// token 签发          http 无状态(认证 token)
// 	token, err := utils.GenerateToken(user.ID, service.UserName, 0)
// 	if err != nil {
// 		code = e.ErrorAuthToken
// 		return serializer.Response{
// 			Status: code,
// 			Msg:    e.GetMsg(code),
// 			Data:   "token 认证失败",
// 		}
// 	}

// 	return serializer.Response{
// 		Status: code,
// 		Msg:    e.GetMsg(code),
// 		Data: serializer.TokenData{
// 			User: serializer.BuildUser(user), Token: token,
// 		},
// 	}
// }
