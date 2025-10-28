package event

// UserLogin 事件数据
type UserLogin struct {
	UserID   int
	Username string
}

// Name 事件名称
func (u UserLogin) Name() string {
	return "user.login"
}
