## English | [中文](readme_zh.md)

- [Project Introduction](#Project-Introduction)
- [License](#License)
- [Installation Instructions](#Installation-Instructions)
- [Directory Structure](#Directory-Structure)
- [Instructions For Use](#Instructions-For-Use)
  - [Start Service](#Start-Service)
    - [Air Hot Update](#Air-Hot-Update)
  - [Configuration File](#Configuration-File)
    -[Project Configuration](#Project-Configuration)
    -[Hot Update Configuration](#Hot-Update-Configuration)
  - [Route](#Route)
    - [Route Creation Help](#Route-Creation-Help)
    - [Route Creation](#Route-Creation)
    - [Route List](#Route-List)
  - [Controller](#Controller)
    - [Controller Creation Help](#Controller-Creation-Help)
    - [Controller Creation](#Controller-Creation)
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

# Project Introduction
> A lightweight framework developed based on the Golang language framework Go Gin, out of the box, inspired by mainstream PHP frameworks such as Laravel and ThinPHP. The project architecture directory has a clear hierarchy, which is a blessing for beginners. The framework uses JWT, middleware, cache, validator, event, routing, etc redis、 Command line tools and other technologies. support multiple languages, simple to develop and easy to use, convenient for extension.
## Project Address
- Github: https://github.com/dsxwk/gin.git
- Gitee: https://gitee.com/dsxwk/gin.git

## Introduction to the Gin Framework
> Gin is a web framework written in Go language. It has the characteristics of simplicity, speed, and efficiency, and is widely used in Go language web development.

## Features of Gin Framework
- Fast: The Gin framework is based on the standard library net/http, using goroutines and channels to implement asynchronous processing and improve performance.
- Simple: The Gin framework provides a range of APIs and middleware, enabling developers to quickly build web applications.
- Efficient: The Gin framework uses sync. Pool to cache objects, reducing memory allocation and release, and improving performance.
> Golang Gin is a lightweight and efficient Golang web framework. It has the characteristics of high performance, ease of use, and flexibility, and is widely used in the development of various web applications.

# License
- 📘 Open source version: Following AGPL-3.0, for learning, research, and non-commercial use only.
- 💼 Commercial version: If closed source or commercial use is required, please contact the author 📧   [ 25076778@qq.com ]Obtain commercial authorization.

# Installation Instructions
> The project is developed based on Golang version 1.25.2, and there may be version differences in lower versions. It is recommended that the version be greater than or equal to 1.25.2.
## Clone Project
```bash
$ git clone https://github.com/dsxwk/gin.git
$ cd gin
```
## Initialize Go environment and dependencies
### Method One
```bash
$ go env -w GOPROXY=https://goproxy.cn,direct
$ go generate ./...
```
### Method Two
```bash
$ go env -w GO111MODULE=on
$ go env -w GOPROXY=https://goproxy.cn,direct
$ go get -u
$ go mod tidy
$ go mod download
$ go mod vendor
```
## Start
```bash
$ go run main.go
```
### Use Air Hot Update
```bash
$ go install github.com/cosmtrek/air@latest
$ air
```

## Compile
```bash
$ go build main.go
```
### Run
```bash
$ ./main
```

# Directory Structure
```
├── app                                 # Application
│   ├── command                         # Command
│   ├── controller                      # Controller
│   ├── middleware                      # Middleware
│   ├── model                           # Model
│   ├── request                         # Validator
│   ├── service                         # Service
├── common                              # Common Module
│   ├── base                            # Base
│   ├── errcode                         # Errcode
│   ├── response                        # Response
│   ├── global                          # Global Variable
│   ├── template                        # Template
├── config                              # Config File
├── database                            # Database Test File 
├── docs                                # Swagger Doc
├── public                              # Static Resources
├── router                              # Router
├── storage                             # Storage
│   ├── cache                           # Disk Cache
│   ├── logs                            # Logs
├── tests                               # Test Case
├── utils                               # Utils
│   ├──├── cache                        # Cache
│   ├──├── i18n                         # Language
│   ├──├──├── locales                   # Translation
├── vendor                              # Vendor
├── .air.linux.toml                     # Air Configuration File
├── .air.toml                           # Air Configuration File
├── .gitignore                          # Gitignore
├── cli.go                              # Command Entry File
├── config.yaml                         # Default Configuration File
├── dev.config.yaml                     # Local Environment Configuration File
├── go.mod                              # go mod
├── LICENSE                             # LICENSE
├── main.go                             # Entry File
├── readme.md                           # English Document
└── readme_zh.md                        # Chinese Document
```

# Instructions For Use
## Start Service
```bash
$ go run main.go
```
### Air Hot Update
```bash
$ go install github.com/cosmtrek/air@latest
$ air
```

## Configuration File
### Project Configuration
> `config.yaml` is the default configuration file and can be modified by oneself. `dev.config.yaml` corresponds to the local environment configuration, and environment variables can be configured through the following app.exe file to switch environments
> ```
> app:
>   env: dev # dev|testing|production dev=local-environment testing=test-environment production=production-environment
> ```

### Hot Update Configuration
> `.air.toml` is the default configuration file in Windows environment, and `.air.Linux.toml` is the default configuration file in Linux environment. You can modify it according to the overall needs of the project.

## Route
> The `router/root.go` file defines global routing rules that can be modified by oneself, and in general, they only need to be defaulted.
### Route Creation Help
```bash
$ go run cli.go make:router -h # --help
make:router - Route Creation

Options:
  -f, --file  File Path, Expample: user                   required:true
  -d, --desc  Route Description, Expample: User-Routing   required:false
```

### Route Creation
```bash
$ go run cli.go make:router --file=user --desc=User-Routing
```
```go
package router

import (
	"gin/app/controller/v1"
	"github.com/gin-gonic/gin"
)

// UserRouter User-Routing
type UserRouter struct{}

func init() {
	Register(&UserRouter{})
}

// RegisterRoutes Register Route
func (r *UserRouter) RegisterRoutes(routerGroup *gin.RouterGroup) {
	var (
		user v1.UserController
	)

	router := routerGroup.Group("api/v1")
	{
		// List
		router.GET("/user", user.List)
		// Create
		router.POST("/user", user.Create)
		// Update
		router.PUT("/user/:id", user.Update)
		// Delete
		router.DELETE("/user/:id", user.Delete)
		// Detail
		router.GET("/user/:id", user.Detail)
	}
}

// IsAuth Do you need authentication
func (r *UserRouter) IsAuth() bool {
	return true
}

```

### Route List
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
A total of 10 routes
```

## Controller
### Controller Creation Help
```bash
$ go run cli.go make:controller -h # --help
make:controller - Controller Creation

Options:
  -f, --file      File Path, Example: v1/user       required:true
  -F, --function  Function Name, Example: list      required:false
  -m, --method    Request Method, Example: get      required:false
  -r, --router    Route Adress, Example: /user      required:false
  -d, --desc      Description, Example: Test-List   required:false
```

### Controller Creation
```bash
$ go run cli.go make:controller --file=v1/test --router=/test --method=get --desc=Test-List --function=list
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

// List Test-List
// @Router /test [get]
func (s *TestController) List(c *gin.Context) {
    // Define your function here
    s.Success(c, errcode.Success().WithMsg("Test Msg").WithData([]string{}))
}
```
