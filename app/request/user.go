package request

import (
	"errors"
	"github.com/gookit/validate"
)

// User 用户请求z验证
type User struct {
	PageValidate
}

// GetValidate 请求验证
func (s User) GetValidate(data User, scene string) error {
	v := validate.Struct(data, scene)
	if !v.Validate(scene) {
		return errors.New(v.Errors.One())
	}

	return nil
}

// ConfigValidation 配置验证
// - 定义验证场景
// - 也可以添加验证设置
func (s User) ConfigValidation(v *validate.Validation) {
	v.WithScenes(validate.SValues{
		"list":   []string{"Page", "PageSize"},
		"create": []string{},
		"update": []string{"ID"},
		"detail": []string{"ID"},
		"delete": []string{"ID"},
	})
}

// Messages 您可以自定义验证器错误消息
func (s User) Messages() map[string]string {
	return validate.MS{
		"required":    "字段 {field} 必填",
		"int":         "字段 {field} 必须为整数",
		"Page.gt":     "字段 {field} 需大于 0",
		"PageSize.gt": "字段 {field} 需大于 0",
	}
}

// Translates 你可以自定义字段翻译
func (s User) Translates() map[string]string {
	return validate.MS{
		"Page":     "页码",
		"PageSize": "每页数量",
		"ID":       "ID",
	}
}
