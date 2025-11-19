package debugger

const (
	TopicSql   = "debug:sql"
	TopicCache = "debug:cache"
	TopicHttp  = "debug:http"
	TopicMQ    = "debug:mq"
)

// SqlEvent Sql事件
type SqlEvent struct {
	Sql  string
	Rows int64
	Ms   float64
}

// CacheEvent 缓存事件
type CacheEvent struct {
	Driver string
	Name   string
	Cmd    string
	Args   any
	Ms     float64
}

// HttpEvent Http事件
type HttpEvent struct {
	Url      string
	Method   string
	Header   map[string]string
	Body     any
	Status   int
	Response any
	Ms       float64
}

// MqEvent Mq事件
type MqEvent struct {
	Topic string
	Body  any
	Type  string
}
