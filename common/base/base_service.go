package base

import "context"

type BaseService struct {
	Context
}

func (s *BaseService) WithContext(ctx context.Context) *BaseService {
	s.Set(ctx)

	return s
}

func (s *Context) GetContext() context.Context {
	return s.ctx
}
