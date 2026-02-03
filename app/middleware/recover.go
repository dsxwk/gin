package middleware

import (
	"fmt"
	"gin/common/base"
	"gin/common/ctxkey"
	"gin/common/errcode"
	"gin/common/response"
	"github.com/gin-gonic/gin"
	"runtime"
)

type Recover struct {
	base.BaseMiddleware
}

type ErrData struct {
	TraceId string      `json:"traceId"`
	Error   interface{} `json:"error"`
	IP      string      `json:"ip"`
	Lang    string      `json:"lang"`
	Path    string      `json:"path"`
	Method  string      `json:"method"`
	Params  interface{} `json:"params"`
	Stack   []string    `json:"stack"`
}

// Handle recover中间件
func (s Recover) Handle() gin.HandlerFunc {
	return func(c *gin.Context) {
		errCode := errcode.SystemError()
		defer func() {
			if e := recover(); e != nil {
				ctx := c.Request.Context()

				errCode = errCode.WithData(&ErrData{
					TraceId: ctx.Value(ctxkey.TraceIdKey).(string),
					Error:   e,
					IP:      ctx.Value(ctxkey.IpKey).(string),
					Lang:    ctx.Value(ctxkey.LangKey).(string),
					Path:    ctx.Value(ctxkey.PathKey).(string),
					Method:  ctx.Value(ctxkey.MethodKey).(string),
					Params:  ctx.Value(ctxkey.ParamsKey),
					Stack:   getStackTrace(3),
				})

				response.Error(c, &errCode)
				return
			}
		}()

		c.Next()
	}
}

// getStackTrace 获取堆栈
func getStackTrace(skip int) []string {
	const maxDepth = 32
	pc := make([]uintptr, maxDepth)
	n := runtime.Callers(skip, pc)
	pc = pc[:n]

	trace := []string{}
	for _, p := range pc {
		fn := runtime.FuncForPC(p)
		if fn == nil {
			trace = append(trace, "unknown")
			continue
		}
		file, line := fn.FileLine(p)
		trace = append(trace, fmt.Sprintf("%s:%d %s", file, line, fn.Name()))
	}

	return trace
}
