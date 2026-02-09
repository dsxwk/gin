package middleware

import (
	"golang.org/x/time/rate"
	"sync"
	"time"
)

type limiterItem struct {
	limiter  *rate.Limiter
	lastSeen int64
}

type limiterStore struct {
	m   sync.Map
	ttl time.Duration
}

// 创建限流ttl时间不访问删除
func newLimiterStore(ttl time.Duration) *limiterStore {
	s := &limiterStore{ttl: ttl}
	go s.clean()
	return s
}

// 获取限流
func (s *limiterStore) get(key string, r rate.Limit, burst int) *rate.Limiter {
	now := time.Now().Unix()

	if v, ok := s.m.Load(key); ok {
		item := v.(*limiterItem)
		item.lastSeen = now
		return item.limiter
	}

	item := &limiterItem{
		limiter:  rate.NewLimiter(r, burst),
		lastSeen: now,
	}

	actual, _ := s.m.LoadOrStore(key, item)
	actual.(*limiterItem).lastSeen = now

	return actual.(*limiterItem).limiter
}

// 自动清理
func (s *limiterStore) clean() {
	ticker := time.NewTicker(time.Minute)

	for range ticker.C {
		expire := time.Now().Add(-s.ttl).Unix()

		s.m.Range(func(key, value any) bool {
			item := value.(*limiterItem)

			if item.lastSeen < expire {
				s.m.Delete(key)
			}
			return true
		})
	}
}
