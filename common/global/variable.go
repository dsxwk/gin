package global

import (
	"gin/config"
	"gin/utils/cache"
	"gin/utils/message"
	"gorm.io/gorm"
)

var (
	Config      *config.Config        // 配置
	DB          *gorm.DB              // 数据库
	Log         *config.Logger        // 日志
	Cache       cache.Cache           // 缓存
	RedisCache  *cache.RedisCache     // redis缓存
	MemoryCache *cache.CacheProxy     // 内存缓存
	DiskCache   *cache.CacheProxy     // 磁盘缓存
	Message     = message.MsgEventBus // 消息事件
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
