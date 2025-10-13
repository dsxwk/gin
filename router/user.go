package router

import (
	"gin/app/controller/v1"
	"github.com/gin-gonic/gin"
)

// UserRouter 用户路由
type UserRouter struct{}

// RegisterRoutes 实现 Router 接口
func (r UserRouter) RegisterRoutes(routerGroup *gin.RouterGroup) {
	var (
		user v1.UserController
	)

	// 列表
	routerGroup.GET("/user", user.List)
}
