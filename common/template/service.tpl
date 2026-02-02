package {{.Package}}

import (
    "gin/common/base"
    "context"
)

type {{.Name}}Service struct {
    base.BaseService
}

// {{.Function}} {{.Description}}
func (s *{{.Name}}Service) {{.Function}}(ctx context.Context) {
    // Define your function here
}