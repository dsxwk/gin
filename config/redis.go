package config

import (
	"gin/utils/cache"
	"gin/utils/message"
)

func RedisInstance() *cache.CacheProxy {
	return cache.NewRedis(Conf.Redis.Address, Conf.Redis.Password, Conf.Redis.DB, message.MsgEventBus)
}
