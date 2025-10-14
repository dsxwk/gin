## 中文 | [English](readme.md)

## 许可说明
- 📘 开源版：遵循 AGPL-3.0，仅供学习、研究及非商业用途。
- 💼 商业版：如需闭源或商业使用，请联系作者📧  [25076778@qq.com] 获取商业授权。

## 项目地址
- Github: https://github.com/dsxwk/gin.git
- Gitee: https://gitee.com/dsxwk/gin.git
## Gin框架介绍
Gin是一个用Go语言编写的Web框架。它具有简单、快速、高效等特点，被广泛应用于Go语言的Web开发中。

## Gin框架的特性包括：
- 快速：Gin框架基于标准库net/http，使用goroutine和channel实现异步处理，提高性能。
- 简单：Gin框架提供了一系列的API和中间件，使得开发人员可以快速构建Web应用程序。
- 高效：Gin框架使用sync.Pool来缓存对象，减少内存分配和释放，提高性能。
Golang Gin 是一个轻量级且高效的 Golang Web 框架。它具有高性能、易用性和灵活性等特点，被广泛应用于各种 Web 应用程序的开发。

## Gin项目介绍
### 命令行生成
- 命令
- 模型
- 控制器
- 服务
- 验证器
- 中间件
- 路由
- 验证器 
  - 自定义验证场景
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
- sql事件监听
- http事件监听
### 日志
- 错误信息记录
- 堆栈信息记录
- sql语句记录
- http请求记录
### 多语言
#### 目前只有登录相关模块有案例,只支持中英文,如需其他模块或语言自行扩展
- Air
- Swagger
- 测试用例
- ...

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
├── app                                 # 应用程序代码
│   ├── command                         # 命令
│   ├── controller                      # 控制期
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
go run cli.go -h # 或go run cli.go --help 查看帮助
```

### 命令创建
```shell
# 获取帮助
go run cli.go make:command -h # 或go run cli.go make:command --help 查看帮助
# 创建
go run cli.go make:command --file=cronjob/demo --desc=测试demo # 执行后会生成公共Name,如Demo::command可自行修改name
```

#### 命令注册
```go
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
```

#### 命令执行
```shell
go run cli.go Demo::command # 执行命令 Demo::command对应命令行文件自定义的name
```
```base
E:\www\dsx\www-go\gin [master +13 ~0 -0 !]> go run cli.go Demo:command       
❌  参数 --args 不能为空
Example: go run cli.go Demo:command --args=arg1 --desc=方法描述
Helper: go run cli.go Demo:command --help
exit status 1
E:\www\dsx\www-go\gin [master +14 ~0 -0 !]> go run cli.go Demo:command -h

Demo:command - demo命令行示例

Options:
  -a, --args  示例参数, 如: arg1 (参数1必填)
E:\www\dsx\www-go\gin [master +13 ~0 -0 !]> go run cli.go Demo:command -a=111
执行命令: Demo:command, 参数: 111
```

### 控制器创建
```shell
# 获取帮助
go run cli.go make:controller -h # 或go run cli.go make:controller --help 查看帮助
# 创建
go run cli.go make:controller --file=v1/user --function=list --method=get --router=/user --desc=列表 
```

### 服务创建
```shell
# 获取帮助
go run cli.go make:service -h # 或go run cli.go make:service --help 查看帮助
# 创建
go run cli.go make:service --file=v1/user --function=list --desc=列表
```

### 中间件创建
```shell
# 获取帮助
go run cli.go make:middleware -h # 或go run cli.go make:middleware --help 查看帮助
# 创建
go run cli.go make:middleware --file=jwt --desc=jwt中间件
```

### 验证器创建
```shell
# 获取帮助
go run cli.go make:request -h # 或go run cli.go make:request --help 查看帮助
# 创建
go run cli.go make:request --file=v1/user --desc=用户请求验证
```

### 路由创建
```shell
# 获取帮助
go run cli.go make:router -h # 或go run cli.go make:router --help 查看帮助
# 创建
go run cli.go make:router --file=user --desc=用户路由
```

### 路由列表
```shell
go run cli.go route:list
```
```
go run cli.go route:list
📦 当前已注册路由
---------------------------------------------------------
Method   Path                                Handler
---------------------------------------------------------
GET      /ping                               gin/router.LoadRouters
GET      /api/v1/user                        gin/app/controller/v1.(*UserController).List
POST     /api/v1/login                       gin/app/controller/v1.(*LoginController).Login
---------------------------------------------------------
✅ 总计 3 条路由
```

## Air
### 安装
```shell
go get github.com/air-go/air
# 运行
air
```

## Swagger 文档生成
```bash
# 安装
go install github.com/swaggo/swag/cmd/swag@latest
# 快速生成文档命令
swag init -g main.go --exclude cli,app/service
```

## 测试用例
```bash
go test -v ./tests/
```