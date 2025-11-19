package cache

import (
	"errors"
	"gin/utils/message"
	"github.com/patrickmn/go-cache"
	"time"
)

// MemoryCache 内存缓存
type MemoryCache struct {
	cache *cache.Cache
}

func NewMemory(defaultExpiration, cleanupInterval time.Duration) *CacheProxy {
	mem := &MemoryCache{
		cache: cache.New(defaultExpiration, cleanupInterval),
	}
	return NewCacheProxy("memory", mem, message.MsgEventBus)
}

func (m *MemoryCache) Set(key string, value interface{}, expire time.Duration) error {
	if expire == 0 {
		expire = cache.NoExpiration
	}
	m.cache.Set(key, value, expire)
	return nil
}

func (m *MemoryCache) Get(key string) (interface{}, bool) {
	return m.cache.Get(key)
}

func (m *MemoryCache) Delete(key string) error {
	m.cache.Delete(key)
	return nil
}

func (m *MemoryCache) Expire(key string) (interface{}, time.Time, bool, error) {
	value, exp, found := m.cache.GetWithExpiration(key)
	if !found {
		return nil, time.Time{}, false, errors.New("cache key not found")
	}
	return value, exp, true, nil
}
