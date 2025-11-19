package config

import (
	"gin/utils/cache"
	"gin/utils/message"
)

func DiskInstance() *cache.CacheProxy {
	return cache.NewDisk(Conf.Cache.Disk.Path, message.MsgEventBus)
}
