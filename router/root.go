package router

import (
	"gin/app/middleware"
	_ "gin/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var (
	jwtMiddleware = middleware.Jwt{}.Handle()
)

// Router 路由接口
type Router interface {
	RegisterRoutes(router *gin.RouterGroup)
}

// LoadRouters 加载路由
func LoadRouters(router *gin.Engine) {
	var (
		// 统一路由分组
		v1    = router.Group("api/v1")
		login LoginRouter
		user  UserRouter
		// ... 其他路由
	)

	// 健康检查
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"code": 0,
			"msg":  "pong",
			"data": []string{},
		})
	})

	// Swagger 路由
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// 登录
	login.RegisterRoutes(v1.Group("")) // new(LoginRouter).RegisterRoutes(v1)

	// 需要权限
	auth := v1.Group("", jwtMiddleware)
	{
		// 用户
		user.RegisterRoutes(auth)
		// ... 其他路由
	}
}
