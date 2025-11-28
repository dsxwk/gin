package config

import (
	"fmt"
	"gin/utils"
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
	Name     string `mapstructure:"name" yaml:"name"`
	Mode     string `mapstructure:"mode" yaml:"mode"`
	Port     int64  `mapstructure:"port" yaml:"port"`
	Timezone string `mapstructure:"timezone" yaml:"timezone"`
	Proxies  string `mapstructure:"proxies" yaml:"proxies"`
	Env      string `mapstructure:"env" yaml:"env"`
}

// Mysql 数据库
type Mysql struct {
	Host             string        `mapstructure:"host" yaml:"host"`
	Port             string        `mapstructure:"port" yaml:"port"`
	Database         string        `mapstructure:"database" yaml:"database"`
	Username         string        `mapstructure:"username" yaml:"username"`
	Password         string        `mapstructure:"password" yaml:"password"`
	SlowQuerySeconds time.Duration `mapstructure:"slow-query-seconds" yaml:"slow-query-seconds"`
}

// Cache 缓存
type Cache struct {
	Driver string `mapstructure:"driver" yaml:"driver"`
	Redis  Redis  `mapstructure:"redis" yaml:"redis"`
	Memory Memory `mapstructure:"memory" yaml:"memory"`
	Disk   Disk   `mapstructure:"disk" yaml:"disk"`
}

// Redis 数据库
type Redis struct {
	Address  string `mapstructure:"address" yaml:"address"`
	Password string `mapstructure:"password" yaml:"password"`
	DB       int    `mapstructure:"db" yaml:"db"`
}

// Memory 内存缓存
type Memory struct {
	DefaultExpire   time.Duration `mapstructure:"default-expire" yaml:"default-expire"`
	CleanupInterval time.Duration `mapstructure:"cleanup-interval" yaml:"cleanup-interval"`
}

// Disk 磁盘缓存
type Disk struct {
	Path string `mapstructure:"path" yaml:"path"`
}

// Cors 跨域
type Cors struct {
	Enabled          bool   `mapstructure:"enabled" yaml:"enabled"`
	AllowOrigin      string `mapstructure:"allow-origin" yaml:"allow-origin"`
	AllowHeaders     string `mapstructure:"allow-headers" yaml:"allow-headers"`
	ExposeHeaders    string `mapstructure:"expose-headers" yaml:"expose-headers"`
	AllowMethods     string `mapstructure:"allow-methods" yaml:"allow-methods"`
	AllowCredentials string `mapstructure:"allow-credentials" yaml:"allow-credentials"`
}

// Jwt token
type Jwt struct {
	Key        string `mapstructure:"key" yaml:"key"`
	Exp        int64  `mapstructure:"exp" yaml:"exp"`
	RefreshExp int64  `mapstructure:"refresh-exp" yaml:"refresh-exp"`
}

// Log 日志
type Log struct {
	Access     bool   `mapstructure:"access" yaml:"access"`           // 是否记录访问日志
	MaxSize    int    `mapstructure:"max-size" yaml:"max-size"`       // 单个日志文件大小（MB）
	MaxBackups int    `mapstructure:"max-backups" yaml:"max-backups"` // 最多保留的旧日志文件数
	MaxDay     int    `mapstructure:"max-day" yaml:"max-day"`         // 保留的最大天数
	Level      string `mapstructure:"level" yaml:"level"`             // 日志级别
}

// I18n 翻译
type I18n struct {
	Dir  string `mapstructure:"dir" yaml:"dir"`   // 翻译文件目录
	Lang string `mapstructure:"lang" yaml:"lang"` // 默认语言,多个语言用逗号分隔
}

type Kafka struct {
	Enabled bool     `mapstructure:"enabled" yaml:"enabled"` // 是否启用
	Brokers []string `mapstructure:"brokers" yaml:"brokers"`
}

type Rabbitmq struct {
	Enabled bool   `mapstructure:"enabled" yaml:"enabled"` // 是否启用
	Url     string `mapstructure:"url" yaml:"url"`
}

// Config 配置
type Config struct {
	App      App      `mapstructure:"app" yaml:"app"`
	Mysql    Mysql    `mapstructure:"mysql" yaml:"mysql"`
	Redis    Redis    `mapstructure:"redis" yaml:"redis"`
	Cors     Cors     `mapstructure:"cors" yaml:"cors"`
	Jwt      Jwt      `mapstructure:"jwt" yaml:"jwt"`
	Log      Log      `mapstructure:"log" yaml:"log"`
	Cache    Cache    `mapstructure:"cache" yaml:"cache"`
	I18n     I18n     `mapstructure:"i18n" yaml:"i18n"`
	Kafka    Kafka    `mapstructure:"kafka" yaml:"kafka"`
	Rabbitmq Rabbitmq `mapstructure:"rabbitmq" yaml:"rabbitmq"`
}

var (
	Conf *Config
	vp   *viper.Viper
)

func init() {
	v := viper.New()

	// 默认配置文件目录为根目录
	configDir := utils.GetRootPath()
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
	configFile := filepath.Join(configDir, fmt.Sprintf("%s.config.yaml", env))
	if _, err := os.Stat(configFile); err == nil {
		v.SetConfigFile(configFile)
		if err = v.MergeInConfig(); err != nil {
			color.Red("❌  合并环境配置失败: %v", err)
			os.Exit(1)
		}
		color.Green("✅  已加载环境配置文件: %s\n", configFile)
	} else {
		color.Yellow("⚠️  未找到环境配置文件: %s，使用默认配置\n", configFile)
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
}

// Get 获取配置项
func Get(key string) interface{} {
	return vp.Get(key)
}

// GetString 快捷方法
func GetString(key string) string {
	return vp.GetString(key)
}

// GetInt 快捷方法
func GetInt(key string) int {
	return vp.GetInt(key)
}
