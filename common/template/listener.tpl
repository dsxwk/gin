package listener

import (
    "encoding/json"
	"fmt"
	"gin/app/event"
	"gin/utils/eventbus"
	"time"
)

type {{.Name}}Listener struct{}

func (l *{{.Name}}Listener) Handle(e event.{{.EventName}}) {
	data, _ := json.Marshal(e)
    fmt.Printf("收到事件: %s 事件描述: %s 事件数据: %s, 时间: %s\n", e.Name(), e.Description(), data, time.Now().Format("2006-01-02 15:04:05"))
}

func init() {
	eventbus.Register(&{{.Name}}Listener{}, event.{{.EventName}}{})
}
