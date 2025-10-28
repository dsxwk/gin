## English | [中文](readme_zh.md)

- [Project Introduction](#Project-Introduction)
- [License](#License)
- [Installation Instructions](#Installation-Instructions)
  -[Clone Project](#Clone-Project)
  -[Initialize Go Environment And Dependencies](#Initialize-Go-Environment-And-Dependencies)
    -[Method One](#Method-One)
    -[Method Two](#Method-Two)
  -[Start](#Start)
    -[Use Air Hot Update](#Use-Air-Hot-Update)
  -[Compile](#Compile)
    -[Compile Project](#Compile-Project)
    -[Compile Command](#Compile-Command)
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
  - [Model](#Model)
    - [Model Creation Help](#Model-Creation-Help)
    - [Model Creation](#Model-Creation)
  - [Form Validation](#Form-Validation)
    - [Validator Creation Help](#Validator-Creation-Help)
    - [Validator Creation](#Validator-Creation)
    - [Validator Rules](#Validator-Rules)
    - [Validator Scenes](#Validator-Scenes)
    - [Prompt Message](#Prompt-Message)
    - [Field Translation](#Field-Translation)
    - [Custom Validation](#Custom-Validation)
      - [Global Rules](#Global-Rules)
      - [Local Rules](#Local-Rules)
      - [Temporary Rules](#Temporary-Rules)
      - [Validator Usage](#Validator-Usage)
      - [Used In The Controller](#Used-In-The-Controller)
  - [Service](#Service)
    - [Service Creation Help](#Service-Creation-Help)
    - [Service Creation](#Service-Creation)
  - [Command](#Command)
    - [Get Version](#Get-Version)
    - [Command Help](#Command-Help)
    - [Command List](#Command-List)
    - [Command Creation Help](#Command-Creation-Help)
    - [Command Creation](#Command-Creation)
    - [Command Structure](#Command-Structure)
    - [Command Registration](#Command-Registration)
    - [Help Options](#Help-Options)
    - [Execute Command](#Execute-Command)
    - [Compile And Execute Commands](#Compile-And-Execute-Commands)
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
  - [Language Support](#Language-Support)
    - [Directory Configuration](#Directory-Configuration) 
    - [Ordinary Translation](#Ordinary-Translation) 
    - [Template Translation](#Template-Translation) 
    - [Add Language Support](#Add-Language-Support) 
  - [Swagger Documents](#Swagger-Documents)

# Project Introduction
> A lightweight framework developed based on the Golang language framework `Go Gin`, out of the box, inspired by mainstream PHP frameworks such as `Laravel` and `ThinkPHP`. The project architecture directory has a clear hierarchy, which is a blessing for beginners. The framework uses `JWT`, `middleware`, `cache`, `validator`, `event`, `routing`, `redis`、 `Command` and other technologies. support multiple languages, simple to develop and easy to use, convenient for extension.
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
## Initialize Go Environment And Dependencies
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
$ go install github.com/air-verse/air@latest
$ air
```

## Compile
### Compile Project
```bash
$ go build main.go
$ ./main
```

### Compile Command
```bash
$ go build cli.go
$ ./cli demo-command --args=11

Excute Command: demo-command, Argument: 11
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
│   ├── locales                         # Translation
│   ├──├── en                           # English Translation
│   ├──├── zh                           # Chinese Translation
├── tests                               # Test Case
├── utils                               # Utils
│   ├──├── cache                        # Cache
│   ├──├── cli                          # Command
│   ├──├── ctx                          # Context
│   ├──├── lang                         # Language
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
$ go install github.com/air-verse/air@latest
$ air

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

// RegisterRoutes Register-Route
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
## Model
### Model Creation Help
```bash
$ go run cli.go make:model -h # --help

make:model - Model Creation

Options:
  -t, --table  Table Name, Example: user or user,menu     required:true
  -p, --path   Output Directory, Example: api/user        required:false
  -c, --camel  Is it a camel hump field, Example: true    required:false
```

### Model Creation
> Support the creation of multiple model files simultaneously. If multiple model files need to be created, please separate the table name parameters of the descendants with commas, such as: user, menu
```bash
$ go run cli.go make:model --table=user,menu --path=api/user --camel=true
```
```go
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package user

import "gin/app/model"

const TableNameUser = "user"

// User User-Table
type User struct {
	ID        int64            `gorm:"column:id;type:int(10) unsigned;primaryKey;autoIncrement:true;comment:ID" json:"id"`                       // ID
	Avatar    string           `gorm:"column:avatar;type:varchar(255);not null;comment:avatar" json:"avatar"`                                    // avatar
	Username  string           `gorm:"column:username;type:varchar(10);not null;comment:username" json:"username"`                               // username
	FullName  string           `gorm:"column:full_name;type:varchar(20);not null;comment:fullname" json:"fullName"`                              // fullName
	Email     string           `gorm:"column:email;type:varchar(50);not null;comment:email" json:"email"`                                        // email
	Password  string           `gorm:"column:password;type:varchar(255);not null;comment:password" json:"password"`                              // password
	Nickname  string           `gorm:"column:nickname;type:varchar(50);not null;comment:nickname" json:"nickname"`                               // nickname
	Gender    int64            `gorm:"column:gender;type:tinyint(1) unsigned;not null;comment:gender 1=male 2=female" json:"gender"`             // gender 1=male 2=female
	Age       int64            `gorm:"column:age;type:int(10) unsigned;not null;comment:age" json:"age"`                                         // age
	Status    int64            `gorm:"column:status;type:tinyint(3) unsigned;not null;default:1;comment:state 1=enable 2=disable" json:"status"` // state 1=enable 2=disable
	CreatedAt *model.DateTime  `gorm:"column:created_at;type:datetime;comment:Creation Time" json:"createdAt"`                                   // Creation Time
	UpdatedAt *model.DateTime  `gorm:"column:updated_at;type:datetime;comment:Update Time" json:"updatedAt"`                                     // Update Time
	DeletedAt *model.DeletedAt `gorm:"column:deleted_at;type:datetime;comment:Delete Time" json:"deletedAt" swaggerignore:"true"`                // Delete Time
}

// TableName User's table name
func (*User) TableName() string {
	return TableNameUser
}
```

## Form Validation
### Validator Creation Help
```bash
$ go run cli.go make:request -h # --help

make:request - Validator Creation

Options:
  -f, --file  File Path, Example: user                         required:true
  -d, --desc  Description, Example: user-request-validation    required:false
```

### Validator Creation
```bash
$ go run cli.go make:request --file=user --desc=user-request-validation
```
```go
package request

import (
    "errors"
    "github.com/gookit/validate"
)

// User User-Request-Validation
type User struct {
    PageListValidate
}

// GetValidate Request-Validation
func (s User) GetValidate(data User, scene string) error {
	v := validate.Struct(data, scene)
	if !v.Validate(scene) {
		return errors.New(v.Errors.One())
	}

	return nil
}

// ConfigValidation Configuration-Validation
// - Define validation scenes
// - You can also add verification settings
func (s User) ConfigValidation(v *validate.Validation) {
	v.WithScenes(validate.SValues{
		"list":   []string{"PageListValidate.Page", "PageListValidate.PageSize"},
		"create": []string{},
		"update": []string{"ID"},
		"detail": []string{"ID"},
		"delete": []string{"ID"},
	})
}

// Messages Validator-Error-Message
func (s User) Messages() map[string]string {
	return validate.MS{
		"required":    "Field {field} Required",
		"int":         "Field {field} Must be an integer",
		"Page.gt":     "Field {field} Must be greater than 0",
		"PageSize.gt": "Field {field} Must be greater than 0",
	}
}

// Translates Field-Translation
func (s User) Translates() map[string]string {
	return validate.MS{
		"Page":     "Page",
		"PageSize": "Page Size",
		"ID":       "ID",
	}
}
```

### Validator Rules
> For more rules, please refer to [gookit/validate](https://github.com/gookit/validate)
```go
// UserCreate User-Create-Validation
type UserCreate struct {
	Username string `json:"username" validate:"required" label:"username"`
	FullName string `json:"fullName" validate:"required" label:"fullname"`
	Nickname string `json:"nickname" validate:"required" label:"nickname"`
	Gender   int    `json:"gender" validate:"required|int" label:"gender"`
	Password string `json:"password" validate:"required" label:"password"`
}

// UserUpdate User-Update-Validation
type UserUpdate struct {
	UserDetail
	UserCreate
}

// UserDetail User-Detail-Validation
type UserDetail struct {
    ID int64 `json:"id" validate:"required|int|gt:0" label:"ID"`
}

// User User-Request-Validation
type User struct {
	UserDetail
	UserCreate
	PageListValidate
}
```

### Validator Scenes
```go
// ConfigValidation Configuration-Validation
// - Define validation scenes
// - You can also add verification settings
func (s User) ConfigValidation(v *validate.Validation) {
	v.WithScenes(validate.SValues{
		// List
		"List": []string{
			"PageListValidate.Page",
			"PageListValidate.PageSize",
		},
		// Create
		"Create": []string{
			"UserCreate.Username",
			"UserCreate.FullName",
			"UserCreate.Nickname",
			"UserCreate.Gender",
			"UserCreate.Password",
		},
		// Update
		"Update": []string{
			"UserUpdate.UserDetail.ID",
			"UserCreate.Username",
			"UserCreate.FullName",
			"UserCreate.Nickname",
			"UserCreate.Gender",
		},
		// Detail
		"Detail": []string{
			"UserDetail.ID",
		},
		// Delete
		"Delete": []string{
			"UserDetail.ID",
		},
	})
}
```

### Prompt Message
```go
// Messages Validator-Error-Message
func (s User) Messages() map[string]string {
    return validate.MS{
        "required":                     "Field {field} Required",
        "int":                          "Field {field} Must be an integer",
        "PageListValidate.Page.gt":     "Field {field} Must be greater than 0",
        "PageListValidate.PageSize.gt": "Field {field} Must be greater than 0",
    }
}
```

### Field Translation
```go
// Translates Field-Translation
func (s User) Translates() map[string]string {
	return validate.MS{
		"Page":                "Page",
		"PageSize":            "Page Size",
		"ID":                  "ID",
		"UserCreate.Username": "Username",
		"UserCreate.FullName": "Fullname",
		"UserCreate.Nickname": "Nickname",
		"UserCreate.Gender":   "Gender",
		"UserCreate.Password": "Password",
	}
}
```

### Custom Validation
#### Global Rules
> Global rules only need to be defined in the entry file `main.go`,applicable to all validators, without the need for repeated definitions.
```go
package main

import (
  "github.com/gookit/validate"
)

// Register during initialization
func init() {
  validate.AddValidator("is_even", func(val any, rule string) bool {
    num, ok := val.(int)
    if !ok {
      return false
    }
    return num%2 == 0
  })
}
```

#### Local Rules
```go
// Define local rule methods (naming convention: Validate<rule name>)
func (s User) ValidateIsEven(val any) bool {
num := val.(int)
return num%2 == 0
}
```

#### Temporary Rules
```go
// GetValidate Request-Validation
func (s User) GetValidate(data User, scene string) error {
    v := validate.Struct(data, scene)
    v.AddValidator("is_even", func(val any, rule string) bool {
        num, ok := val.(int)
        if !ok {
            return false
        }
        return num%2 == 0
    })
	if !v.Validate(scene) {
		return errors.New(v.Errors.One())
	}

    return nil
}
```

#### Validator Usage
```go
type User struct {
    Age int `json:"gender" validate:"required|is_even" label:"age"`
}
```

#### Used In The Controller
```go
// List User-List
// @Tags User
// @Summary List
// @Description User-List
// @Param token header string true "Authentication Token"
// @Param page query string true "Page"
// @Param pageSize query string true "Page Size"
// @Success 200 {object} errcode.SuccessResponse{data=request.PageData{list=[]model.User}} "Login Successful"
// @Failure 400 {object} errcode.ArgsErrorResponse "Argument Error"
// @Failure 500 {object} errcode.SystemErrorResponse "System Error"
// @Router /api/v1/user [get]
func (s *UserController) List(c *gin.Context) {
	var (
		srv service.UserService
		req request.User
	)

	err := c.ShouldBind(&req)
	if err != nil {
		s.Error(c, errcode.SystemError().WithMsg(err.Error()))
		return
	}

	// Validator
	err = request.User{}.GetValidate(req, "List")
	if err != nil {
		s.Error(c, errcode.ArgsError().WithMsg(err.Error()))
		return
	}

	res, err := srv.List(req)
	if err != nil {
		s.Error(c, errcode.SystemError().WithMsg(err.Error()))
		return
	}

	s.Success(c, errcode.Success().WithData(res))
}
```

## Service
### Service Creation Help
```bash
$ go run cli.go make:service -h # --help

make:service - Service Creation

Options:
  -f, --file      File Path, Example: v1/user      required:true
  -F, --function  Function Name, Example: list     required:false
  -d, --desc      Description, Example: list       required:false
exit status 3
```

### Service Creation
```bash
$ go run cli.go make:service -f=user --function=list --desc="list"
```

## Command
### Get Version
```bash
$ go run cli.go --version # -v
Gin CLI v1.0.0
```

### Command Help
```bash
$ go run cli.go -h # --help

Usage: go run cli.go [command] [options]
Available commands:
  db:migrate               Database Migration(Automatic Table Creation/Schema Updates)
  db:seed                  Data Initialization
  demo-command             test-demo
  make:command             Service Creation
  make:controller          Controller Creation
  make:middleware          Middleware Creation
  make:model               Model Creation
  make:request             Validator Creation
  make:router              Route Creation
  make:service             Service Creation
  route:list               Route List

Options:
  -f, --format   The output format (txt, json) [default: txt]
  -h, --help     Display help for the given command. When no command is given display help for the list command
  -v, --version  Display this application version
```

### Command List
```bash
$ go run cli.go --format=json # -f=json

{
  "commands": [
    {
      "description": "Database Migration(Automatic Table Creation/Schema Updates)",
      "name": "db:migrate"
    },
    {
      "description": "Data Initialization",
      "name": "db:seed"
    },
    {
      "description": "test-demo",
      "name": "demo-command"
    },
    {
      "description": "Service Creation",
      "name": "make:command"
    },
    {
      "description": "Controller Creation",
      "name": "make:controller"
    },
    {
      "description": "Middleware Creation",
      "name": "make:middleware"
    },
    {
      "description": "Model Creation",
      "name": "make:model"
    },
    {
      "description": "Validator Creation",
      "name": "make:request"
    },
    {
      "description": "Route Creation",
      "name": "make:router"
    },
    {
      "description": "Service Creation",
      "name": "make:service"
    },
    {
      "description": "Route List",
      "name": "route:list"
    }
  ],
  "version": "Gin CLI v1.0.0"
}
```

## Command Creation Help
```bash
$ go run cli.go make:command -h # --help

make:command - Command Creation

Options:
  -f, --file  File Path, Example: cronjob/demo     required:true
  -n, --name  Command Name, Example: demo-test     required:false
  -d, --desc  Description, Example: command-desc   required:false
```

## Command Creation
```bash
$ go run cli.go make:command --file=cronjob/demo --name=demo-test --desc=command-desc
```

## Command Structure
> After generating the command, appropriate values should be defined for the ` Name() ` and ` Descript() ` functions. These properties will be used when displaying the command list. The `Name()` function also allows you to define the expected input value for the command. It will call the `Execute()` function when executing the command. You can put the command logic in this method. Let's take a look at an example command.
```go
package cronjob

import (
	"gin/common/base"
	"gin/utils/cli"
	"github.com/fatih/color"
)

type DemoCommand struct {
	base.BaseCommand
}

func (m *DemoCommand) Name() string {
    return "demo-test"
}

func (m *DemoCommand) Description() string {
	return "command-desc"
}

func (m *DemoCommand) Help() []base.CommandOption {
	return []base.CommandOption{
        {
            base.Flag{
                Short: "a",
                Long:  "args",
            },
            "Example Argument, Example: arg1",
            true,
        },
    }
}

func (m *DemoCommand) Execute(args []string) {
    values := m.ParseFlags(m.Name(), args, m.Help())
    color.Green("Execute Command: %s %s", m.Name(), m.FormatArgs(values))
}

func init() {
	cli.Register(&DemoCommand{})
}

```

## Command Registration
> `cli. go` registers all commands in the `command` package under the `gin/app/command` directory by default. If the command you registered is not `command` package, you can add the path to import the package in `cli. go`.
```go
package main

import (
	_ "gin/app/command"
	_ "gin/app/command/cronjob"
	"gin/utils/cli"
	_ "gin/utils/cli/db"
	_ "gin/utils/cli/make"
	_ "gin/utils/cli/route"
)

func main() {
	cli.Execute()
}

```

## Help Options
> Command option parameters are defined using the `base. CommandOption` structure. The `base. CommandOption` struct contains two attributes: `Flag` and `Description`. The `Flag` attribute is used to define the flag of command options, which can be a short flag (such as `- a `) or a long flag (such as `--args`). The `Description` attribute is used to define the description of command options. The `base. CommandOption` struct also contains a `Required` attribute that specifies whether a command option is required. At the same time, this method supports the console `--help` parameter and automatically generates help information.
```go
func (m *DemoCommand) Help() []base.CommandOption {
	return []base.CommandOption{
        {
            base.Flag{
                Short: "a",
                Long:  "args",
            },
            "Example Argument, Example: arg1",
            true,
        },
    }
}
```
```bash
$ go run cli.go demo-test -h # --help

demo-test - command-desc

Options:
  -a, --args  Example Argument, Example: arg1  required:true
```

## Execute Command
```bash
$ go run cli.go demo-test --args=arg1

Execute Command: demo-test --args=arg1
```

## Compile And Execute Commands
```bash
$ go build cli.go
$ ./cli demo-test --args=arg1
```

# Language Support
> Implement multilingual support using the `i18n` package, supporting both `zh` and `en` languages, and supporting custom extensions. Language transmission defaults to transmitting the `Accept-Language` parameter in the `header`, such as `zh` or `en`, which is not case sensitive and does not pass the default language as `zh`.
## Directory Configuration
> The storage path for translation files is `storage/scales`, the default language is `zh`, and multiple languages are separated by commas. Languages are stored in the corresponding language directory without distinguishing between subdirectories. For example, Chinese is stored in `storage/scales/zh` and can support `json` and `yaml` format files in any directory.
```yaml
# Translation Configuration
i18n:
  dir: "storage/locales" # Translation file storage path
  lang: "zh,en" # Default language, multiple languages separated by commas
```

## Ordinary Translation
```go
import (
    "gin/utils/lang"
)

func Test()  {
    trans := lang.T("login.username", nil)
	fmt.Println(trans) // Output: 用户名, English Output: Username
}
```

## Template Translation
> Template translation is supported in the translation file, such as `{{. name}}`, using `map[string]interface{}` to pass parameters.
```json
[
  {
    "id": "login.success",
    "translation": "{{.name}},Login Success"
  }
]
```
```go
import (
    "gin/utils/lang"
)

func Test()  {
    trans := lang.T("login.success", map[string]interface{}{
        "name": "admin",
    }),
	fmt.Println(trans) // Output: admin,登录成功 English Output: admin,Login Success
}
```

## Add Language Support
> Add the corresponding language directory, such as `en`, in the `storage/scales` directory, and then add a translation file in the directory. The translation file supports `json` and `yaml` formats, with `id` as the unique identifier and `translation` as the translation content. Any number of translation contents can be added to the translation file. The configuration language support requires adjusting the `i18n.lang` parameter in the configuration file.
```yaml
# Translation Configuration
i18n:
  dir: "storage/locales" # Translation file storage path
  lang: "zh,en" # Default language, multiple languages separated by commas
```

# Swagger Documents
```bash
$ go install github.com/swaggo/swag/cmd/swag@latest
$ swag init -g main.go # --exclude cli,app/service
2025/10/23 16:26:42 Generate swagger docs....
2025/10/23 16:26:42 Generate general API Info, search dir:./
2025/10/23 16:26:43 Generating request.UserLogin
2025/10/23 16:26:43 Generating errcode.SuccessResponse
2025/10/23 16:26:43 Generating v1.LoginResponse
2025/10/23 16:26:43 Generating v1.Token
2025/10/23 16:26:43 Generating model.User
2025/10/23 16:26:43 Generating model.DateTime
2025/10/23 16:26:43 Generating errcode.ArgsErrorResponse
2025/10/23 16:26:43 Generating errcode.SystemErrorResponse
2025/10/23 16:26:43 Generating request.PageData
2025/10/23 16:26:43 Generating request.UserCreate
2025/10/23 16:26:43 Generating request.UserUpdate
2025/10/23 16:26:43 Generating request.UserDetail
2025/10/23 16:26:43 create docs.go at docs/docs.go
2025/10/23 16:26:43 create swagger.json at docs/swagger.json
2025/10/23 16:26:43 create swagger.yaml at docs/swagger.yaml
```
