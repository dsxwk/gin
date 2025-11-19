package ctx

import (
	"context"
	"github.com/gin-gonic/gin"
	"runtime"
	"strconv"
	"strings"
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
	traceMap     sync.Map // goroutineId -> traceId
	debuggers    sync.Map // traceId -> *Debugger
)

// Debugger 记录调试信息
type Debugger struct {
	Sql      []map[string]any
	Redis    []map[string]any
	Http     []map[string]any
	Rabbitmq []map[string]any
}

// 获取当前goroutine Id
func goId() int64 {
	buf := make([]byte, 64)
	n := runtime.Stack(buf, false)
	// goroutine 18 [running]:
	parts := strings.Split(string(buf[:n]), " ")
	id, _ := strconv.ParseInt(parts[1], 10, 64)
	return id
}

// BindTraceId 中间件设置：绑定 goroutine → traceId
func BindTraceId(traceId string) {
	traceMap.Store(goId(), traceId)
}

// CurrentTrace 获取当前traceId
func CurrentTrace() string {
	id := goId()
	v, ok := traceMap.Load(id)
	if !ok {
		return ""
	}
	return v.(string)
}

// GetDebugger 获取debugger
func GetDebugger() *Debugger {
	traceId := CurrentTrace()
	if traceId == "" {
		return nil
	}
	v, _ := debuggers.LoadOrStore(traceId, &Debugger{})
	return v.(*Debugger)
}

// Clear 清除
func Clear() {
	id := goId()
	v, ok := traceMap.Load(id)
	if ok {
		debuggers.Delete(v.(string))
		traceMap.Delete(id)
	}
}

func AddSql(data map[string]any) {
	if dbg := GetDebugger(); dbg != nil {
		dbg.Sql = append(dbg.Sql, data)
	}
}

func AddRedis(data map[string]any) {
	if dbg := GetDebugger(); dbg != nil {
		dbg.Redis = append(dbg.Redis, data)
	}
}

func AddHttp(data map[string]any) {
	if dbg := GetDebugger(); dbg != nil {
		dbg.Http = append(dbg.Http, data)
	}
}

func AddRabbitmq(data map[string]any) {
	if dbg := GetDebugger(); dbg != nil {
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
