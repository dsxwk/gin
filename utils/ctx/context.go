package ctx

import (
	"context"
	"github.com/gin-gonic/gin"
	"sync"
)

const (
	KeyTraceId   = "traceId"
	KeyIp        = "ip"
	KeyPath      = "path"
	KeyMethod    = "method"
	KeyParams    = "params"
	KeyMs        = "ms"
	KeyLang      = "lang"
	KeyDebugger  = "debugger"
	KeyStartTime = "startTime"
)

var (
	contextStore = make(map[string]interface{})
	mu           sync.RWMutex
	debuggerMap  sync.Map // traceId -> *Debugger
)

// Debugger 记录调试信息
type Debugger struct {
	Sql      []map[string]any
	Redis    []map[string]any
	Http     []map[string]any
	Rabbitmq []map[string]any
}

// InitDebugger 初始化调试器并绑定traceId
func InitDebugger(traceId string) *Debugger {
	dbg := &Debugger{}
	debuggerMap.Store(traceId, dbg)
	return dbg
}

// GetDebugger 通过traceId获取debugger
func GetDebugger(traceId string) *Debugger {
	if v, ok := debuggerMap.Load(traceId); ok {
		return v.(*Debugger)
	}
	return nil
}

func AddSql(ctx context.Context, data map[string]any) {
	traceId, _ := ctx.Value(KeyTraceId).(string)
	if dbg := GetDebugger(traceId); dbg != nil {
		dbg.Sql = append(dbg.Sql, data)
	}
}

func AddRedis(ctx context.Context, data map[string]any) {
	traceId, _ := ctx.Value(KeyTraceId).(string)
	if dbg := GetDebugger(traceId); dbg != nil {
		dbg.Redis = append(dbg.Redis, data)
	}
}

func AddHttp(ctx context.Context, data map[string]any) {
	traceId, _ := ctx.Value(KeyTraceId).(string)
	if dbg := GetDebugger(traceId); dbg != nil {
		dbg.Http = append(dbg.Http, data)
	}
}

func AddMq(ctx context.Context, data map[string]any) {
	traceId, _ := ctx.Value(KeyTraceId).(string)
	if dbg := GetDebugger(traceId); dbg != nil {
		dbg.Rabbitmq = append(dbg.Rabbitmq, data)
	}
}

// GetValue 通用获取context
func GetValue(ctx context.Context, key string) any {
	return ctx.Value(key)
}

// SetContext 全局gin.Context缓存(可选)
func SetContext(key string, value *gin.Context) {
	mu.Lock()
	defer mu.Unlock()
	contextStore[key] = value
}

func GetContext(key string) *gin.Context {
	mu.RLock()
	defer mu.RUnlock()
	if ctx, ok := contextStore[key]; ok {
		return ctx.(*gin.Context)
	}
	return nil
}

func ClearContext(key string) {
	mu.Lock()
	defer mu.Unlock()
	delete(contextStore, key)
}
