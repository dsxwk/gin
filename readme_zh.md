## 中文 | [English](readme.md)

- [项目简介](#项目简介)
- [许可证](#许可证)
- [安装说明](#安装说明)
- [目录结构](#目录结构)
- [使用方法](#使用方法)
  - [启动服务](#启动服务)
    - [air热更新](#air热更新)
  - [配置文件](#配置文件)
    -[项目配置](#项目配置)
    -[热更新配置](#热更新配置)
  - [路由](#路由)
    - [路由创建帮助](#路由创建帮助)
    - [路由创建](#路由创建)
    - [路由列表](#路由列表)
  - [控制器](#控制器)
    - [控制器创建帮助](#控制器创建帮助)
    - [控制器创建](#控制器创建)
  - [表单验证](#表单验证)
    - [验证创建帮助](#验证创建帮助)
    - [验证创建](#验证创建)
    - [验证场景](#验证场景)
    - [验证规则](#验证规则)
    - [提示信息](#提示信息)
    - [翻译](#翻译)
    - [自定义验证](#自定义验证)
  - [模型](#模型)
    - [模型创建帮助](#模型创建帮助)
    - [模型创建](#模型创建)
  - [服务](#服务)
    - [服务创建帮助](#服务创建帮助)
    - [服务创建](#服务创建)
  - [命令行](#命令行)
    - [简介](#简介)
    - [命令帮助](#命令帮助)
    - [编写命令](#编写命令)
    - [创建命令帮助](#创建命令帮助)
    - [创建命令](#创建命令)
    - [命令结构](#命令结构)
    - [选项参数](#选项参数)
    - [注册命令](#注册命令)
    - [执行命令](#执行命令)
  - [缓存](#缓存)
    - [全局缓存](#全局缓存)
    - [Redis缓存](#Redis缓存)
    - [内存缓存](#内存缓存)
    - [磁盘缓存](#磁盘缓存)
  - [事件](#事件)
  - [响应](#响应)
    - [成功响应](#成功响应)
      - [添加提示](#添加提示)
      - [添加数据](#添加数据)
    - [失败响应](#失败响应)
      - [添加提示](#添加提示)
      - [添加数据](#添加数据)
  - [错误处理](#错误处理)
  - [日志](#日志)
    - [错误调试](#错误调试)
  - [多语言](#多语言)
  - [swagger文档](#swagger文档)

# 项目简介
> 基于Golang语言框架Go Gin开发的轻量级框架，开箱即用，设计灵感基于laravel、thinphp等主流php框架，项目架构目录层次分明，初学者的福音，框架使用JWT、中间件、缓存、验证器、事件、路由、redis、命令行工具等、支持多语言,开发简单易于上手, 方便扩展。
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

# 许可证
- 📘 开源版: 遵循 AGPL-3.0，仅供学习、研究及非商业用途。
- 💼 商业版: 如需闭源或商业使用，请联系作者📧  [25076778@qq.com] 获取商业授权。

# 安装说明
> 项目基于Golang 1.25.2版本开发, 低版本可能存在版本差异, 建议版本 >= 1.25.2。
## 克隆项目
```bash
$ git clone https://github.com/dsxwk/gin.git
$ cd gin
```
## 初始化Go环境与依赖
### 方式一
```bash
$ go env -w GOPROXY=https://goproxy.cn,direct
$ go generate ./...
```
### 方式二
```bash
$ go env -w GO111MODULE=on
$ go env -w GOPROXY=https://goproxy.cn,direct
$ go get -u
$ go mod tidy
$ go mod download
$ go mod vendor
```
## 启动
```bash
$ go run main.go
```
### 使用air热更新
```bash
$ go install github.com/cosmtrek/air@latest
$ air
```

## 编译
```bash
$ go build main.go
```
### 运行
```bash
$  ./main
```

# 目录结构
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
│   ├──├── cache                        # 缓存
│   ├──├── i18n                         # 多语言
│   ├──├──├── locales                   # 翻译文件
├── vendor                              # 依赖包
├── .air.linux.toml                     # air配置文件
├── .air.toml                           # air配置文件
├── .gitignore                          # git忽略文件
├── cli.go                              # 命令行入口文件
├── config.yaml                         # 默认配置文件
├── dev.config.yaml                     # 本地环境配置文件
├── go.mod                              # go mod
├── LICENSE                             # 开源协议
├── main.go                             # 入口文件
├── readme.md                           # 英文文档
└── readme_zh.md                        # 中文文档
```

# 使用方法
## 启动服务
```bash
$ go run main.go
```
### air热更新
```bash
$ go install github.com/cosmtrek/air@latest
$ air
```

## 配置文件
### 项目配置
> `config.yaml`为默认配置文件, 可自行修改。`dev.config.yaml`对应本地环境配置, 通过以下app.env文件配置环境变量来切换环境
> ```
> app:
>   env: dev # dev|testing|production dev=本地环境 testing=测试环境 production=生产环境
> ```

### 热更新配置
> `.air.toml`为Windows环境下默认配置文件, `.air.linux.toml`为Linux环境下默认配置文件。可自行根据项目整体需要自行修改。

## 路由
> `router/root.go` 文件中定义了全局路由规则可自行修改,  一般情况只需要默认即可。
### 路由创建帮助
```bash
$ go run cli.go make:router -h # --help
make:router - 路由创建

Options:
  -f, --file  文件路径, 如: user      required:true
  -d, --desc  路由描述, 如: 用户路由   required:false
```

### 路由创建
```bash
$ go run cli.go make:router --file=user --desc=用户路由
```
```go
package router

import (
	"gin/app/controller/v1"
	"github.com/gin-gonic/gin"
)

// UserRouter 用户路由
type UserRouter struct{}

func init() {
	Register(&UserRouter{})
}

// RegisterRoutes 注册路由
func (r *UserRouter) RegisterRoutes(routerGroup *gin.RouterGroup) {
	var (
		user v1.UserController
	)

	router := routerGroup.Group("api/v1")
	{
		// 列表
		router.GET("/user", user.List)
		// 创建
		router.POST("/user", user.Create)
		// 更新
		router.PUT("/user/:id", user.Update)
		// 删除
		router.DELETE("/user/:id", user.Delete)
		// 详情
		router.GET("/user/:id", user.Detail)
	}
}

// IsAuth 是否需要鉴权
func (r *UserRouter) IsAuth() bool {
	return true
}

```

### 路由列表
```bash
$ go run cli.go route:list
---------------------------------------------------------
Method   Path                                Handler
---------------------------------------------------------
POST     /api/v1/login                       gin/app/controller/v1.(*LoginController).Login
GET      /api/v1/user                        gin/app/controller/v1.(*UserController).List
POST     /api/v1/user                        gin/app/controller/v1.(*UserController).Create
GET      /api/v1/user/:id                    gin/app/controller/v1.(*UserController).Detail
PUT      /api/v1/user/:id                    gin/app/controller/v1.(*UserController).Update
DELETE   /api/v1/user/:id                    gin/app/controller/v1.(*UserController).Delete
GET      /ping                               gin/router.LoadRouters
GET      /public/*filepath                   github.com/gin-gonic/gin.(*RouterGroup).createStaticHandler
HEAD     /public/*filepath                   github.com/gin-gonic/gin.(*RouterGroup).createStaticHandler
GET      /swagger/*any                       github.com/swaggo/gin-swagger.CustomWrapHandler
---------------------------------------------------------
总计 10 条路由
```

## 控制器
### 控制器创建帮助
```bash
$ go run cli.go make:controller -h # --help
make:controller - 控制器创建

Options:
  -f, --file      文件路径, 如: v1/user  required:true
  -F, --function  方法名称, 如: list     required:false
  -m, --method    请求方式, 如: get      required:false
  -r, --router    路由地址, 如: /user    required:false
  -d, --desc      描述, 如: 列表         required:false
```

### 控制器创建
```bash
$ go run cli.go make:controller --file=v1/test --router=/test --method=get --desc=列表 --function=list
```
```go
package v1

import (
    "gin/common/base"
    "gin/common/errcode"
    "github.com/gin-gonic/gin"
)

type TestController struct {
    base.BaseController
}

// List 列表
// @Router /test [get]
func (s *TestController) List(c *gin.Context) {
    // Define your function here
    s.Success(c, errcode.Success().WithMsg("Test Msg").WithData([]string{}))
}
```
