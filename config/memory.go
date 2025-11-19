package config

import "gin/utils/cache"

func MemoryInstance() *cache.CacheProxy {
	return cache.NewMemory(Conf.Cache.Memory.DefaultExpire, Conf.Cache.Memory.CleanupInterval)
}
