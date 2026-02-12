package container

import (
	"context"
	"gin/common/ctxkey"
	"gin/config"
	"gin/pkg/cache"
	"gorm.io/gorm"
	"sync"
)

type Container struct {
	Config      *config.Config
	Log         *config.Logger
	DB          *gorm.DB
	Cache       *cache.CacheProxy
	RedisCache  *cache.CacheProxy
	MemoryCache *cache.CacheProxy
	DiskCache   *cache.CacheProxy
}

var (
	instance *Container
	once     sync.Once
)

func GetContainer() *Container {
	once.Do(func() {
		instance = &Container{
			Config:      config.GetConfig(),
			Log:         config.GetLogger(),
			DB:          config.GetDB(),
			Cache:       config.GetCache(),
			RedisCache:  config.GetRedisCache(),
			MemoryCache: config.GetMemoryCache(),
			DiskCache:   config.GetDiskCache(),
		}
	})
	return instance
}

func (c *Container) WithContext(ctx context.Context) *Container {
	return &Container{
		Config:      c.Config,
		Log:         c.Log,
		DB:          c.DB.WithContext(ctx),
		Cache:       c.Cache.WithContext(ctx),
		RedisCache:  c.RedisCache.WithContext(ctx),
		MemoryCache: c.MemoryCache.WithContext(ctx),
		DiskCache:   c.DiskCache.WithContext(ctx),
	}
}

// Set 保存Container到Context
func Set(ctx context.Context, c *Container) context.Context {
	return context.WithValue(ctx, ctxkey.ContainerKey, c)
}

// Get 从Context获取Container
func Get(ctx context.Context) *Container {
	if c, ok := ctx.Value(ctxkey.ContainerKey).(*Container); ok {
		return c
	}
	return nil
}
