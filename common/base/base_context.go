package base

import (
	"context"
	"gin/common/ctxkey"
)

// Context 统一管理context
type Context struct {
	ctx context.Context
}

// Set 设置ctx
func (s *Context) Set(ctx context.Context) {
	s.ctx = ctx
}

// Get 获取ctx
func (s *Context) Get() context.Context {
	return s.ctx
}

// TraceId 获取traceId
func (s *Context) TraceId() string {
	traceId, ok := s.ctx.Value(ctxkey.TraceIdKey).(string)
	if !ok || traceId == "" {
		traceId = "unknown"
	}
	return traceId
}

// GetLang 获取语言
func (s *Context) GetLang() string {
	return s.ctx.Value(ctxkey.LangKey).(string)
}

// GetIp 获取ip
func (s *Context) GetIp() string {
	return s.ctx.Value(ctxkey.IpKey).(string)
}

// GetPath 获取请求路径
func (s *Context) GetPath() string {
	return s.ctx.Value(ctxkey.PathKey).(string)
}

// GetMethod 获取请求方法
func (s *Context) GetMethod() string {
	return s.ctx.Value(ctxkey.MethodKey).(string)
}

// GetParams 获取请求参数
func (s *Context) GetParams() string {
	return s.ctx.Value(ctxkey.ParamsKey).(string)
}

// GetMs 获取耗时
func (s *Context) GetMs() string {
	return s.ctx.Value(ctxkey.MsKey).(string)
}

// GetStartTime 获取请求开始时间
func (s *Context) GetStartTime() string {
	return s.ctx.Value(ctxkey.StartTimeKey).(string)
}
