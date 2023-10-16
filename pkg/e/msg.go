package e

var MsgFlags = map[int]string{
	SUCCESS:       "ok",
	Error:         "fail",
	InvalidParams: "参数错误",

	ErrorExistUser:             "用户名已存在",
	ErrorFailEncryption:        "密码加密失败",
	ErrorExistUserNotFound:     "用户不存在",
	ErrorNotCompare:            "密码错误",
	ErrorAuthCheckTokenFail:    "token 认证失败",
	ErrorAuthCheckTokenTimeout: "token 过期",
	ErrorUploadFail:            "图片上传失败",

	ErrorSendEmail: "邮件发送失败",
	ErrorAuthToken: "Token生成失败",

	ErrorProductImgUpload: "图片上传错误",
}

// GetMsg 获取状态码对应的信息
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if !ok {
		return MsgFlags[Error]
	}
	return msg
}
