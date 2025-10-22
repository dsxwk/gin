package middleware

import (
	"bytes"
	"gin/common/base"
	"gin/common/global"
	"gin/utils/ctx"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"io"
	"net/http"
)

type Logger struct {
	base.BaseMiddleware
}

// Handle 日志中间件
func (s Logger) Handle() gin.HandlerFunc {
	return func(c *gin.Context) {
		traceId := uuid.New().String()
		c.Set(ctx.KeyTraceID, traceId)
		c.Header("Trace-Id", traceId)

		var (
			body []byte
		)
		if c.Request.Body != nil {
			body, _ = io.ReadAll(c.Request.Body)
			c.Request.Body = io.NopCloser(bytes.NewBuffer(body))
		}

		c.Set(ctx.KeyTraceID, traceId)
		if c.Request.Method == http.MethodGet || c.Request.Method == http.MethodDelete {
			c.Set(ctx.KeyParams, c.Request.URL.RawQuery)
		} else {
			c.Set(ctx.KeyParams, string(body))
		}

		ctx.SetContext(ctx.KeyLogger, c)

		// 是否记录请求日志
		if global.Config.Log.Access {
			global.Log.Info("Access Log")
		}

		c.Next()

		// 清理
		ctx.ClearContext(ctx.KeyLogger)
	}
}
