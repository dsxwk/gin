package middleware

import (
	"gin/common/base"
	"github.com/gin-gonic/gin"
)

type Jwt struct {
	base.BaseMiddleware
}

// Handle jwt中间件
func (s Jwt) Handle() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Define your middleware logic here
		c.Next()
	}
}
