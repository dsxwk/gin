package queue

// Handler 消费者处理接口
type Handler interface {
	Handle(msg string) error // 消费者处理
}
