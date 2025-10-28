package event

// UserLogin 事件数据
type UserLogin struct {
	UserId   int64
	Username string
}

// Name 事件名称
func (u UserLogin) Name() string {
	return "user.login"
}

// Description 事件描述
func (u UserLogin) Description() string {
	return "用户登录事件"
}
