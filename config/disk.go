package config

import "gin/utils/cache"

func DiskInstance() *cache.DiskCache {
	instance, err := cache.NewDisk(Conf.Cache.Disk.Path)
	if err != nil {
		ZapLogger.Fatal("disk cache init failed: " + err.Error())
	}
	return instance
}
