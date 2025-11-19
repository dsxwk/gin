package middleware

import (
	"bytes"
	"gin/common/base"
	"gin/common/global"
	"gin/utils/ctx"
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"github.com/google/uuid"
	"github.com/gookit/goutil/strutil"
	"io"
	"net/http"
	"strings"
	"time"
)

type Logger struct {
	base.BaseMiddleware
}

// Handle 日志中间件
func (s Logger) Handle() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			body   []byte
			params any
			m      map[string]any
		)
		if c.Request.Body != nil {
			body, _ = io.ReadAll(c.Request.Body)
			c.Request.Body = io.NopCloser(bytes.NewBuffer(body))
		}

		if c.Request.Method == http.MethodGet || c.Request.Method == http.MethodDelete {
			params = c.Request.URL.Query()
		} else if len(body) > 0 {
			if err := json.Unmarshal(body, &m); err != nil {
				params = string(body)
			} else {
				params = m
			}
		} else {
			params = map[string]any{}
		}

		lang := strings.ToLower(c.GetHeader("Accept-Language"))
		if strutil.StartsWith(lang, "en") {
			lang = "en"
		} else {
			lang = "zh"
		}

		traceId := uuid.New().String()
		// 绑定 goroutine → traceId
		ctx.BindTraceId(traceId)
		c.Set(ctx.KeyTraceId, traceId)
		c.Set(ctx.KeyIp, c.ClientIP())
		c.Set(ctx.KeyPath, c.Request.URL.Path)
		c.Set(ctx.KeyMethod, c.Request.Method)
		c.Set(ctx.KeyParams, params)
		c.Set(ctx.KeyLang, lang)
		ctx.SetContext(ctx.KeyLang, c)
		c.Header("Trace-Id", traceId)
		reqCtx := c.Request.Context()
		c.Request = c.Request.WithContext(reqCtx)

		start := time.Now()
		c.Set(ctx.KeyStartTime, start) // 保存开始时间
		c.Next()
		cost := float64(time.Since(start).Nanoseconds()) / 1e6
		c.Set(ctx.KeyMs, cost)

		// 记录请求日志
		if global.Config.Log.Access {
			global.Log.Info(c, "Access Log")
		}
	}
}
