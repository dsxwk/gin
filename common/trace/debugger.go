package trace

import "sync"

type TraceDebugger struct {
	mu            sync.Mutex
	Sql           []map[string]any
	Cache         []map[string]any
	Http          []map[string]any
	Mq            []map[string]any
	ListenerEvent []map[string]any
}
