package middleware

import (
	"gin/common/base"
	"gin/config"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Cors struct {
	base.BaseMiddleware
}

// Handle 跨域中间件
func (s Cors) Handle() gin.HandlerFunc {
	return func(c *gin.Context) {
		if config.Conf.Cors.Enabled {
			c.Header("Access-Control-Allow-Origin", config.Conf.Cors.AllowOrigin)
			c.Header("Access-Control-Allow-Headers", config.Conf.Cors.AllowHeaders)
			c.Header("Access-Control-Expose-Headers", config.Conf.Cors.ExposeHeaders)
			c.Header("Access-Control-Allow-Methods", config.Conf.Cors.AllowMethods)
			c.Header("Access-Control-Allow-Credentials", config.Conf.Cors.AllowCredentials)

			// 放行所有OPTIONS方法
			if c.Request.Method == "OPTIONS" {
				c.AbortWithStatus(http.StatusNoContent)
			}
		}

		c.Next()
	}
}
