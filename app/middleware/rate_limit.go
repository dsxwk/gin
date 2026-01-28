package middleware

import (
	"gin/common/base"
	"gin/common/errcode"
	"gin/common/response"
	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
	"sync"
)

type RateLimit struct {
	base.BaseMiddleware
	limiter *rate.Limiter
}

var err = errcode.RateLimitError()

// Handle 限流中间件
func (s RateLimit) Handle() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !s.limiter.Allow() {
			response.Error(c, &err)
			return
		}
		c.Next()
	}
}

// NewRateLimit 创建限流中间件
// @param r rate.Limit 每秒产生多少token
// @param burst int 桶容量
func NewRateLimit(r rate.Limit, burst int) *RateLimit {
	return &RateLimit{
		limiter: rate.NewLimiter(r, burst),
	}
}

var userLimiters = struct {
	sync.Mutex
	m map[string]*rate.Limiter
}{
	m: make(map[string]*rate.Limiter),
}

// UserRateLimit 用户限流
// @param r rate.Limit 每秒产生多少token
// @param burst int 桶容量
func UserRateLimit(r rate.Limit, burst int) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.GetString("user.id")
		if userID == "" {
			c.Next()
			return
		}

		userLimiters.Lock()
		limiter, ok := userLimiters.m[userID]
		if !ok {
			limiter = rate.NewLimiter(r, burst)
			userLimiters.m[userID] = limiter
		}
		userLimiters.Unlock()

		if !limiter.Allow() {
			response.Error(c, &err)
			return
		}

		c.Next()
	}
}

var ipLimiters = struct {
	sync.Mutex
	m map[string]*rate.Limiter
}{
	m: make(map[string]*rate.Limiter),
}

// IpRateLimit ip限流
// @param r rate.Limit 每秒产生多少token
// @param burst int 桶容量
func IpRateLimit(r rate.Limit, burst int) gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()

		ipLimiters.Lock()
		limiter, ok := ipLimiters.m[ip]
		if !ok {
			limiter = rate.NewLimiter(r, burst)
			ipLimiters.m[ip] = limiter
		}
		ipLimiters.Unlock()

		if !limiter.Allow() {
			response.Error(c, &err)
			return
		}

		c.Next()
	}
}
