package ctx

import (
	"github.com/gin-gonic/gin"
	"sync"
)

const (
	KeyLogger  string = "logger"
	KeyTraceID string = "traceId"
	KeyParams  string = "params"
)

var (
	contextStore = make(map[string]interface{})
	mu           sync.RWMutex
)

// SetContext 设置 Context
func SetContext(key string, value *gin.Context) {
	mu.Lock()
	defer mu.Unlock()
	contextStore[key] = value
}

// GetContext 获取 Context
func GetContext(key string) *gin.Context {
	mu.RLock()
	defer mu.RUnlock()
	if ctx, ok := contextStore[key]; ok {
		return ctx.(*gin.Context)
	}
	return nil
}

// ClearContext 清理 Context
func ClearContext(key string) {
	mu.Lock()
	defer mu.Unlock()
	delete(contextStore, key)
}
