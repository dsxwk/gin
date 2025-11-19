package cache

import (
	"gin/utils/debugger"
	"gin/utils/message"
	"time"
)

// Cache 缓存接口
type Cache interface {
	Set(key string, value interface{}, expire time.Duration) error // 设置缓存
	Get(key string) (interface{}, bool)                            // 获取缓存
	Delete(key string) error                                       // 删除缓存
	Expire(key string) (interface{}, time.Time, bool, error)       // 获取缓存过期时间
}

type CacheProxy struct {
	driver string
	c      Cache
	bus    *message.EventBus
}

func NewCacheProxy(driver string, c Cache, bus *message.EventBus) *CacheProxy {
	return &CacheProxy{
		driver: driver,
		c:      c,
		bus:    bus,
	}
}

func (p *CacheProxy) Set(key string, value interface{}, expire time.Duration) error {
	start := time.Now()
	err := p.c.Set(key, value, expire)
	p.publish("Set", key, value, time.Since(start))
	return err
}

func (p *CacheProxy) Get(key string) (interface{}, bool) {
	start := time.Now()
	val, ok := p.c.Get(key)
	p.publish("Get", key, val, time.Since(start))
	return val, ok
}

func (p *CacheProxy) Delete(key string) error {
	start := time.Now()
	err := p.c.Delete(key)
	p.publish("Delete", key, nil, time.Since(start))
	return err
}

func (p *CacheProxy) Expire(key string) (interface{}, time.Time, bool, error) {
	start := time.Now()
	val, exp, ok, err := p.c.Expire(key)
	p.publish("Expire", key, val, time.Since(start))
	return val, exp, ok, err
}

func (p *CacheProxy) publish(method, key string, val interface{}, cost time.Duration) {
	if p.bus != nil {
		p.bus.Publish(debugger.TopicCache, debugger.CacheEvent{
			Driver: p.driver,
			Name:   method,
			Cmd:    key,
			Args:   val,
			Ms:     float64(cost.Nanoseconds()) / 1e6,
		})
	}
}
