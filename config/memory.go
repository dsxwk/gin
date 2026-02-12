package config

import (
	"gin/pkg/cache"
	"sync"
)

var (
	memoryCache *cache.CacheProxy
	memoryOnce  sync.Once
)

func GetMemoryCache() *cache.CacheProxy {
	memoryOnce.Do(func() {
		memoryCache = cache.NewMemory(
			Conf.Cache.Memory.DefaultExpire,
			Conf.Cache.Memory.CleanupInterval,
		)
	})
	return memoryCache
}
