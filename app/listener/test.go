package listener

import (
	"encoding/json"
	"fmt"
	"gin/app/event"
	"gin/utils/eventbus"
	"time"
)

type TestListener struct{}

func (l *TestListener) Handle(e event.UserLoginEvent) {
	data, _ := json.Marshal(e)
	fmt.Printf("收到事件: %s 事件描述: %s 事件数据: %s, 时间: %s\n", e.Name(), e.Description(), data, time.Now().Format("2006-01-02 15:04:05"))
}

func init() {
	eventbus.Register(&TestListener{}, event.UserLoginEvent{})
}
