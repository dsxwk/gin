package config

import "gin/pkg/cache"

func MemoryInstance() *cache.CacheProxy {
	return cache.NewMemory(Conf.Cache.Memory.DefaultExpire, Conf.Cache.Memory.CleanupInterval)
}
