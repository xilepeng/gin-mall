package e

const (
	SUCCESS       = 200
	Error         = 500
	InvalidParams = 400

	// User 模块错误 3xxxx
	ErrorExistUser             = 3001
	ErrorFailEncryption        = 3002
	ErrorExistUserNotFound     = 3003
	ErrorNotCompare            = 3004
	ErrorAuthCheckTokenFail    = 3005
	ErrorAuthCheckTokenTimeout = 30006
	ErrorUploadFail            = 3007

	// Product 模块错误 4xxxx

)
