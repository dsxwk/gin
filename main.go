package main

import (
	"fmt"
	"gin/config"
	"gin/router"
	"gin/utils"
	"github.com/gin-gonic/gin"
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

	config.Init()

	// 运行环境模式 debug模式, test测试模式, release生产模式, 默认是debug,根据当前配置文件读取
	gin.SetMode(config.Conf.App.Mode)

	if config.Conf.App.Env != "production" {
		// 开发环境和测试环境允许所有代理
		_ = r.SetTrustedProxies(nil)
	}

	// 设置 HTTP 请求处理文件上传时可以使用的最大内存为 90MB
	r.MaxMultipartMemory = 90 << 20

	// 加载路由
	router.LoadRouters(r)

	// 启动提示
	fmt.Printf("✅ 应用：%s\n", config.Conf.App.Name)
	fmt.Printf("🌍 环境：%s\n", config.GetString("app.env"))
	fmt.Printf("🚪 端口：%d\n", config.Conf.App.Port)
	fmt.Printf("🗄️ 数据库：%s\n", config.Conf.Mysql.Database)
	fmt.Println("✅  Gin server started successfully!")
	fmt.Println("✅  0.0.0.0:" + utils.IntToString(config.Conf.App.Port))
	fmt.Println("👉  Open Swagger: http://127.0.0.1:" + utils.IntToString(config.Conf.App.Port) + "/swagger/index.html")
	fmt.Println("👉  Test API: http://127.0.0.1:" + utils.IntToString(config.Conf.App.Port) + "/ping")

	_ = r.Run(":" + utils.IntToString(config.Conf.App.Port))
}
