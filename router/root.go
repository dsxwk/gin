package router

import (
	"gin/app/middleware"
	_ "gin/docs"
	"gin/utils"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

var (
	jwtMiddleware = middleware.Jwt{}.Handle()
)

// LoadRouters 加载路由
func LoadRouters(router *gin.Engine) {
	// 健康检查
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"code": 0,
			"msg":  "pong",
			"data": []string{},
		})
	})

	// 静态文件
	router.StaticFS("/public", http.Dir(utils.GetRootPath()+"/public"))

	// Swagger 文档
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// 路由分组
	v1 := router.Group("api/v1")
	auth := v1.Group("", jwtMiddleware)

	// 自动注册
	AutoLoads(v1, auth)
}
