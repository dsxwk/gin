package pkg

// Ter 三元操作符
func Ter[T any](cond bool, a, b T) T {
	if cond {
		return a
	}
	return b
}
