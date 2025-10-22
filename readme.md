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
D:\www\dsx\go\gin [master]> go install github.com/air-verse/air@latest
# Run
D:\www\dsx\go\gin [master]> air

  __    _   ___
 / /\  | | | |_)
/_/--\ |_| |_| \_ v1.62.0, built with Go go1.24.2

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
!exclude storage
watching tests
!exclude tmp
watching utils
watching utils\cli
watching utils\cli\db
watching utils\cli\make
watching utils\cli\route
watching utils\ctx
!exclude vendor
building...
running...
✅ 已加载环境配置文件: config\dev.config.yaml
[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

✅ 已加载环境配置文件: config\dev.config.yaml
[GIN-debug] GET    /ping                     --> gin/router.LoadRouters.func1 (3 handlers)
[GIN-debug] GET    /public/*filepath         --> github.com/gin-gonic/gin.(*RouterGroup).createStaticHandler.func1 (3 handlers)
[GIN-debug] HEAD   /public/*filepath         --> github.com/gin-gonic/gin.(*RouterGroup).createStaticHandler.func1 (3 handlers)
[GIN-debug] GET    /swagger/*any             --> github.com/swaggo/gin-swagger.CustomWrapHandler.func1 (3 handlers)
[GIN-debug] POST   /api/v1/login             --> gin/app/controller/v1.(*LoginController).Login-fm (4 handlers)
[GIN-debug] GET    /api/v1/user              --> gin/app/controller/v1.(*UserController).List-fm (5 handlers)
[GIN-debug] POST   /api/v1/user              --> gin/app/controller/v1.(*UserController).Create-fm (5 handlers)
[GIN-debug] PUT    /api/v1/user/:id          --> gin/app/controller/v1.(*UserController).Update-fm (5 handlers)
[GIN-debug] DELETE /api/v1/user/:id          --> gin/app/controller/v1.(*UserController).Delete-fm (5 handlers)
[GIN-debug] GET    /api/v1/user/:id          --> gin/app/controller/v1.(*UserController).Detail-fm (5 handlers)
应用:    gin
环境:    dev
端口:    8080
数据库:  gin
🌐 Address:    http://0.0.0.0:8080
👉 Swagger:    http://127.0.0.1:8080/swagger/index.html
👉 Test API:   http://127.0.0.1:8080/ping
✅  Success:   Gin server started successfully!
```

## Swagger Document Generation
```bash
# Install
D:\www\dsx\go\gin [master]> go install github.com/swaggo/swag/cmd/swag@latest
# Quickly generate document command
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

## Test Case
```bash
D:\www\dsx\go\gin [master]> go test -v ./tests/
```