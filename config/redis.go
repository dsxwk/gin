package config

import (
	"gin/pkg/cache"
	"sync"
)

var (
	redisCache *cache.CacheProxy
	redisOnce  sync.Once
)

func GetRedisCache() *cache.CacheProxy {
	redisOnce.Do(func() {
		redisCache = cache.NewRedis(
			Conf.Redis.Address,
			Conf.Redis.Password,
			Conf.Redis.DB,
		)
	})
	return redisCache
}
