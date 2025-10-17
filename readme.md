## English | [中文](readme_zh.md)

## License Description
- 📘 Open Source Version: Following AGPL-3.0, for learning, research, and non-commercial purposes only.
- 💼 Commercial Version: If closed source or commercial use is required, please contact the author 📧   [ 25076778@qq.com ]Obtain commercial authorization.

## Project Address
- Github: https://github.com/dsxwk/gin.git
- Gitee: https://gitee.com/dsxwk/gin.git

## Introduction to Gin Framework
> Gin is a web framework written in Go language. It has the characteristics of simplicity, speed, and efficiency, and is widely used in Go language web development.

## Features of Gin Framework
- Fast: The Gin framework is based on the standard library net/http, using goroutines and channels to implement asynchronous processing and improve performance.
- Simple: The Gin framework provides a range of APIs and middleware, enabling developers to quickly build web applications.
- Efficient: The Gin framework uses sync. Pool to cache objects, reducing memory allocation and release, and improving performance.
> Golang Gin is a lightweight and efficient Golang web framework. It has the characteristics of high performance, ease of use, and flexibility, and is widely used in the development of various web applications.

## Introduction to Gin Project
### Command Line Creation
- Command
- Model
- Controller
- Service
- Validator
    - Customize Verification Scenarios
- Middleware
- Router
    
### Middleware
- Cors
- Language
- JWT
- Log

### Cache
- Memory Cache
- Redis Cache
- Disk Cache

### Event
- Publish
- Subscribe
- Unsubscribe
- Listener
    - Mysql
    - Http

### Log
- Error Information
- Stack Information
- Sql Statements
- Http Requests

### Language
> At present, only login related modules have cases, supporting only Chinese and English. If you need other modules or languages, please expand them yourself

## Tech Stack
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

## Project Structure
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
│   ├── extend                          # Extend
│   ├──├── cache                        # Cache
│   ├──├── i18n                         # Language
│   ├──├──├── locales                   # Translation
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
├── vendor                              # Vendor
├── .air.linux.toml                     # Air Configuration File
├── .air.toml                           # Air Configuration File
├── .gitignore                          # Gitignore
├── cli.go                              # Command Entry File
├── go.mod                              # go mod
├── LICENSE                             # LICENSE
├── LICENSE.COMMERCIAL                  # Business Agreement
├── main.go                             # Entry File
├── readme.md                           # English Document
└── readme_zh.md                        # Chinese Document
```

## Command
```shell
# Get Help
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

# Format Output
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

### Create Command
```shell
# Get Help
go run cli.go make:command -h # go run cli.go make:command --help
# Create
go run cli.go make:command --file=demo --name=demo-command --desc=command-desc
```
> Note: The created -- file is a subdirectories, for example: test/demo, the subdirectories need to be imported in cli.go
```go
import (
	_ "gin/command/test"
)
```

### Execute Command
```shell
# go run cli.go demo-command # Execute command demo-command corresponds to the custom name of the command line file
E:\www\dsx\www-go\gin [master]> go run cli.go demo-command       
❌  Parameter --args cannot be empty
Example: go run cli.go demo-command --args=arg1 --desc=command-desc
Helper: go run cli.go demo-command --help
exit status 1
E:\www\dsx\www-go\gin [master]> go run cli.go demo-command -h

demo-command - command-desc

Options:
  -a, --args  Example parameter, such as arg1 (arg1 is required)
E:\www\dsx\www-go\gin [master]> go run cli.go demo-command -a=111
Execute Command: demo-command, Parameter: 111
```

### Create Controller
```shell
# Get Help
go run cli.go make:controller -h # go run cli.go make:controller --help
# Create
go run cli.go make:controller --file=v1/user --function=list --method=get --router=/user --desc=list 
```

### Create Service
```shell
# Get Help
go run cli.go make:service -h # go run cli.go make:service --help 
# Create
go run cli.go make:service --file=v1/user --function=list --desc=list
```

### Create Middleware
```shell
# Get Help
go run cli.go make:middleware -h # go run cli.go make:middleware --help
# Create
go run cli.go make:middleware --file=jwt --desc=jwt-middleware
```

### Create Validator
```shell
# Get Help
go run cli.go make:request -h # go run cli.go make:request --help 
# Create
go run cli.go make:request --file=v1/user --desc=user-validator
```

### Create Router
```shell
# Get Help
go run cli.go make:router -h # go run cli.go make:router --help
# Create
go run cli.go make:router --file=user --desc=user-router
```

### Route List
```shell
E:\www\dsx\www-go\gin [master]> go run cli.go route:list
📦 Currently registered route
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
✅ A total of 6 routes
```

## Create Model
```shell
# Get Help
go run cli.go make:model -h # go run cli.go make:model --help 
# Create
go run cli.go make:model --table=user
```

## Air
### Install
```shell
go get github.com/air-go/air
# Run
air
```

## Swagger Document Generation
```bash
# Install
go install github.com/swaggo/swag/cmd/swag@latest
# Quickly generate document command
swag init -g main.go --exclude cli,app/service
```

## Test Case
```bash
go test -v ./tests/
```