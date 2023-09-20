package serializer

type Response struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
	Msg    string      `json:"msg"`
	Error  string      `json:"error"`
}

// 带 token 的 data
type TokenData struct {
	User  interface{} `json:"user"`
	Token string      `json:"token"`
}
