package config

import "gin/utils/cache"

func MemoryInstance() *cache.MemoryCache {
	return cache.NewMemory(Conf.Cache.Memory.DefaultExpire, Conf.Cache.Memory.CleanupInterval)
}
