package main

import (
	"context"
	"fmt"
	"gin/config"
	"gin/router"
	"gin/utils"
	"github.com/gin-gonic/gin"
	"github.com/mattn/go-runewidth"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
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

	data := map[string]interface{}{
		"应用":  config.Conf.App.Name,
		"环境":  config.GetString("app.env"),
		"端口":  config.Conf.App.Port,
		"数据库": config.Conf.Mysql.Database,
	}

	// 启动提示
	PrintAligned(data)
	fmt.Println("✅  Gin server started successfully!")
	fmt.Println("✅  0.0.0.0:" + utils.IntToString(config.Conf.App.Port))
	fmt.Println("👉 Open Swagger: http://127.0.0.1:" + utils.IntToString(config.Conf.App.Port) + "/swagger/index.html")
	fmt.Println("👉 Test API: http://127.0.0.1:" + utils.IntToString(config.Conf.App.Port) + "/ping")

	srv := &http.Server{
		Addr:         ":" + utils.IntToString(config.Conf.App.Port),
		Handler:      r,
		ReadTimeout:  3 * time.Second,  // 设置读取超时
		WriteTimeout: 3 * time.Second,  // 设置写入超时
		IdleTimeout:  30 * time.Second, // 设置空闲超时
	}

	go func() {
		_ = srv.ListenAndServe()
	}()

	// 等待中断信号以优雅地关闭服务器（设置5秒的超时时间）
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_ = srv.Shutdown(ctx)
}

// PrintAligned 打印冒号对齐,支持中文
func PrintAligned(data map[string]interface{}) {
	// 找出最长key的显示宽度
	maxLen := 0
	for k := range data {
		w := runewidth.StringWidth(k)
		if w > maxLen {
			maxLen = w
		}
	}

	// 打印
	for k, v := range data {
		padding := maxLen - runewidth.StringWidth(k) + 2
		fmt.Printf("%s:%s%v\n", k, spaces(padding), v)
	}
}

// spaces 生成n个空格
func spaces(n int) string {
	if n <= 0 {
		return ""
	}
	return fmt.Sprintf("%*s", n, "")
}
