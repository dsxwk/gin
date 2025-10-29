package base

// Listener 泛型接口,处理指定事件类型
type Listener[T Event] interface {
	Handle(e T) // 处理事件
}
