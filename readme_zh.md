## 中文 | [English](readme.md)

- [项目简介](#项目简介)
- [许可证](#许可证)
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
