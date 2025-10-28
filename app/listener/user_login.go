package listener

import (
	"fmt"
	"gin/app/event"
	"gin/utils/eventbus"
	"time"
)

type UserLoginListener struct{}

func (l *UserLoginListener) Handle(e event.UserLogin) {
	fmt.Printf("用户登录事件: %s 用户: %s, 用户ID: %d, 时间: %v\n", e.Name(), e.Username, e.UserID, time.Now().Format("2006-01-02 15:04:05"))
}

func init() {
	eventbus.Register(&UserLoginListener{}, event.UserLogin{})
}
