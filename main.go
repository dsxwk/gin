package main

import (
	"fmt"
	"gin/router"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @title Gin Swagger API
// @version 2.0
// @description Gin API 文档
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email 25076778@qq.com
// @host 127.0.0.1:8080
func main() {
	var (
		r = gin.Default()
	)

	// 运行环境模式 debug模式, test测试模式, release生产模式, 默认是debug,根据当前配置文件读取
	gin.SetMode("debug")

	// 静态文件
	r.StaticFS("/public", http.Dir("./public"))

	// 设置 HTTP 请求处理文件上传时可以使用的最大内存为 90MB
	r.MaxMultipartMemory = 90 << 20

	// 加载路由
	router.LoadRouters(r)

	// 启动提示
	fmt.Println("✅  Gin server started successfully!")
	fmt.Println("✅  0.0.0.0:8080")
	fmt.Println("👉  Open Swagger: http://127.0.0.1:8080/swagger/index.html")
	fmt.Println("👉  Test API: http://127.0.0.1:8080/ping")

	_ = r.Run(":8080")
}
