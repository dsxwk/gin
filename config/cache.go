package config

import (
	"gin/utils/cache"
)

func CacheInstance() cache.Cache {
	var (
		instance cache.Cache
	)

	switch Conf.Cache.Driver {
	case "redis":
		instance = RedisInstance()
	case "", "memory":
		instance = MemoryInstance()
	case "disk":
		instance = DiskInstance()
	default:
		ZapLogger.Fatal(nil, "不支持的缓存驱动: "+Conf.Cache.Driver)
	}

	return instance
}
