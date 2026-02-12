package base

import "context"

type BaseMiddleware struct {
	Context
}

func (s *BaseMiddleware) GetContext() context.Context {
	return s.ctx
}
