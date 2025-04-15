package controller

type MyCode int64

const (
	Success          MyCode = 200
	ParamErr         MyCode = 400
	TokenInvalidErr  MyCode = 401
	TokenCheckErr    MyCode = 402
	ParamLostErr     MyCode = 4001
	SystemErr        MyCode = 500
	UserExistErr     MyCode = 5001
	UserNotLogin     MyCode = 5002
	UserAccountErr   MyCode = 5003
	UserNotExistErr  MyCode = 5004
	ParamPasswordErr MyCode = 5005
	ActivitySignErr  MyCode = 5006
)

var msgFlags = map[MyCode]string{
	Success:          "success",
	ParamErr:         "参数异常",
	TokenInvalidErr:  "无效的token",
	TokenCheckErr:    "token验证失败，请重新登录",
	ParamLostErr:     "参数缺失",
	SystemErr:        "系统异常",
	UserExistErr:     "用户已存在",
	UserNotLogin:     "用户未登录",
	UserAccountErr:   "账号或密码错误",
	UserNotExistErr:  "用户不存在",
	ParamPasswordErr: "原密码输入错误",
	ActivitySignErr:  "活动已报名",
}

func (c MyCode) Msg() string {
	msg, ok := msgFlags[c]
	if ok {
		return msg
	}
	return msgFlags[SystemErr]
}
