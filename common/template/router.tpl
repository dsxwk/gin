package router

import (
	"github.com/gin-gonic/gin"
)

type {{.Name}}Router struct {}

// RegisterRoutes 实现 Router 接口
func (r {{.Name}}Router) RegisterRoutes(routerGroup *gin.RouterGroup) {
	// Define your routes here ...
    // routerGroup.GET("your-route" ...)
}
