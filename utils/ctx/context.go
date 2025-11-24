package ctx

import (
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
	KeyStartTime = "startTime"
)

var (
	contextStore = make(map[string]interface{})
	mu           sync.RWMutex
	debuggers    sync.Map // traceId -> *Debugger
)

// Debugger 收集调试信息（Sql、Cache、Http、Mq）
type Debugger struct {
	Sql   []map[string]any
	Cache []map[string]any
	Http  []map[string]any
	Mq    []map[string]any
}

func TraceId() string {
	return GetContext(KeyTraceId).GetString(KeyTraceId)
}

// InitDebugger 初始化
func InitDebugger(traceId string) *Debugger {
	dbg := &Debugger{}
	debuggers.Store(traceId, dbg)
	return dbg
}

func GetDebugger(traceId string) *Debugger {
	if v, ok := debuggers.Load(traceId); ok {
		return v.(*Debugger)
	}
	return nil
}

// AddSql 添加Sql记录
func AddSql(traceId string, data map[string]any) {
	if dbg := GetDebugger(traceId); dbg != nil {
		dbg.Sql = append(dbg.Sql, data)
	}
}

// AddCache 添加Cache记录
func AddCache(traceId string, data map[string]any) {
	if dbg := GetDebugger(traceId); dbg != nil {
		dbg.Sql = append(dbg.Sql, data)
	}
}

// AddHttp 添加Http记录
func AddHttp(traceId string, data map[string]any) {
	if dbg := GetDebugger(traceId); dbg != nil {
		dbg.Http = append(dbg.Http, data)
	}
}

// AddMq 添加Mq记录
func AddMq(traceId string, data map[string]any) {
	if dbg := GetDebugger(traceId); dbg != nil {
		dbg.Mq = append(dbg.Mq, data)
	}
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
