package errcode

import "fmt"

// ErrorCode 公共错误码结构体
type ErrorCode struct {
	Code int64       `json:"code"` // 错误码
	Msg  string      `json:"msg"`  // 错误描述
	Data interface{} `json:"data"` // 返回数据
}

// Error 实现error接口
func (e ErrorCode) Error() string {
	return fmt.Sprintf("%s", e.Msg)
}

func (e ErrorCode) WithMsg(msg string) ErrorCode {
	return ErrorCode{
		Code: e.Code,
		Msg:  msg,
		Data: e.Data,
	}
}

func (e ErrorCode) WithData(data interface{}) ErrorCode {
	return ErrorCode{
		Code: e.Code,
		Msg:  e.Msg,
		Data: data,
	}
}

func Success() ErrorCode {
	return ErrorCode{Code: 0, Msg: "Success"}
}

func Redirect() ErrorCode {
	return ErrorCode{Code: 301, Msg: "Redirect"}
}

func ArgsError() ErrorCode {
	return ErrorCode{Code: 400, Msg: "Invalid arguments"}
}

func Unauthorized() ErrorCode {
	return ErrorCode{Code: 401, Msg: "Unauthorized"}
}

func NotFound() ErrorCode {
	return ErrorCode{Code: 404, Msg: "Resource not found"}
}

func SystemError() ErrorCode {
	return ErrorCode{Code: 500, Msg: "Internal server error"}
}
