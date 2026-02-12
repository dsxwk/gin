package config

import (
	"gin/pkg/cache"
	"sync"
)

var (
	diskCache *cache.CacheProxy
	diskOnce  sync.Once
)

func GetDiskCache() *cache.CacheProxy {
	diskOnce.Do(func() {
		diskCache = cache.NewDisk(
			Conf.Cache.Disk.Path,
		)
	})
	return diskCache
}
