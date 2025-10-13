package main

import (
	"gin/app/command"
	"gin/utils/cli"
	"gin/utils/cli/make"
	"gin/utils/cli/route"
)

func main() {
	// 注册命令
	cli.Register(&make.MakeCommand{})    // 命令行创建
	cli.Register(&make.MakeController{}) // 控制器创建
	cli.Register(&make.MakeService{})    // 服务创建
	cli.Register(&make.MakeRequest{})    // 验证请求创建
	cli.Register(&make.MakeMiddleware{}) // 中间件创建
	cli.Register(&make.MakeRouter{})     // 路由创建
	cli.Register(&route.RouteList{})     // 路由列表
	cli.Register(&command.DemoCommand{}) // 命令行demo
	// ... 注册其他命令

	// 执行命令
	cli.Execute()
}
