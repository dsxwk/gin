## English | [中文](readme_zh.md)

- [Project Introduction](#Project Introduction)
- [License](#License)
- [Installation Instructions](#Installation Instructions)
- [Directory Structure](#Directory Structure)
- [使用方法](#使用方法)
  - [启动服务](#启动服务)
    - [air热更新](#air热更新)
  - [配置文件](#配置文件)
  - [路由](#路由)
    - [路由创建](#路由创建)
      - [路由创建帮助](#路由创建帮助)
    - [路由列表](#路由列表)
  - [控制器](#控制器)
    - [控制器创建](#控制器创建)
      - [控制器创建帮助](#控制器创建帮助)
  - [表单验证](#表单验证)
    - [验证创建](#验证创建)
      - [验证创建帮助](#验证创建帮助)
    - [验证场景](#验证场景)
    - [验证规则](#验证规则)
    - [提示信息](#提示信息)
    - [翻译](#翻译)
    - [自定义验证](#自定义验证)
  - [模型](#模型)
    - [模型创建](#模型创建)
      - [模型创建帮助](#模型创建帮助)
  - [服务](#服务)
    - [服务创建](#服务创建)
      - [服务创建帮助](#服务创建帮助)
  - [命令行](#命令行)
    - [简介](#简介)
    - [命令帮助](#命令帮助)
    - [编写命令](#编写命令)
    - [创建命令](#创建命令)
      - [创建命令帮助](#创建命令帮助)
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
> 

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
