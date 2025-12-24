package main

import (
	"context"
	"fmt"
	_ "gin/app/listener"
	_ "gin/app/queue/kafka/consumer"
	_ "gin/app/queue/rabbitmq/consumer"
	"gin/config"
	"gin/router"
	"gin/utils"
	"gin/utils/debugger"
	"gin/utils/lang"
	"gin/utils/message"
	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
	"github.com/mattn/go-runewidth"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go get # -u
//go:generate go mod tidy
//go:generate go mod download
//go:generate go mod vendor

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
	gin.SetMode(config.Conf.App.Mode)

	if config.Conf.App.Env != "production" {
		// 开发环境和测试环境允许所有代理
		_ = r.SetTrustedProxies(nil)
	}

	// 设置 HTTP 请求处理文件上传时可以使用的最大内存为 90MB
	r.MaxMultipartMemory = 90 << 20

	// 加载路由
	router.LoadRouters(r)

	// 加载翻译
	lang.LoadLang()
	// debugger订阅
	dbg := debugger.NewDebugger(message.MsgEventBus)
	dbg.Start()
	defer dbg.Stop()

	data := map[string]interface{}{
		"应用":  config.Conf.App.Name,
		"环境":  config.GetString("app.env"),
		"端口":  config.Conf.App.Port,
		"数据库": config.Conf.Mysql.Database,
	}

	// 启动提示
	PrintAligned(data, []string{"应用", "环境", "端口", "数据库"})

	var port = utils.IntToString(config.Conf.App.Port)
	run := map[string]interface{}{
		utils.Network + " Address:":  "http://0.0.0.0:" + port,
		utils.Pointer + " Swagger:":  "http://127.0.0.1:" + port + "/swagger/index.html",
		utils.Pointer + " Test API:": "http://127.0.0.1:" + port + "/ping",
	}
	PrintAligned(run, []string{utils.Network + " Address:", utils.Pointer + " Swagger:", utils.Pointer + " Test API:"})
	fmt.Println("Gin server started successfully!")

	srv := &http.Server{
		Addr:              ":" + port,
		Handler:           r,
		ReadTimeout:       10 * time.Second, // 设置读取超时
		WriteTimeout:      10 * time.Second, // 设置写入超时
		IdleTimeout:       30 * time.Second, // 设置空闲超时
		ReadHeaderTimeout: 5 * time.Second,  // 设置读取头超时
	}

	go func() {
		err := srv.ListenAndServe()
		if err != nil {
			log.Fatal(err.Error())
		}
	}()

	// 等待中断信号以优雅地关闭服务器（设置5秒的超时时间）
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	color.Yellow("服务正在关闭...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		color.Red(utils.Error+" 服务关闭异常: %v", err)
	}

	select {}
}

// PrintAligned 打印冒号对齐,支持中文
func PrintAligned(data map[string]interface{}, order []string) {
	// 找出最长key的显示宽度
	maxLen := 0
	for k := range data {
		w := runewidth.StringWidth(k)
		if w > maxLen {
			maxLen = w
		}
	}

	for _, k := range order {
		key := ensureEmojiSpace(strings.TrimSuffix(k, ":"))
		padding := maxLen - runewidth.StringWidth(key) + 2
		fmt.Printf("%s:%s%v\n", key, spaces(padding), data[k])
	}
}

func ensureEmojiSpace(s string) string {
	r := []rune(s)
	if len(r) > 0 && (r[0] > 0x1F000 && r[0] < 0x1FAFF) {
		if len(r) > 1 && r[1] != ' ' {
			return string(r[0]) + " " + string(r[1:])
		}
	}
	return s
}

// spaces 生成n个空格
func spaces(n int) string {
	if n <= 0 {
		return ""
	}

	return fmt.Sprintf("%*s", n, "")
}
