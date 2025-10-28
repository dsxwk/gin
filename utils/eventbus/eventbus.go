package eventbus

import (
	"fmt"
	"github.com/fatih/color"
	"os"
	"sync"
)

// Event 事件接口
type Event interface {
	Name() string        // 事件名称
	Description() string // 事件描述
}

// Listener 泛型接口,处理指定事件类型
type Listener[T Event] interface {
	Handle(e T) // 处理事件
}

var (
	listenerMap      sync.Map // key: event name -> []Listener[T]
	registeredEvents sync.Map // key: event name -> struct{},用于重复检测
)

// Register 注册监听器,同时检测事件名称重复
func Register[T Event](l Listener[T], example T) {
	name := example.Name()

	// 检测事件是否重复注册
	if _, loaded := registeredEvents.LoadOrStore(name, struct{}{}); !loaded {
		color.Green(fmt.Sprintf("注册事件: %s\n", name))
	} else {
		color.Red(fmt.Sprintf("警告: 事件名称重复注册: %s\n", name))
		os.Exit(1)
	}

	var lst []Listener[T]
	if v, ok := listenerMap.Load(name); ok {
		lst = v.([]Listener[T])
	}
	lst = append(lst, l)
	listenerMap.Store(name, lst)
}

// Emit 发布事件
func Emit[T Event](e T) {
	if v, ok := listenerMap.Load(e.Name()); ok {
		for _, l := range v.([]Listener[T]) {
			go l.Handle(e)
		}
	} else {
		fmt.Printf("[EventBus] 未找到事件: %s\n", e.Name())
	}
}
