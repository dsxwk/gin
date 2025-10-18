package {{.Package}}

import (
    "errors"
    {{- if ne .Package "request" }}
    "gin/app/request"
    {{- end }}
    "github.com/gookit/validate"
)

// {{.Name}} {{.Description}}
type {{.Name}} struct {
    {{- if eq .Package "router" }}
    PageListValidate
    {{- else }}
    request.PageListValidate
    {{- end }}
}

// GetValidate 请求验证
func (s {{.Name}}) GetValidate(data {{.Name}}, scene string) error {
	v := validate.Struct(data, scene)
	if !v.Validate(scene) {
		return errors.New(v.Errors.One())
	}

	return nil
}

// ConfigValidation 配置验证
// - 定义验证场景
// - 也可以添加验证设置
func (s {{.Name}}) ConfigValidation(v *validate.Validation) {
	v.WithScenes(validate.SValues{
		"list":   []string{"PageListValidate.Page", "PageListValidate.PageSize"},
		"create": []string{},
		"update": []string{"ID"},
		"detail": []string{"ID"},
		"delete": []string{"ID"},
	})
}

// Messages 验证器错误消息
func (s {{.Name}}) Messages() map[string]string {
	return validate.MS{
		"required":    "字段 {field} 必填",
		"int":         "字段 {field} 必须为整数",
		"Page.gt":     "字段 {field} 需大于 0",
		"PageSize.gt": "字段 {field} 需大于 0",
	}
}

// Translates 字段翻译
func (s {{.Name}}) Translates() map[string]string {
	return validate.MS{
		"Page":     "页码",
		"PageSize": "每页数量",
		"ID":       "ID",
	}
}