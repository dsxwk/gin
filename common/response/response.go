package response

import (
	"gin/common/errcode"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Response 通用响应结构体
type Response struct {
	Code int64       `json:"code"` // 错误码
	Msg  string      `json:"msg"`  // 提示信息
	Data interface{} `json:"data"` // 返回数据
}

// JSON 输出 JSON 响应
func (r Response) JSON(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, r)
	c.Abort()
}

// Success 返回成功响应
func Success(c *gin.Context, e *errcode.ErrorCode) {
	var (
		resp Response
	)

	switch e {
	case nil:
		resp = Response{
			Code: errcode.Success().Code,
			Msg:  errcode.Success().Msg,
			Data: []interface{}{},
		}
	default:
		if e.Data == nil {
			e.Data = []interface{}{}
		}
		resp = Response{
			Code: e.Code,
			Msg:  e.Msg,
			Data: e.Data,
		}
	}

	resp.JSON(c)
}

// Error 返回失败响应,可传ErrorCode
func Error(c *gin.Context, e *errcode.ErrorCode) {
	var (
		resp Response
	)

	switch e {
	case nil:
		resp = Response{
			Code: errcode.SystemError().Code,
			Msg:  errcode.SystemError().Msg,
			Data: []interface{}{},
		}
	default:
		if e.Data == nil {
			e.Data = []interface{}{}
		}
		resp = Response{
			Code: e.Code,
			Msg:  e.Msg,
			Data: e.Data,
		}
	}

	resp.JSON(c)
}
