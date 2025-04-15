package ginx

import (
	"fmt"
)

type MyError struct {
	Code int    `json:"code"`
	Msg  string `json:"message"`
	Err  error  `json:"-"` // 原始错误，不暴露给客户端
}

// 错误信息
func (e *MyError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %v", e.Msg, e.Err)
	}
	return e.Msg
}

// Unwrap 实现 errors.Unwrap 接口
func (e *MyError) Unwrap() error {
	return e.Err
}

// 预定义错误
var (
	Success          = &MyError{Code: 200, Msg: "success"}
	ParamErr         = &MyError{Code: 400, Msg: "参数异常"}
	TokenInvalidErr  = &MyError{Code: 401, Msg: "无效的token"}
	TokenCheckErr    = &MyError{Code: 402, Msg: "token验证失败，请重新登录"}
	ParamLostErr     = &MyError{Code: 4001, Msg: "参数缺失"}
	SystemErr        = &MyError{Code: 500, Msg: "系统异常"}
	UserExistErr     = &MyError{Code: 5001, Msg: "用户已存在"}
	UserNotLogin     = &MyError{Code: 5002, Msg: "用户未登录"}
	UserAccountErr   = &MyError{Code: 5003, Msg: "账号或密码错误"}
	UserNotExistErr  = &MyError{Code: 5004, Msg: "用户不存在"}
	ParamPasswordErr = &MyError{Code: 5005, Msg: "原密码输入错误"}
	ActivitySignErr  = &MyError{Code: 5006, Msg: "活动已报名"}
)
