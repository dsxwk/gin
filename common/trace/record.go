package trace

// AddSql 记录sql调试信息
func AddSql(traceId string, data map[string]any) {
	if traceId == "" {
		return
	}
	d := Store.Get(traceId)
	d.mu.Lock()
	d.Sql = append(d.Sql, data)
	d.mu.Unlock()
}

// AddCache 记录缓存调试信息
func AddCache(traceId string, data map[string]any) {
	if traceId == "" {
		return
	}
	d := Store.Get(traceId)
	d.mu.Lock()
	d.Cache = append(d.Cache, data)
	d.mu.Unlock()
}

// AddHttp 记录http调试信息
func AddHttp(traceId string, data map[string]any) {
	if traceId == "" {
		return
	}
	d := Store.Get(traceId)
	d.mu.Lock()
	d.Http = append(d.Http, data)
	d.mu.Unlock()
}

// AddMq 记录mq调试信息
func AddMq(traceId string, data map[string]any) {
	if traceId == "" {
		return
	}
	d := Store.Get(traceId)
	d.mu.Lock()
	d.Mq = append(d.Mq, data)
	d.mu.Unlock()
}

// AddListener 记录监听调试信息
func AddListener(traceId string, data map[string]any) {
	if traceId == "" {
		return
	}
	d := Store.Get(traceId)
	d.mu.Lock()
	d.ListenerEvent = append(d.ListenerEvent, data)
	d.mu.Unlock()
}
