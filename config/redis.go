package config

import (
	"gin/pkg/cache"
	"gin/pkg/message"
)

func RedisInstance() *cache.CacheProxy {
	return cache.NewRedis(Conf.Redis.Address, Conf.Redis.Password, Conf.Redis.DB, message.MsgEventBus)
}
