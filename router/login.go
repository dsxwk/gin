package router

import (
	"gin/app/controller/v1"
	"github.com/gin-gonic/gin"
)

// LoginRouter 登录路由
type LoginRouter struct{}

// RegisterRoutes 实现 Router 接口
func (r LoginRouter) RegisterRoutes(routerGroup *gin.RouterGroup) {
	var (
		login v1.LoginController
	)

	// 登录
	routerGroup.POST("/login", login.Login)
}
