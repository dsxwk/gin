package {{.Package}}

import (
    {{- if ne .Package "router" }}
    "gin/router"
    {{- end }}
	"github.com/gin-gonic/gin"
)

// {{.Name}}Router {{.Description}}
type {{.Name}}Router struct {}

func init() {
	{{- if eq .Package "router" }}
	Register(&{{.Name}}Router{})
	{{- else }}
	router.Register(&{{.Name}}Router{})
	{{- end }}
}

// RegisterRoutes 注册路由
func (r *{{.Name}}Router) RegisterRoutes(routerGroup *gin.RouterGroup) {
	// var (
    //	 login v1.LoginController
    // )
    //
    // router := routerGroup.Group("api/v1")
    // {
    //	 // 登录
    //	 router.POST("/login", login.Login)
    //   // ... 其他路由
    // }
	// todo Define your route here ...
}

// IsAuth 是否需要鉴权
func (r *{{.Name}}Router) IsAuth() bool {
	return true
}
