package global

import (
	"gin/config"
	"gin/utils/cache"
	"gorm.io/gorm"
)

var (
	Config      *config.Config     // 配置
	DB          *gorm.DB           // 数据库
	Log         *config.Logger     // 日志
	Cache       cache.Cache        // 缓存
	RedisCache  *cache.RedisCache  // redis缓存
	MemoryCache *cache.MemoryCache // 内存缓存
	DiskCache   *cache.DiskCache   // 磁盘缓存
)

func init() {
	Config = config.Conf
	DB = config.DB
	Log = config.ZapLogger
	Cache = config.CacheInstance()
	RedisCache = config.RedisInstance()
	MemoryCache = config.MemoryInstance()
	DiskCache = config.DiskInstance()
}
