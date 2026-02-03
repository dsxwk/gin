package config

import (
	"gin/pkg/cache"
)

func CacheInstance() *cache.CacheProxy {
	var instance *cache.CacheProxy

	switch Conf.Cache.Driver {
	case "redis":
		instance = RedisInstance()
	case "", "memory":
		instance = MemoryInstance()
	case "disk":
		instance = DiskInstance()
	default:
		ZapLogger.Fatal("不支持的缓存驱动: " + Conf.Cache.Driver)
	}

	return instance
}
