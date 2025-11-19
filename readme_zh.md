## 中文 | [English](readme.md)

- [项目简介](#项目简介)
- [许可证](#许可证)
- [版本记录](#版本记录)
- [安装说明](#安装说明)
  -[克隆项目](#克隆项目)
  -[初始化Go环境与依赖](#初始化Go环境与依赖)
    -[方式一](#方式一)
    -[方式二](#方式二)
  -[启动](#启动)
    -[使用air热更新](#使用air热更新)
  -[编译](#编译)
    -[编译项目](#编译项目)
    -[编译命令行](#编译命令行)
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
  - [模型](#模型)
    - [模型创建帮助](#模型创建帮助)
    - [模型创建](#模型创建)
  - [表单验证](#表单验证)
    - [验证创建帮助](#验证创建帮助)
    - [验证创建](#验证创建)
    - [验证规则](#验证规则)
    - [验证场景](#验证场景)
    - [提示信息](#提示信息)
    - [字段翻译](#字段翻译)
    - [自定义验证](#自定义验证)
      - [全局规则](#全局规则)
      - [局部规则](#局部规则)
      - [临时规则](#临时规则)
      - [验证使用](#验证使用)
      - [在控制器中使用](#在控制器中使用)
  - [服务](#服务)
    - [服务创建帮助](#服务创建帮助)
    - [服务创建](#服务创建)
  - [命令行](#命令行)
    - [获取版本](#获取版本)
    - [命令帮助](#命令帮助)
    - [命令列表](#命令列表)
    - [命令创建帮助](#命令创建帮助)
    - [命令创建](#命令创建)
    - [命令结构](#命令结构)
    - [命令注册](#命令注册)
    - [帮助选项](#帮助选项)
    - [执行命令](#执行命令)
    - [编译执行](#编译执行)
  - [缓存](#缓存)
    - [全局缓存](#全局缓存)
    - [Redis缓存](#Redis缓存)
    - [内存缓存](#内存缓存)
    - [磁盘缓存](#磁盘缓存)
  - [事件](#事件)
    - [事件创建帮助](#事件创建帮助)
    - [事件创建](#事件创建)
  - [监听](#监听)
    - [监听创建帮助](#监听创建帮助)
    - [监听创建](#监听创建)
  - [发布事件](#发布事件)
    - [测试事件](#测试事件)
  - [事件列表](#事件列表)
    - [事件监听列表](#事件监听列表)
  - [响应](#响应)
    - [成功响应](#成功响应)
      - [成功提示](#成功提示)
      - [成功数据](#成功数据)
    - [失败响应](#失败响应)
      - [失败错误码](#失败错误码)
      - [失败提示](#失败提示)
      - [失败数据](#失败数据)
  - [日志](#日志)
    - [记录日志](#记录日志)
    - [错误调试](#错误调试)
  - [多语言](#多语言)
    - [目录配置](#目录配置)
    - [常规翻译](#常规翻译)
    - [模版翻译](#模版翻译)
    - [添加语言](#添加语言)
  - [swagger文档](#swagger文档)

# 项目简介
> 基于`Golang`语言框架`Go Gin`开发的轻量级框架，开箱即用，设计灵感基于`Laravel`、`ThinPHP`等主流`PHP`框架，项目架构目录层次分明，初学者的福音，框架使用`JWT`、`中间件`、`缓存`、`验证器`、`事件`、`路由`、`redis`、`命令行`等,支持多语言,开发简单易于上手, 方便扩展。
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

# 版本记录
## v1.2.1
> - 优化上下文处理
> - 优化日志处理以及redis、http、mysql等日志处理
> - 优化后readme文档完善

## v1.2.0
> - 优化上下文处理
> - 优化日志处理
> - 新增消息发布订阅
> - 优化后readme文档完善

## v1.1.0
> 完善日志调试以及使用文档, 完成版本v1.0.0。

## v1.0.3
> 完善公共响应使用文档。

## v1.0.2
> 错误码优化。

## v1.0.1
> 新增以公共包函数`FilterFields`, 调整公共包函数`StructToMap`, 调整json序列化使用包`go-json`。

## v1.0.0
> 除响应、错误处理、日志文档未完善其他已完成更新。

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
$ go install github.com/air-verse/air@latest
$ air
```

## 编译
### 编译项目
```bash
$ go build main.go
$ ./main
```

### 编译命令行
```bash
$ go build cli.go
$ ./cli demo-command --args=11

执行命令: demo-command, 参数: 11
```

# 目录结构
```
├── app                                 # 应用程序
│   ├── command                         # 命令
│   ├── controller                      # 控制器
│   ├── event                           # 事件
│   ├── listener                        # 监听
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
│   ├── locales                         # 翻译文件
│   ├──├── en                           # 英文翻译
│   ├──├── zh                           # 中文翻译
├── tests                               # 测试用例
├── utils                               # 工具包
│   ├──├── cache                        # 缓存
│   ├──├── cli                          # 命令行
│   ├──├── ctx                          # 上下文
│   ├──├── eventbus                     # 事件
│   ├──├── lang                         # 多语言
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

## 模型
### 模型创建帮助
```bash
$ go run cli.go make:model -h # --help

make:model - 模型创建

Options:
  -t, --table  表名, 如: user 或 user,menu  required:true
  -p, --path   输出目录, 如: api/user       required:false
  -c, --camel  是否驼峰字段, 如: true       required:false
```

### 模型创建
> 支持创建同时多个模型文件, 如需创建多个模型文件, 传人的表名参数请使用逗号分隔, 如: user,menu
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

// User 用户表
type User struct {
	ID        int64            `gorm:"column:id;type:int(10) unsigned;primaryKey;autoIncrement:true;comment:ID" json:"id"`           // ID
	Avatar    string           `gorm:"column:avatar;type:varchar(255);not null;comment:头像" json:"avatar"`                            // 头像
	Username  string           `gorm:"column:username;type:varchar(10);not null;comment:用户名" json:"username"`                        // 用户名
	FullName  string           `gorm:"column:full_name;type:varchar(20);not null;comment:姓名" json:"fullName"`                        // 姓名
	Email     string           `gorm:"column:email;type:varchar(50);not null;comment:邮箱" json:"email"`                               // 邮箱
	Password  string           `gorm:"column:password;type:varchar(255);not null;comment:密码" json:"password"`                        // 密码
	Nickname  string           `gorm:"column:nickname;type:varchar(50);not null;comment:昵称" json:"nickname"`                         // 昵称
	Gender    int64            `gorm:"column:gender;type:tinyint(1) unsigned;not null;comment:性别 1=男 2=女" json:"gender"`             // 性别 1=男 2=女
	Age       int64            `gorm:"column:age;type:int(10) unsigned;not null;comment:年龄" json:"age"`                              // 年龄
	Status    int64            `gorm:"column:status;type:tinyint(3) unsigned;not null;default:1;comment:状态 1=启用 2=停用" json:"status"` // 状态 1=启用 2=停用
	CreatedAt *model.DateTime  `gorm:"column:created_at;type:datetime;comment:创建时间" json:"createdAt"`                                // 创建时间
	UpdatedAt *model.DateTime  `gorm:"column:updated_at;type:datetime;comment:更新时间" json:"updatedAt"`                                // 更新时间
	DeletedAt *model.DeletedAt `gorm:"column:deleted_at;type:datetime;comment:删除时间" json:"deletedAt" swaggerignore:"true"`                                // 删除时间
}

// TableName User's table name
func (*User) TableName() string {
	return TableNameUser
}
```

## 表单验证
### 验证创建帮助
```bash
$ go run cli.go make:request -h # --help

make:request - 验证请求创建

Options:
  -f, --file  文件路径, 如: user     required:true
  -d, --desc  描述, 如: 用户请求验证  required:false
```

### 验证创建
```bash
$ go run cli.go make:request --file=user --desc=用户请求验证
```
```go
package request

import (
    "errors"
    "github.com/gookit/validate"
)

// User 用户请求验证
type User struct {
    PageListValidate
}

// GetValidate 请求验证
func (s User) GetValidate(data User, scene string) error {
	v := validate.Struct(data, scene)
	if !v.Validate(scene) {
		return errors.New(v.Errors.One())
	}

	return nil
}

// ConfigValidation 配置验证
// - 定义验证场景
// - 也可以添加验证设置
func (s User) ConfigValidation(v *validate.Validation) {
	v.WithScenes(validate.SValues{
		"list":   []string{"PageListValidate.Page", "PageListValidate.PageSize"},
		"create": []string{},
		"update": []string{"ID"},
		"detail": []string{"ID"},
		"delete": []string{"ID"},
	})
}

// Messages 验证器错误消息
func (s User) Messages() map[string]string {
	return validate.MS{
		"required":    "字段 {field} 必填",
		"int":         "字段 {field} 必须为整数",
		"Page.gt":     "字段 {field} 需大于 0",
		"PageSize.gt": "字段 {field} 需大于 0",
	}
}

// Translates 字段翻译
func (s User) Translates() map[string]string {
	return validate.MS{
		"Page":     "页码",
		"PageSize": "每页数量",
		"ID":       "ID",
	}
}
```

### 验证规则
> 更多规则请查看 [gookit/validate](https://github.com/gookit/validate)
```go
// UserCreate 用户创建验证
type UserCreate struct {
	Username string `json:"username" validate:"required" label:"用户名"`
	FullName string `json:"fullName" validate:"required" label:"姓名"`
	Nickname string `json:"nickname" validate:"required" label:"昵称"`
	Gender   int    `json:"gender" validate:"required|int" label:"性别"`
	Password string `json:"password" validate:"required" label:"密码"`
}

// UserUpdate 用户更新验证
type UserUpdate struct {
	UserDetail
	UserCreate
}

// UserDetail 用户详情验证
type UserDetail struct {
    ID int64 `json:"id" validate:"required|int|gt:0" label:"ID"`
}

// User 用户请求验证
type User struct {
	UserDetail
	UserCreate
	PageListValidate
}
```

### 验证场景
```go
// ConfigValidation 配置验证
// - 定义验证场景
// - 也可以添加验证设置
func (s User) ConfigValidation(v *validate.Validation) {
	v.WithScenes(validate.SValues{
		// 列表
		"List": []string{
			"PageListValidate.Page",
			"PageListValidate.PageSize",
		},
		// 创建
		"Create": []string{
			"UserCreate.Username",
			"UserCreate.FullName",
			"UserCreate.Nickname",
			"UserCreate.Gender",
			"UserCreate.Password",
		},
		// 更新
		"Update": []string{
			"UserUpdate.UserDetail.ID",
			"UserCreate.Username",
			"UserCreate.FullName",
			"UserCreate.Nickname",
			"UserCreate.Gender",
		},
		// 详情
		"Detail": []string{
			"UserDetail.ID",
		},
		// 删除
		"Delete": []string{
			"UserDetail.ID",
		},
	})
}
```

### 提示信息
```go
// Messages 验证器错误消息
func (s User) Messages() map[string]string {
	return validate.MS{
        "required":                     "字段 {field} 必填",
        "int":                          "字段 {field} 必须为整数",
        "PageListValidate.Page.gt":     "字段 {field} 需大于 0",
        "PageListValidate.PageSize.gt": "字段 {field} 需大于 0",
	}
}
```

### 字段翻译
```go
// Translates 字段翻译
func (s User) Translates() map[string]string {
	return validate.MS{
		"Page":                "页码",
		"PageSize":            "每页数量",
		"ID":                  "ID",
		"UserCreate.Username": "用户名",
		"UserCreate.FullName": "姓名",
		"UserCreate.Nickname": "昵称",
		"UserCreate.Gender":   "性别",
		"UserCreate.Password": "密码",
	}
}
```

### 自定义验证
#### 全局规则
> 全局规则只需要在入口文件`main.go`中定义, 适用于所有验证器, 无需重复定义。
```go
package main

import (
	"github.com/gookit/validate"
)

// 初始化时注册
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

#### 局部规则
```go
// 定义局部规则方法(命名规则：Validate<规则名>)
func (s User) ValidateIsEven(val any) bool {
	num := val.(int)
	return num%2 == 0
}
```

#### 临时规则
```go
// GetValidate 请求验证
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

#### 验证使用
```go
type User struct {
    Age int `json:"gender" validate:"required|is_even" label:"年龄"`
}
```

#### 在控制器中使用
```go
// List 列表
// @Tags 用户管理
// @Summary 列表
// @Description 用户列表
// @Param token header string true "认证Token"
// @Param page query string true "页码"
// @Param pageSize query string true "分页大小"
// @Success 200 {object} errcode.SuccessResponse{data=request.PageData{list=[]model.User}} "登录成功"
// @Failure 400 {object} errcode.ArgsErrorResponse "参数错误"
// @Failure 500 {object} errcode.SystemErrorResponse "系统错误"
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

	// 验证
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

## 服务
### 服务创建帮助
```bash
$ go run cli.go make:service -h # --help

make:service - 服务创建

Options:
  -f, --file      文件路径, 如: v1/user  required:true
  -F, --function  方法名称, 如: list     required:false
  -d, --desc      描述, 如: 列表         required:false
exit status 3
```

### 服务创建
```bash
$ go run cli.go make:service -f=user --function=list --desc="列表"
```

## 命令行
### 获取版本
```bash
$ go run cli.go --version # -v
Gin CLI v1.0.0
```

### 命令帮助
```bash
$ go run cli.go -h # --help

Usage: go run cli.go [command] [options]
Available commands:
  db:migrate               数据库迁移(自动建表/更新结构)
  db:seed                  数据初始化
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
  -f, --format   The output format (txt, json) [default: txt]
  -h, --help     Display help for the given command. When no command is given display help for the list command
  -v, --version  Display this application version
```

### 命令列表
```bash
$ go run cli.go --format=json # -f=json

{
  "commands": [
    {
      "description": "数据库迁移(自动建表/更新结构)",
      "name": "db:migrate"
    },
    {
      "description": "数据初始化",
      "name": "db:seed"
    },
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

## 命令创建帮助
```bash
$ go run cli.go make:command -h # --help

make:command - 命令创建

Options:
  -f, --file  文件路径, 如: cronjob/demo  required:true
  -n, --name  命令名称, 如: demo-test     required:false
  -d, --desc  描述, 如: command-desc      required:false
```

## 命令创建
```bash
$ go run cli.go make:command --file=cronjob/demo --name=demo-test --desc=command-desc
```

## 命令结构
> 生成命令后，应为`Name()` 和 `Description()` 方法定义适当的值。当在显示命令列表时，将使用这些属性。 `Name()` 方法还允许你定义命令的输入期望值。 `Execute()` 执行命令时将调用该方法。你可以将命令逻辑放在此方法中。 让我们看一个示例命令。
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
            "示例参数, 如: arg1",
            true,
        },
    }
}

func (m *DemoCommand) Execute(args []string) {
    values := m.ParseFlags(m.Name(), args, m.Help())
    color.Green("执行命令: %s %s", m.Name(), m.FormatArgs(values))
}

func init() {
	cli.Register(&DemoCommand{})
}

```

## 命令注册
> `cli.go` 默认注册了 `gin/app/command` 目录下的 `command` 包的所有命令，如果你注册的命令不是一个包，可以在 `cli.go` 中添加导入包的路径。
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

## 帮助选项
> 命令选项参数使用 `base.CommandOption` 结构体来定义。 `base.CommandOption` 结构体包含两个属性： `Flag` 和 `Description`。 `Flag` 属性用于定义命令选项的标志，可以是短标志（如 `-a`）或长标志（如 `--args`）。 `Description` 属性用于定义命令选项的描述。 `base.CommandOption` 结构体还包含一个 `Required` 属性，用于指定命令选项是否为必需的。同时该方法支持控制台 `--help` 参数，自动生成帮助信息。
```go
func (m *DemoCommand) Help() []base.CommandOption {
	return []base.CommandOption{
        {
            base.Flag{
                Short: "a",
                Long:  "args",
            },
            "示例参数, 如: arg1",
            true,
        },
    }
}
```
```bash
$ go run cli.go demo-test -h # --help

demo-test - command-desc

Options:
  -a, --args  示例参数, 如: arg1  required:true
```

## 执行命令
```bash
$ go run cli.go demo-test --args=arg1

执行命令: demo-test --args=arg1
```

## 编译执行
```bash
$ go build cli.go
$ ./cli demo-test --args=arg1
```

# 缓存
> 使用了全局缓存, 默认使用 `memory` 作为缓存驱动, 支持自定义扩展。默认支持`内存缓存`、`Redis缓存`、`磁盘缓存`三种模式, 可使用全局缓存也可单独使用任意缓存。全局缓存默认只集成了`Set`、`Get`、`Delete`、`Expire`公共方法如需使用更多可以单独使用,你也可以自己集成。
## 全局缓存
> 配置全局缓存可通过`yaml`配置文件中的`cache.driver`配置进行切换。
```go
import (
	"fmt"
    "gin/common/global"
)

func Test()  {
    // Set 设置缓存	
    key := "test_key"
    value := "test_value"
    err := global.Cache.Set(key, value, time.Second*10)
	if err != nil {
	    // 处理错误	
    }
	
    // Get 获取缓存
    key := "test_key"
    value := "test_value"
    result, ok := global.Cache.Get(key)
	if ok {
	    println(result) // test_value	
    }
	
	// Delete 删除缓存
	key := "test_key"
	err := global.Cache.Delete(key)
	if err != nil {
        // 处理错误	
    }
	
	// Expire 获取缓存过期时间
	key := "test_key"
    val, expireAt, ok, err := global.Cache.Expire(key)
	if err != nil {
	    // 处理错误
    }
	if ok {
      fmt.Println(val) // test_value
      fmt.Printf("ExpireAt: %v\n", expireAt) // ExpireAt: 2025-10-28 11:23:38.7416956 +0800 CST
    }
}
```

## Redis缓存
```go
import (
	"fmt"
    "gin/common/global"
)

func Test()  {
    // Set 设置缓存	
    key := "test_key"
    value := "test_value"
    err := global.RedisCache.Set(key, value, time.Second*10)
	if err != nil {
	    // 处理错误	
    }
	
    // Get 获取缓存
    key := "test_key"
    value := "test_value"
    result, ok := global.RedisCache.Get(key)
	if ok {
	    println(result) // test_value	
    }
	
	// Delete 删除缓存
	key := "test_key"
	err := global.RedisCache.Delete(key)
	if err != nil {
        // 处理错误	
    }
	
	// Expire 获取缓存过期时间
	key := "test_key"
    val, expireAt, ok, err := global.RedisCache.Expire(key)
	if err != nil {
	    // 处理错误
    }
	if ok {
      fmt.Println(val) // test_value
      fmt.Printf("ExpireAt: %v\n", expireAt) // ExpireAt: 2025-10-28 11:23:38.7416956 +0800 CST
    }
	
	// ... 其他
}
```

## 内存缓存
```go
import (
	"fmt"
    "gin/common/global"
)

func Test()  {
    // Set 设置缓存	
    key := "test_key"
    value := "test_value"
    err := global.MemoryCache.Set(key, value, time.Second*10)
	if err != nil {
	    // 处理错误	
    }
	
    // Get 获取缓存
    key := "test_key"
    value := "test_value"
    result, ok := global.MemoryCache.Get(key)
	if ok {
	    println(result) // test_value	
    }
	
	// Delete 删除缓存
	key := "test_key"
	err := global.MemoryCache.Delete(key)
	if err != nil {
        // 处理错误	
    }
	
	// Expire 获取缓存过期时间
	key := "test_key"
    val, expireAt, ok, err := global.MemoryCache.Expire(key)
	if err != nil {
	    // 处理错误
    }
	if ok {
      fmt.Println(val) // test_value
      fmt.Printf("ExpireAt: %v\n", expireAt) // ExpireAt: 2025-10-28 11:23:38.7416956 +0800 CST
    }
	
	// ... 其他
}
```

## 磁盘缓存
```go
// Set 设置缓存	
    key := "test_key"
    value := "test_value"
    err := global.DiskCache.Set(key, value, time.Second*10)
	if err != nil {
	    // 处理错误	
    }
	
    // Get 获取缓存
    key := "test_key"
    value := "test_value"
    result, ok := global.DiskCache.Get(key)
	if ok {
	    println(result) // test_value	
    }
	
	// Delete 删除缓存
	key := "test_key"
	err := global.DiskCache.Delete(key)
	if err != nil {
        // 处理错误	
    }
	
	// Expire 获取缓存过期时间
	key := "test_key"
    val, expireAt, ok, err := global.DiskCache.Expire(key)
	if err != nil {
	    // 处理错误
    }
	if ok {
      fmt.Println(val) // test_value
      fmt.Printf("ExpireAt: %v\n", expireAt) // ExpireAt: 2025-10-28 11:23:38.7416956 +0800 CST
    }
	
	// ... 其他
```

# 事件
## 事件创建帮助
```bash
$ go run cli.go make:event -h # --help

make:event - 创建事件

Options:
  -f, --file  文件路径, 如: login/test  required:true
  -n, --name  事件名称, 如: test-event  required:false
  -d, --desc  事件描述, 如: 测试事件     required:false
```

## 事件创建
```bash
$ go run cli.go make:event -f=user_login -n='user.login' -d=用户登录事件
```
```go
package event

// UserLoginEvent 事件数据
type UserLoginEvent struct {
	UserId   int64
	Username string
}

// Name 事件名称
func (u UserLoginEvent) Name() string {
	return "user.login"
}

// Description 事件描述
func (u UserLoginEvent) Description() string {
	return "用户登录事件"
}

```

# 监听
## 监听创建帮助
```bash
$ go run cli.go make:listener -h # --help

make:listener - 创建监听

Options:
  -f, --file   文件路径, 如: login/test  required:true
  -e, --event  事件数据, 如: UserLogin   required:true
```

## 监听创建
```bash
$ go run cli.go make:listener -f=user_login -e=UserLoginEvent
```
```go
package listener

import (
	"github.com/goccy/go-json"
	"fmt"
	"gin/app/event"
	"gin/utils/eventbus"
	"time"
)

type UserLoginListener struct{}

func (l *UserLoginListener) Handle(e event.UserLoginEvent) {
	data, _ := json.Marshal(e)
	fmt.Printf("收到事件: %s 事件描述: %s 事件数据: %s, 时间: %s\n", e.Name(), e.Description(), data, time.Now().Format("2006-01-02 15:04:05"))
}

func init() {
	eventbus.Register(&UserLoginListener{}, event.UserLoginEvent{})
}

```

# 发布事件
```go
package v1

import (
	"gin/app/event"
	"gin/app/middleware"
	"gin/app/model"
	"gin/app/request"
	"gin/app/service"
	"gin/common/base"
	"gin/common/errcode"
	"gin/common/global"
	"gin/utils/eventbus"
	"gin/utils/lang"
	"github.com/gin-gonic/gin"
)

type LoginController struct {
	base.BaseController
}

// Token token信息
type Token struct {
	AccessToken        string `json:"accessToken"`
	RefreshToken       string `json:"refreshToken"`
	TokenExpire        int64  `json:"tokenExpire" example:"7200"`
	RefreshTokenExpire int64  `json:"refreshTokenExpire" example:"172800"`
}

type LoginResponse struct {
	Token Token `json:"token"`
	User  model.User
}

// Login 登录
// @Tags 登录相关
// @Summary 登录
// @Description 用户登录
// @Accept json
// @Produce json
// @Param data body request.UserLogin true "登录参数"
// @Success 200 {object} errcode.SuccessResponse{data=LoginResponse} "成功"
// @Failure 400 {object} errcode.ArgsErrorResponse "参数错误"
// @Failure 500 {object} errcode.SystemErrorResponse "系统错误"
// @Router /api/v1/login [post]
func (s *LoginController) Login(c *gin.Context) {
	var (
		srv service.LoginService
		req request.Login
		jwt middleware.Jwt
	)

	err := c.ShouldBind(&req)
	if err != nil {
		s.Error(c, errcode.SystemError().WithMsg(err.Error()))
		return
	}

	// 验证
	err = request.Login{}.GetValidate(req, "Login")
	if err != nil {
		s.Error(c, errcode.ArgsError().WithMsg(err.Error()))
		return
	}

	userModel, err := srv.Login(req.Username, req.Password)
	if err != nil {
		s.Error(c, errcode.SystemError().WithMsg(lang.T(err.Error(), nil)))
		return
	}

	accessToken, refreshToken, tokenExpire, refreshTokenExpire, err := jwt.WithRefresh(userModel.ID, global.Config.Jwt.Exp, global.Config.Jwt.RefreshExp)
	if err != nil {
		s.Error(c, errcode.ArgsError().WithMsg(err.Error()))
		return
	}

	// 发布事件
	eventbus.Publish(event.UserLoginEvent{
		UserId:   userModel.ID,
		Username: userModel.Username,
	})

	s.Success(
		c, errcode.Success().WithMsg(
			lang.T("login.success", map[string]interface{}{
				"name": userModel.Username,
			}),
		).WithData(LoginResponse{
			Token{
				AccessToken:        accessToken,
				RefreshToken:       refreshToken,
				TokenExpire:        tokenExpire,
				RefreshTokenExpire: refreshTokenExpire,
			},
			userModel,
		}),
	)
}
```
## 测试事件
```bash
$ POST /api/v1/login HTTP/1.1
Host: 127.0.0.1:8080
Accept-Language: en-Us
Content-Type: application/json
Content-Length: 56

{
    "username": "admin",
    "password": "123456"
}

收到事件: user.login 事件描述: 用户登录事件 事件数据: {"UserId":1,"Username":"admin"}, 时间: 2025-11-04 15:32:12
```

# 事件列表
```bash
$ go run cli.go event:list

user.login 用户登录事件
```

## 事件监听列表
```bash
$ go run cli.go event-listener:list

==== 当前已注册事件 ====
事件: user.login
描述: 用户登录事件
监听:
  - *listener.TestListener
  - *listener.UserLoginListener
----------------------
```

# 响应
## 成功响应
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

func (s *TestController) Test(c *gin.Context) {
    return s.Success(c, errcode.Success())
}
```

### 成功提示
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

func (s *TestController) Test(c *gin.Context) {
    return s.Success(c, errcode.Success().WithMsg("Success"))
}
```

### 成功数据
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

func (s *TestController) Test(c *gin.Context) {
    return s.Success(c, errcode.Success().WithData([]string{"test data"}))
}
```

## 失败响应
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

func (s *TestController) Test(c *gin.Context) {
    return s.Error(c, errcode.SystemError())
}
```

### 失败错误码
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

func (s *TestController) Test(c *gin.Context) {
    return s.Error(c, errcode.SystemError().WithCode(500))
}
```

### 失败提示
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

func (s *TestController) Test(c *gin.Context) {
    return s.Error(c, errcode.SystemError().WithMsg("System Error"))
}
```

### 失败数据
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

func (s *TestController) Test(c *gin.Context) {
    return s.Error(c, errcode.SystemError().WithData([]string{"test data"}))
}
```

# 日志
> 使用 `zap` 包实现日志记录，日志文件存放路径为 `storage/logs`, 默认日志级别为 `debug`, 返回错误码不为0时自动记录日志TraceId、堆栈、sql、http、redis等调用信息, 也可以直接调用日志记录也会自动记录调试信息。配置文件`yaml`中`log.access`支持是否自动记录请求日志，如若开启会自动记录请求日志。
```json
{
  "level": "info",
  "timestamp": "2025-11-17 16:35:09.402",
  "caller": "middleware/logger.go:83",
  "msg": "Access Log",
  "traceId": "fa505122-d31e-4d4f-a05c-13c1641d6c6c",
  "ip": "127.0.0.1",
  "path": "/api/v1/login",
  "method": "POST",
  "params": {
    "password": "1234561",
    "username": "admin"
  },
  "ms": 59,
  "debugger": {
    "Sql": [
      {
        "ms": 2.5008,
        "rows": 1,
        "sql": "SELECT * FROM `user` WHERE username = 'admin' AND `user`.`deleted_at` IS NULL ORDER BY `user`.`id` LIMIT 1"
      }
    ],
    "Redis": null,
    "Http": null,
    "Rabbitmq": null
  }
}
```
## 记录日志
> 已封装全局日志在`global`包中，可直接使用`global.Log`记录日志, 日志级别支持debug、info、warn、error、panic、fatal, 默认为`debug`。
```go
package v1

import (
    "gin/common/base"
    "gin/common/global"
    "github.com/gin-gonic/gin"
)

type TestController struct {
    base.BaseController
}

func (s *TestController) Test(c *gin.Context) {
  global.Log.Error(c, "System Error")
}
```

## 错误调试
> 返回错误码不为0时自动记录日志TraceId、堆栈、sql、http、redis等调用信息, 直接调用日志记录也会自动记录调试信息, 可根据debug调试信息和trace堆栈信息调试, 日志文件存放路径为 `storage/logs`。
```json
{
  "level": "error",
  "timestamp": "2025-11-17 16:35:09.401",
  "caller": "response/response.go:60",
  "msg": "Login Password Error",
  "traceId": "fa505122-d31e-4d4f-a05c-13c1641d6c6c",
  "ip": "127.0.0.1",
  "path": "/api/v1/login",
  "method": "POST",
  "params": {
    "password": "1234561",
    "username": "admin"
  },
  "ms": 58,
  "debugger": {
    "Sql": [
      {
        "ms": 2.5008,
        "rows": 1,
        "sql": "SELECT * FROM `user` WHERE username = 'admin' AND `user`.`deleted_at` IS NULL ORDER BY `user`.`id` LIMIT 1"
      }
    ],
    "Redis": null,
    "Http": null,
    "Rabbitmq": null
  },
  "stackTrace": "gin/common/response.Error\n\tE:/www/dsx/www-go/gin/common/response/response.go:60\ngin/common/base.(*BaseController).Error\n\tE:/www/dsx/www-go/gin/common/base/base_controller.go:25\ngin/app/controller/v1.(*LoginController).Login\n\tE:/www/dsx/www-go/gin/app/controller/v1/login.go:67\ngithub.com/gin-gonic/gin.(*Context).Next\n\tE:/www/dsx/www-go/gin/vendor/github.com/gin-gonic/gin/context.go:192\ngin/router.init.Cors.Handle.func2\n\tE:/www/dsx/www-go/gin/app/middleware/cors.go:30\ngithub.com/gin-gonic/gin.(*Context).Next\n\tE:/www/dsx/www-go/gin/vendor/github.com/gin-gonic/gin/context.go:192\ngin/router.init.Logger.Handle.func1\n\tE:/www/dsx/www-go/gin/app/middleware/logger.go:76\ngithub.com/gin-gonic/gin.(*Context).Next\n\tE:/www/dsx/www-go/gin/vendor/github.com/gin-gonic/gin/context.go:192\ngithub.com/gin-gonic/gin.CustomRecoveryWithWriter.func1\n\tE:/www/dsx/www-go/gin/vendor/github.com/gin-gonic/gin/recovery.go:92\ngithub.com/gin-gonic/gin.(*Context).Next\n\tE:/www/dsx/www-go/gin/vendor/github.com/gin-gonic/gin/context.go:192\ngithub.com/gin-gonic/gin.LoggerWithConfig.func1\n\tE:/www/dsx/www-go/gin/vendor/github.com/gin-gonic/gin/logger.go:249\ngithub.com/gin-gonic/gin.(*Context).Next\n\tE:/www/dsx/www-go/gin/vendor/github.com/gin-gonic/gin/context.go:192\ngithub.com/gin-gonic/gin.(*Engine).handleHTTPRequest\n\tE:/www/dsx/www-go/gin/vendor/github.com/gin-gonic/gin/gin.go:689\ngithub.com/gin-gonic/gin.(*Engine).ServeHTTP\n\tE:/www/dsx/www-go/gin/vendor/github.com/gin-gonic/gin/gin.go:643\nnet/http.serverHandler.ServeHTTP\n\tE:/go-sdk/go1.25.2/src/net/http/server.go:3340\nnet/http.(*conn).serve\n\tE:/go-sdk/go1.25.2/src/net/http/server.go:2109"
}
```

# 多语言
> 使用 `i18n` 包实现多语言支持，支持 `zh` 和 `en` 两种语言, 可支持自定义扩展。语言传输默认在`header`中传输 `Accept-Language` 参数, 如 `zh` 或 `en`, 不区分大小写, 不传递默认语言为 `zh`。
## 目录配置
> 翻译文件存放路径为 `storage/locales`, 默认语言为 `zh`, 多个语言用逗号分隔。语言存放在对应的语言目录下不区分子目录, 如中文就放在`storage/locales/zh`下,可以支持任意目录下的`json`和`yaml`格式文件。
```yaml
# 翻译配置
i18n:
  dir: "storage/locales" # 翻译文件存放路径
  lang: "zh,en" # 默认语言,多个语言用逗号分隔
```

## 常规翻译
```go
import (
    "gin/utils/lang"
)

func Test()  {
    trans := lang.T("login.username", nil)
	fmt.Println(trans) // 输出: 用户名, 英文输出: Username
}
```

## 模版翻译
> 翻译文件中支持模版翻译, 如 `{{.name}}`, 使用 `map[string]interface{}` 传递参数。
ms
[
  {
    "id": "login.success",
    "translation": "{{.name}},登录成功"
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
	fmt.Println(trans) // 输出: admin,登录成功 英文输出: admin,Login Success
}
```

## 添加语言
> 在 `storage/locales` 目录下添加对应语言目录, 如 `en`, 然后在目录下添加翻译文件, 翻译文件支持 `json` 和 `yaml` 格式, 翻译文件中 `id` 为唯一标识, `translation` 为翻译内容, 翻译文件中可以添加任意数量的翻译内容。配置语言支持需调整配置文件i18n.lang参数。
```yaml
# 翻译配置
i18n:
  dir: "storage/locales" # 翻译文件存放路径
  lang: "zh,en" # 默认语言,多个语言用逗号分隔
```

# swagger文档
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
