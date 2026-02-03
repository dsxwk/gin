package config

import (
	"gin/pkg/cache"
	"gin/pkg/message"
)

func DiskInstance() *cache.CacheProxy {
	return cache.NewDisk(Conf.Cache.Disk.Path, message.MsgEventBus)
}
