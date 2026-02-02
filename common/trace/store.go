package trace

import "sync"

// traceStore 按traceId存储TraceDebugger
type traceStore struct {
	mu sync.RWMutex
	m  map[string]*TraceDebugger
}

// Store 全局Trace存储
var Store = &traceStore{
	m: make(map[string]*TraceDebugger),
}

// Get 获取或创建TraceDebugger(线程安全)
func (s *traceStore) Get(traceId string) *TraceDebugger {
	// 先读
	s.mu.RLock()
	d, ok := s.m[traceId]
	s.mu.RUnlock()
	if ok {
		return d
	}

	// 再写(双重检查)
	s.mu.Lock()
	defer s.mu.Unlock()

	if d, ok = s.m[traceId]; ok {
		return d
	}

	d = &TraceDebugger{
		Sql:   make([]map[string]any, 0),
		Cache: make([]map[string]any, 0),
		Http:  make([]map[string]any, 0),
		Mq:    make([]map[string]any, 0),
	}
	s.m[traceId] = d
	return d
}

// Delete 删除trace(防止内存泄漏)
func (s *traceStore) Delete(traceId string) {
	s.mu.Lock()
	delete(s.m, traceId)
	s.mu.Unlock()
}
