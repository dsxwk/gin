package config

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
	"strings"
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

var (
	Conf *Config
	vp   *viper.Viper
)

func InitConfig() *Config {
	v := viper.New()

	// 默认配置文件目录 ./config
	configDir := "./config"
	v.AddConfigPath(configDir)
	v.SetConfigName("config")
	v.SetConfigType("yaml")

	// 允许使用环境变量覆盖
	v.AutomaticEnv()
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// 读取主配置文件 config.yaml
	if err := v.ReadInConfig(); err != nil {
		color.Red("❌  读取配置文件失败: %v", err)
	}

	// 获取环境类型
	env := v.GetString("app.env")
	if env == "" {
		env = "dev"
	}

	// 加载对应环境的配置文件，如 config.dev.yaml
	envConfigFile := filepath.Join(configDir, fmt.Sprintf("%s.config.yaml", env))
	if _, err := os.Stat(envConfigFile); err == nil {
		v.SetConfigFile(envConfigFile)
		if err = v.MergeInConfig(); err != nil {
			color.Red("❌  合并环境配置失败: %v", err)
			os.Exit(1)
		}
		color.Green("✅  已加载环境配置文件: %s\n", envConfigFile)
	} else {
		color.Yellow("⚠️  未找到环境配置文件: %s，使用默认配置\n", envConfigFile)
	}

	// 自动映射到结构体
	cfg := &Config{}
	if err := v.Unmarshal(cfg); err != nil {
		color.Red("❌  解析配置文件失败: %v", err)
		os.Exit(1)
	}

	// 监听配置变化
	v.WatchConfig()

	var lastEventTime int64
	v.OnConfigChange(func(e fsnotify.Event) {
		if e.Op&fsnotify.Write != fsnotify.Write {
			return
		}

		now := time.Now().UnixNano()
		// 如果两次事件间隔小于200ms则忽略
		if now-lastEventTime < 200*1e6 {
			return
		}
		lastEventTime = now

		color.Green("🔄  配置文件修改: %s\n", e.Name)
		if err := v.Unmarshal(cfg); err != nil {
			color.Red("⚠️  配置热更新失败: %v", err)
			os.Exit(1)
		}
	})

	Conf = cfg
	vp = v

	return Conf
}

// Get 获取配置项
func Get(key string) interface{} {
	if vp == nil {
		InitConfig()
	}

	return vp.Get(key)
}

// GetString 快捷方法
func GetString(key string) string {
	if vp == nil {
		InitConfig()
	}

	return vp.GetString(key)
}

// GetInt 快捷方法
func GetInt(key string) int {
	if vp == nil {
		InitConfig()
	}

	return vp.GetInt(key)
}
