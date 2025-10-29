package eventbus

import (
	"fmt"
	"gin/common/base"
	"github.com/fatih/color"
	"os"
	"sync"
)

var (
	listenerMap      sync.Map // key: event name -> []Listener[T]
	registeredEvents sync.Map // key: event name -> struct{},用于重复检测
)

// Register 注册监听
func Register[T base.Event](l base.Listener[T], e T) {
	name := e.Name()

	// 检测事件是否重复注册
	if _, loaded := registeredEvents.LoadOrStore(name, struct{}{}); !loaded {
		color.Green(fmt.Sprintf("注册事件: %s, 事件描述: %s\n", e.Name(), e.Description()))
	} else {
		color.Red(fmt.Sprintf("事件已注册: %s\n", e.Name()))
		os.Exit(1)
	}

	var lsn []base.Listener[T]
	if v, ok := listenerMap.Load(name); ok {
		lsn = v.([]base.Listener[T])
	}
	lsn = append(lsn, l)
	listenerMap.Store(name, lsn)
}

// Publish 发布事件
func Publish[T base.Event](e T) {
	if v, ok := listenerMap.Load(e.Name()); ok {
		for _, l := range v.([]base.Listener[T]) {
			go l.Handle(e)
		}
	} else {
		fmt.Printf("未找到事件: %s\n", e.Name())
	}
}
