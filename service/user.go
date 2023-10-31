package service

import (
	"context"
	"fmt"
	"mime/multipart"
	"strings"
	"time"

	"github.com/xilepeng/gin-mall/conf"
	"github.com/xilepeng/gin-mall/dao"
	"github.com/xilepeng/gin-mall/model"
	"github.com/xilepeng/gin-mall/pkg/e"
	util "github.com/xilepeng/gin-mall/pkg/utils"
	"github.com/xilepeng/gin-mall/serializer"
	"gopkg.in/mail.v2"
)

type UserService struct {
	NickName string `json:"nick_name" form:"nick_name"`
	UserName string `json:"user_name" form:"user_name"`
	Password string `json:"password" form:"password"`
	Key      string `json:"key" form:"key"` // 秘钥：现阶段前端验证
}

type SendEmailService struct {
	Email         string `json:"email" form:"email"`
	Password      string `json:"password" form:"password"`
	OperationType uint   `json:"operation_type" form:"operation_type"` // 1-绑定邮箱 2-解绑邮箱 3-修改密码
}

type ValidEmailService struct {
}

// Register 注册
func (service *UserService) Register(ctx context.Context) serializer.Response {
	var user model.User
	code := e.SUCCESS
	if service.Key == "" || len(service.Key) != 16 {
		code = e.ERROR
		return serializer.Response{
			Status: code,
			Data:   e.GetMsg(code),
			Msg:    "秘钥长度不足",
		}
	}
	// 1000 --->密文存储 对称加密操作
	util.Encrypt.SetKey(service.Key)
	userDao := dao.NewUserDao(ctx)
	_, exiest, err := userDao.ExistOrNotByUserName(service.UserName)
	if err != nil {
		code = e.ERROR
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	if exiest {
		code = e.ErrorExistUser
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	user = model.User{
		UserName: service.UserName,
		NiceName: service.NickName,
		Status:   model.Active,
		Avatar:   "avatar.jpeg",
		Money:    util.Encrypt.AesEncoding("999999"), // 初始金额加密
	}
	// 密码加密
	if err = user.SetPassword(service.Password); err != nil {
		code = e.ErrorFailEncryption
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	// 创建用户
	err = userDao.CreateUser(&user)
	if err != nil {
		code = e.ERROR
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}

// Login 登录
func (service *UserService) Login(ctx context.Context) serializer.Response {
	var user *model.User
	code := e.SUCCESS
	userDao := dao.NewUserDao(ctx)

	// 判断用户存不存在
	user, exist, err := userDao.ExistOrNotByUserName(service.UserName)
	//  存在 exist == true
	if !exist || err != nil { // ❌ if exist || err == nil
		code = e.ErrorNotExistUser
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Data:   "用户不存在，请先注册！",
		}
	}
	// 校验密码
	if !user.CheckPassword(service.Password) {
		code = e.ErrorNotCompare
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Data:   "密码错误，请重新登录！",
		}
	}
	// token 签发
	token, err := util.GenerateToken(user.ID, service.UserName, 0)
	if err != nil {
		code = e.ErrorAuthCheckTokenFail
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Data:   "token 认证失败!",
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.TokenData{User: serializer.BuildUser(user), Token: token},
	}

}

// Update 用户修改信息
// postman 测试：
// 环境变量：登录后获取 token, 将 token 添加到环境变量
// Headers 中添加  Key: authorization  Value:{{token}}
func (service *UserService) Update(ctx context.Context, uId uint) serializer.Response {
	var user *model.User
	var err error
	code := e.SUCCESS

	// 找到这个用户
	userDao := dao.NewUserDao(ctx)
	user, err = userDao.GetUserById(uId)
	if err != nil {
		code = e.ERROR
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	// 修改昵称 nickname
	if service.NickName != "" {
		user.NiceName = service.NickName
	}
	err = userDao.UpdateUserById(uId, user)
	if err != nil {
		code = e.ERROR
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildUser(user),
	}
}

// Post 用户上传头像（头像更新到本地）
func (service *UserService) Post(ctx context.Context, uId uint, file multipart.File, fileSize int64) serializer.Response {
	code := e.SUCCESS
	var user *model.User
	var err error
	userDao := dao.NewUserDao(ctx)
	user, err = userDao.GetUserById(uId)
	if err != nil {
		code = e.ERROR
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	// 保存图片到本地
	path, err := UploadAvatarToLocalStatic(file, uId, user.UserName)
	if err != nil {
		code = e.ErrorUploadFile
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	user.Avatar = path
	err = userDao.UpdateUserById(uId, user)
	if err != nil {
		code = e.ERROR
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildUser(user),
	}
}

// SendEmailService 发送邮件服务
func (service *SendEmailService) Send(ctx context.Context, uId uint) serializer.Response {
	code := e.SUCCESS
	var address string
	var notice *model.Notice // 绑定邮箱，修改密码 模版通知
	token, err := util.GenerateEmailToken(uId, service.OperationType, service.Email, service.Password)
	if err != nil {
		code = e.ErrorAuthToken
		util.LogrusObj.Infoln("err", err)
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	noticeDao := dao.NewNoticeDao(ctx)
	notice, err = noticeDao.GetNoticeById(service.OperationType)
	if err != nil {
		code = e.ErrorAuthCheckTokenFail
		util.LogrusObj.Infoln("err", err)
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	address = conf.ValidEmail + token // 发送方
	fmt.Println("address = ", address)
	mailStr := notice.Text
	mailText := strings.Replace(mailStr, "Email", address, -1)
	fmt.Println("mailText = ", mailText)
	m := mail.NewMessage()
	m.SetHeader("From", conf.SmtpEmail)
	m.SetHeader("To", service.Email)
	m.SetHeader("Subject", "好饭不怕晚")
	m.SetBody("text/html", mailText)
	d := mail.NewDialer(conf.SmtpHost, 465, conf.SmtpEmail, conf.SmtpPass)
	d.StartTLSPolicy = mail.MandatoryStartTLS
	if err = d.DialAndSend(m); err != nil { // 错误修改 err := d.DialAndSend(m)
		code = e.ErrorSendEmail
		util.LogrusObj.Infoln("err", err)
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}

// Valid 验证邮箱
func (service *ValidEmailService) Valid(ctx context.Context, token string) serializer.Response {
	var userId uint
	var email string
	var password string
	var operationType uint
	code := e.SUCCESS
	// 验证 token
	if token == "" {
		code = e.InvalidParams
	} else {
		claims, err := util.ParseEmailToken(token)
		if err != nil {
			code = e.ErrorAuthToken
			util.LogrusObj.Infoln("err", err)
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
				Error:  err.Error(),
			}
		} else if time.Now().Unix() > claims.ExpiresAt {
			code = e.ErrorAuthCheckTokenTimeout
		} else {
			userId = claims.UserID
			email = claims.Email
			password = claims.Password
			operationType = claims.OperationType
		}
	}
	if code != e.SUCCESS {
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	// 获取该用户的信息
	userDao := dao.NewUserDao(ctx)
	user, err := userDao.GetUserById(userId)
	if err != nil {
		code = e.ERROR
		util.LogrusObj.Infoln("err", err)
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	if operationType == 1 {
		// 绑定邮箱
		user.Email = email
	} else if operationType == 2 {
		// 解绑邮箱
		user.Email = ""
	} else if operationType == 3 {
		// 修改密码
		err = user.SetPassword(password)
		if err != nil {
			code = e.ERROR
			util.LogrusObj.Infoln("err", err)
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
				Error:  err.Error(),
			}
		}
	}
	err = userDao.UpdateUserById(userId, user)
	if err != nil {
		code = e.ERROR
		util.LogrusObj.Infoln("err", err)
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildUser(user),
	}
}
