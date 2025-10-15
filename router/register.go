package router

import "github.com/gin-gonic/gin"

// Router 路由接口
type Router interface {
	RegisterRoutes(router *gin.RouterGroup)
	IsAuth() bool // 是否需要鉴权
}

var routerRegistry []Router

// Register 注册路由模块
func Register(r Router) {
	routerRegistry = append(routerRegistry, r)
}

// AutoLoads 自动注册
func AutoLoads(public *gin.RouterGroup, auth *gin.RouterGroup) {
	for _, r := range routerRegistry {
		if r.IsAuth() {
			r.RegisterRoutes(auth)
		} else {
			r.RegisterRoutes(public)
		}
	}
}
