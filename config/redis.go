package config

import "gin/utils/cache"

func RedisInstance() *cache.RedisCache {
	return cache.NewRedis(Conf.Redis.Address, Conf.Redis.Password, Conf.Redis.DB)
}
