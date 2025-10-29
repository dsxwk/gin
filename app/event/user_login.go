package event

// UserLoginEvent 事件数据
type UserLoginEvent struct {
	UserId   int64
	Username string
}

// Name 事件名称
func (u UserLoginEvent) Name() string {
	return "user.login"
}

// Description 事件描述
func (u UserLoginEvent) Description() string {
	return "用户登录事件"
}
