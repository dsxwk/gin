package cache

import "time"

// Cache 缓存接口
type Cache interface {
	Set(key string, value interface{}, expire time.Duration) error // 设置缓存
	Get(key string) (interface{}, bool)                            // 获取缓存
	Delete(key string) error                                       // 删除缓存
	Expire(key string) (interface{}, time.Time, bool, error)       // 获取缓存过期时间
}
