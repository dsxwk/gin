package config

import (
	"gin/pkg/cache"
	"sync"
)

var (
	instance  *cache.CacheProxy
	cacheOnce sync.Once
)

func GetCache() *cache.CacheProxy {
	cacheOnce.Do(func() {
		switch Conf.Cache.Driver {

		case "redis":
			instance = GetRedisCache()

		case "", "memory":
			instance = GetMemoryCache()

		case "disk":
			instance = GetDiskCache()

		default:
			GetLogger().Fatal("不支持的缓存驱动: " + Conf.Cache.Driver)
		}
	})

	return instance
}
