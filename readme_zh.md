## 中文 | [English](readme.md)

## 许可说明
- 📘 开源版: 遵循 AGPL-3.0，仅供学习、研究及非商业用途。
- 💼 商业版: 如需闭源或商业使用，请联系作者📧  [25076778@qq.com] 获取商业授权。

## 项目地址
- Github: https://github.com/dsxwk/gin.git
- Gitee: https://gitee.com/dsxwk/gin.git

## Gin框架介绍
> Gin是一个用Go语言编写的Web框架。它具有简单、快速、高效等特点，被广泛应用于Go语言的Web开发中。

## Gin框架的特性
- 快速: Gin框架基于标准库net/http，使用goroutine和channel实现异步处理，提高性能。
- 简单: Gin框架提供了一系列的API和中间件，使得开发人员可以快速构建Web应用程序。
- 高效: Gin框架使用sync.Pool来缓存对象，减少内存分配和释放，提高性能。
> Golang Gin 是一个轻量级且高效的 Golang Web 框架。它具有高性能、易用性和灵活性等特点，被广泛应用于各种 Web 应用程序的开发。

## Gin项目介绍
### 命令行创建
- 命令
- 模型
- 控制器
- 服务
- 验证器
  - 自定义验证场景
- 中间件
- 路由
  
### 中间件
- 跨域
- 多语言
- JWT
- 日志

### 缓存
- 内存缓存
- redis缓存
- 磁盘缓存

### 事件
- 发布
- 订阅
- 取消订阅
- 事件监听
  - Mysql
  - Http

### 日志
- 错误信息
- 堆栈信息
- Sql语句
- Http请求

### 多语言
> 目前只有登录相关模块有案例,只支持中英文,如需其他模块或语言自行扩展

## 技术栈
- Gin
- Gorm
- Jwt
- Mysql
- Middleware
- Validator
- Cache
- Event
- Viper
- Swagger
- Air
- ...

## 项目结构
```
├── app                                 # 应用程序
│   ├── command                         # 命令
│   ├── controller                      # 控制器
│   ├── middleware                      # 中间件
│   ├── model                           # 模型
│   ├── request                         # 验证器
│   ├── service                         # 服务
├── common                              # 公共模块
│   ├── base                            # 基类
│   ├── errcode                         # 错误码
│   ├── response                        # 响应
│   ├── extend                          # 扩展
│   ├──├── cache                        # 缓存
│   ├──├── i18n                         # 多语言
│   ├──├──├── locales                   # 翻译文件
│   ├── global                          # 全局变量
│   ├── template                        # 模版
├── config                              # 配置文件
├── database                            # 数据库测试文件
├── docs                                # 文档
├── public                              # 静态资源
├── router                              # 路由
├── storage                             # 存储
│   ├── cache                           # 磁盘缓存
│   ├── logs                            # 日志
├── tests                               # 测试用例
├── utils                               # 工具包
├── vendor                              # 依赖包
├── .air.linux.toml                     # air配置文件
├── .air.toml                           # air配置文件
├── .gitignore                          # git忽略文件
├── cli.go                              # 命令行入口文件
├── go.mod                              # go mod
├── LICENSE                             # 开源协议
├── LICENSE.COMMERCIAL                  # 商业协议
├── main.go                             # 入口文件
├── readme.md                           # 英文文档
└── readme_zh.md                        # 中文文档
```

## 命令行
```shell
# 获取帮助
E:\www\dsx\www-go\gin [master]> go run cli.go -h # go run cli.go --help
Usage: go run cli.go [command] [options]
Available commands:
  demo-command             test-demo
  make:command             服务创建
  make:controller          控制器创建
  make:middleware          中间件创建
  make:model               模型创建
  make:request             验证请求创建
  make:router              路由创建
  make:service             服务创建
  route:list               路由列表

Options:
  -f, --format   The output format (txt, json) [default: "txt"]
  -h, --help     Display help for the given command. When no command is given display help for the list command
  -v, --version  Display this application version

# 格式化输出
E:\www\dsx\www-go\gin [master]> go run cli.go -f=json
{
  "commands": [
    {
      "description": "test-demo",
      "name": "demo-command"
    },
    {
      "description": "服务创建",
      "name": "make:command"
    },
    {
      "description": "控制器创建",
      "name": "make:controller"
    },
    {
      "description": "中间件创建",
      "name": "make:middleware"
    },
    {
      "description": "模型创建",
      "name": "make:model"
    },
    {
      "description": "验证请求创建",
      "name": "make:request"
    },
    {
      "description": "路由创建",
      "name": "make:router"
    },
    {
      "description": "服务创建",
      "name": "make:service"
    },
    {
      "description": "路由列表",
      "name": "route:list"
    }
  ],
  "version": "Gin CLI v1.0.0"
}
```

### 命令创建
```shell
# 获取帮助
go run cli.go make:command -h # go run cli.go make:command --help
# 创建
go run cli.go make:command --file=demo --name=demo-command --desc=command-desc
```
> 注意: 创建的--file为子目录,如: test/demo 则需要在cli.go 导入子目录
```go
import (
	_ "gin/command/test"
)
```

### 执行命令
```shell
# go run cli.go demo-command # 执行命令 demo-command对应命令行文件自定义的name
E:\www\dsx\www-go\gin [master]> go run .\cli.go demo-command                                   
❌  参数 --args 不能为空
Example: go run cli.go Demo:command --args=arg1 --desc=test-demo
Helper: go run cli.go Demo:command --help
exit status 1
E:\www\dsx\www-go\gin [master]> go run .\cli.go demo-command -h

demo-command - command-desc

Options:
  -a, --args  示例参数, 如: arg1 (参数1必填)

E:\www\dsx\www-go\gin [master]> go run cli.go demo-command -a=111
执行命令: demo-command, 参数: 111
```

### 控制器创建
```shell
# 获取帮助
go run cli.go make:controller -h # go run cli.go make:controller --help 
# 创建
go run cli.go make:controller --file=v1/user --function=list --method=get --router=/user --desc=列表 
```

### 服务创建
```shell
# 获取帮助
go run cli.go make:service -h # go run cli.go make:service --help 
# 创建
go run cli.go make:service --file=v1/user --function=list --desc=列表
```

### 中间件创建
```shell
# 获取帮助
go run cli.go make:middleware -h # go run cli.go make:middleware --help 
# 创建
go run cli.go make:middleware --file=jwt --desc=jwt中间件
```

### 验证器创建
```shell
# 获取帮助
go run cli.go make:request -h # go run cli.go make:request --help 
# 创建
go run cli.go make:request --file=v1/user --desc=用户请求验证
```

### 路由创建
```shell
# 获取帮助
go run cli.go make:router -h # go run cli.go make:router --help 
# 创建
go run cli.go make:router --file=user --desc=用户路由
```

### 路由列表
```shell
E:\www\dsx\www-go\gin [master]> go run cli.go route:list
📦 当前已注册路由
---------------------------------------------------------
Method   Path                                Handler
---------------------------------------------------------
GET      /ping                               gin/router.LoadRouters
GET      /public/*filepath                   github.com/gin-gonic/gin.(*RouterGroup).createStaticHandler
GET      /swagger/*any                       github.com/swaggo/gin-swagger.CustomWrapHandler
HEAD     /public/*filepath                   github.com/gin-gonic/gin.(*RouterGroup).createStaticHandler
POST     /api/v1/login                       gin/app/controller/v1.(*LoginController).Login
POST     /api/v1/user                        gin/app/controller/v1.(*UserController).List
---------------------------------------------------------
✅ 总计 6 条路由
```

## 模型创建
```shell
# 获取帮助
go run cli.go make:model -h # go run cli.go make:model --help 
# 创建
go run cli.go make:model --table=user
```

## Air
### 安装
```shell
D:\www\dsx\go\gin [master]> go install github.com/air-verse/air@latest
# 运行
D:\www\dsx\go\gin [master]> air

  __    _   ___
 / /\  | | | |_)
/_/--\ |_| |_| \_ v1.63.0, built with Go go1.25.2

watching .
watching app
watching app\command
watching app\controller
watching app\controller\v1
watching app\middleware
watching app\model
watching app\request
watching app\service
watching common
watching common\base
watching common\errcode
watching common\global
watching common\response
watching common\template
watching config
watching database
watching docs
watching public
watching router
watching tests
!exclude tmp
watching utils
watching utils\cli
watching utils\cli\make
watching utils\cli\route
!exclude vendor
building...
running...
✅ 已加载基础配置文件: D:\www\dsx\go\gin\config\config.yaml
✅ 已加载环境配置文件: config\dev.config.yaml
[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

✅ 已加载基础配置文件: D:\www\dsx\go\gin\config\config.yaml
✅ 已加载环境配置文件: config\dev.config.yaml
[GIN-debug] GET    /ping                     --> gin/router.LoadRouters.func1 (3 handlers)
[GIN-debug] GET    /public/*filepath         --> github.com/gin-gonic/gin.(*RouterGroup).createStaticHandler.func1 (3 handlers)
[GIN-debug] HEAD   /public/*filepath         --> github.com/gin-gonic/gin.(*RouterGroup).createStaticHandler.func1 (3 handlers)
[GIN-debug] GET    /swagger/*any             --> github.com/swaggo/gin-swagger.CustomWrapHandler.func1 (3 handlers)
[GIN-debug] POST   /api/v1/login             --> gin/app/controller/v1.(*LoginController).Login-fm (3 handlers)
[GIN-debug] GET    /api/v1/user              --> gin/app/controller/v1.(*UserController).List-fm (4 handlers)
[GIN-debug] POST   /api/v1/user              --> gin/app/controller/v1.(*UserController).Create-fm (4 handlers)
[GIN-debug] PUT    /api/v1/user/:id          --> gin/app/controller/v1.(*UserController).Update-fm (4 handlers)
[GIN-debug] DELETE /api/v1/user/:id          --> gin/app/controller/v1.(*UserController).Delete-fm (4 handlers)
[GIN-debug] GET    /api/v1/user/:id          --> gin/app/controller/v1.(*UserController).Detail-fm (4 handlers)
端口:    8080
数据库:  gin
应用:    gin
环境:    dev
✅  Gin server started successfully!
✅  0.0.0.0:8080
👉 Open Swagger: http://127.0.0.1:8080/swagger/index.html
👉 Test API: http://127.0.0.1:8080/ping
[GIN-debug] Listening and serving HTTP on :8080
```

## Swagger 文档生成
```bash
# 安装
D:\www\dsx\go\gin [master]> go install github.com/swaggo/swag/cmd/swag@latest
# 快速生成文档命令
D:\www\dsx\go\gin [master]> swag init -g main.go --exclude cli,app/service
2025/10/18 21:32:48 Generate swagger docs....
2025/10/18 21:32:48 Generate general API Info, search dir:./
2025/10/18 21:32:49 Generating request.UserLogin
2025/10/18 21:32:49 Generating errcode.SuccessResponse
2025/10/18 21:32:49 Generating v1.LoginResponse
2025/10/18 21:32:49 Generating v1.Token
2025/10/18 21:32:49 Generating model.User
2025/10/18 21:32:49 Generating model.JsonTime
2025/10/18 21:32:49 Generating errcode.ArgsErrorResponse
2025/10/18 21:32:49 Generating errcode.SystemErrorResponse
2025/10/18 21:32:49 Generating request.PageData
2025/10/18 21:32:49 Generating request.UserCreate
2025/10/18 21:32:49 Generating request.UserUpdate
2025/10/18 21:32:49 Generating request.UserDetail
2025/10/18 21:32:49 create docs.go at docs/docs.go
2025/10/18 21:32:49 create swagger.json at docs/swagger.json
2025/10/18 21:32:49 create swagger.yaml at docs/swagger.yaml
```

## 测试用例
```bash
D:\www\dsx\go\gin [master]> go test -v ./tests/
```