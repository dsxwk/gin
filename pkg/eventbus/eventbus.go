package eventbus

import (
	"context"
	"fmt"
	"gin/common/base"
	"gin/common/ctxkey"
	"gin/pkg/debugger"
	"gin/pkg/message"
	"github.com/fatih/color"
	"sync"
)

type EventInfo struct {
	Name        string
	Description string
	Listeners   []string
}

var (
	listenerMap sync.Map // key: event name -> []base.Listener[T]
	eventInfos  sync.Map // key: event name -> EventInfo
)

// Register 注册监听
func Register[T base.Event](listener base.Listener[T], event T) {
	name := event.Name()
	desc := event.Description()

	// 获取当前已注册监听
	var listen []base.Listener[T]
	if v, ok := listenerMap.Load(name); ok {
		listen = v.([]base.Listener[T])
	}

	// 添加新的监听
	listen = append(listen, listener)
	listenerMap.Store(name, listen)

	// 更新事件信息
	info := EventInfo{
		Name:        name,
		Description: desc,
	}
	if v, ok := eventInfos.Load(name); ok {
		existing := v.(EventInfo)
		existing.Listeners = append(existing.Listeners, fmt.Sprintf("%T", listener))
		eventInfos.Store(name, existing)
	} else {
		info.Listeners = []string{fmt.Sprintf("%T", listener)}
		eventInfos.Store(name, info)
		// color.Green("注册事件: %s (%s)", name, desc)
	}
}

// Publish 发布事件
func Publish[T base.Event](ctx context.Context, e T) {
	message.MsgEventBus.Publish(debugger.TopicListener, debugger.ListenerEvent{
		TraceId:     ctx.Value(ctxkey.TraceIdKey).(string),
		Name:        e.Name(),
		Description: e.Description(),
		Data:        e,
	})

	if v, ok := listenerMap.Load(e.Name()); ok {
		for _, listener := range v.([]base.Listener[T]) {
			go listener.Handle(e)
		}
	} else {
		color.Yellow("未找到事件监听: %s", e.Name())
	}
}

// EventList 已注册事件列表
func EventList() []EventInfo {
	var list []EventInfo
	eventInfos.Range(func(_, value any) bool {
		list = append(list, value.(EventInfo))
		return true
	})
	return list
}

// DebugPrint 打印所有注册事件信息
func DebugPrint() {
	color.Cyan("==== 当前已注册事件 ====")
	eventInfos.Range(func(_, value any) bool {
		info := value.(EventInfo)
		fmt.Printf("事件: %s\n描述: %s\n监听:\n", info.Name, info.Description)
		for _, l := range info.Listeners {
			fmt.Printf("  - %s\n", l)
		}
		fmt.Println("----------------------")
		return true
	})
}
