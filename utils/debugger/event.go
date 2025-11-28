package debugger

const (
	TopicSql   = "debug:sql"
	TopicCache = "debug:cache"
	TopicHttp  = "debug:http"
	TopicMq    = "debug:mq"
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

// MqEvent 消息队列事件
type MqEvent struct {
	TraceId string                 // tranceId
	Driver  string                 // kafka rabbitmq redis-stream
	Topic   string                 // topic queue stream
	Message string                 // 消息内容
	Key     string                 // 用于Kafka
	Group   string                 // 消费组
	Ms      float64                // 耗时ms
	Extra   map[string]interface{} // 扩展信息
}
