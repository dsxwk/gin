package middleware

import (
	"bytes"
	"gin/common/base"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"io"
)

type Log struct {
	base.BaseMiddleware
}

// Handle 日志中间件
func (s Log) Handle() gin.HandlerFunc {
	return func(c *gin.Context) {
		traceId := uuid.New().String()
		c.Set("traceId", traceId)
		c.Header("Trace-Id", traceId)
		// 读取请求体(防止被读取一次后丢失)
		var bodyBytes []byte
		if c.Request.Body != nil {
			bodyBytes, _ = io.ReadAll(c.Request.Body)
			c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
		}

		c.Next()
	}
}
