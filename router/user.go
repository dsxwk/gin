package router

import (
	"gin/app/controller/v1"
	"github.com/gin-gonic/gin"
)

// UserRouter 用户路由
type UserRouter struct{}

func init() {
	Register(&UserRouter{})
}

// RegisterRoutes 注册路由
func (r *UserRouter) RegisterRoutes(routerGroup *gin.RouterGroup) {
	var (
		user v1.UserController
	)

	router := routerGroup.Group("api/v1")
	{
		// 列表
		router.POST("/user", user.List)
	}
}

// IsAuth 是否需要鉴权
func (r *UserRouter) IsAuth() bool {
	return true
}
