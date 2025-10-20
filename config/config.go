package config

import (
	"time"
)

// App 应用
type App struct {
	Name     string `yaml:"name"`
	Mode     string `yaml:"mode"`
	Port     int64  `yaml:"port"`
	Timezone string `yaml:"timezone"`
	Proxies  string `yaml:"proxies"`
	Env      string `yaml:"env"`
}

// Mysql 数据库
type Mysql struct {
	Host             string        `yaml:"host"`
	Port             string        `yaml:"port"`
	Database         string        `yaml:"database"`
	Username         string        `yaml:"username"`
	Password         string        `yaml:"password"`
	SlowQuerySeconds time.Duration `yaml:"slow-query-seconds"`
}

// Redis 数据库
type Redis struct {
	Address  string `yaml:"address"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
}

// Cache 缓存
type Cache struct {
	Driver string `yaml:"driver"`
}

// Cors 跨域
type Cors struct {
	Enabled          bool   `yaml:"enabled"`
	AllowOrigin      string `yaml:"allow-origin"`
	AllowHeaders     string `yaml:"allow-headers"`
	ExposeHeaders    string `yaml:"expose-headers"`
	AllowMethods     string `yaml:"allow-methods"`
	AllowCredentials string `yaml:"allow-credentials"`
}

// Jwt token
type Jwt struct {
	Key        string `yaml:"key"`
	Exp        int64  `yaml:"exp"`
	RefreshExp int64  `yaml:"refresh-exp"`
}

// Log 日志
type Log struct {
	Access     bool   `yaml:"access"`      // 是否记录访问日志
	MaxSize    int    `yaml:"max-size"`    // 单个日志文件大小（MB）
	MaxBackups int    `yaml:"max-backups"` // 最多保留的旧日志文件数
	MaxDay     int    `yaml:"max-day"`     // 保留的最大天数
	Level      string `yaml:"level"`       // 日志级别
}

// Config 配置
type Config struct {
	App
	Mysql
	Redis
	Cors
	Jwt
	Log
}
