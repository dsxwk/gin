package route

import (
	"fmt"
	"gin/common/base"
	"gin/router"
	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
	"strings"
)

type RouteList struct{}

func (s *RouteList) Name() string {
	return "route:list"
}

func (s *RouteList) Description() string {
	return "路由列表"
}

func (s *RouteList) Help() []base.CommandOption {
	return []base.CommandOption{}
}

func (s *RouteList) Execute(args []string) {
	// 初始化Gin引擎(不要Run)
	gin.SetMode(gin.ReleaseMode)
	engine := gin.New()

	// 加载项目路由
	router.LoadRouters(engine)

	// 获取所有路由
	routes := engine.Routes()

	// 打印路由列表
	fmt.Println("📦 当前已注册路由")
	fmt.Println("---------------------------------------------------------")
	fmt.Printf("%-8s %-35s %-40s\n", "Method", "Path", "Handler")
	fmt.Println("---------------------------------------------------------")

	for _, route := range routes {
		color.Green(
			"%-8s %-35s %-40s\n",
			route.Method,
			route.Path,
			s.formatHandlerName(route.Handler),
		)
	}

	fmt.Println("---------------------------------------------------------")
	fmt.Printf("✅ 总计 %d 条路由\n", len(routes))
}

func (s *RouteList) formatHandlerName(handler string) string {
	// 去掉 -fm 结尾
	handler = strings.TrimSuffix(handler, "-fm")
	// 去掉 .func1
	handler = strings.TrimSuffix(handler, ".func1")

	return handler
}
