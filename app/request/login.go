package request

import (
	"errors"
	"github.com/gookit/validate"
)

// Login Validator
type Login struct {
	Username string `json:"username" validate:"required" label:"用户名"`
	Password string `json:"password" validate:"required" label:"密码"`
}

// GetValidate 请求验证
func (s Login) GetValidate(data Login, scene string) error {
	v := validate.Struct(data, scene)
	if !v.Validate(scene) {
		return errors.New(v.Errors.One())
	}

	return nil
}

// ConfigValidation 配置验证
// - 定义验证场景
// - 也可以添加验证设置
func (s Login) ConfigValidation(v *validate.Validation) {
	v.WithScenes(validate.SValues{
		"login": []string{"Username", "Password"},
	})
}

// Messages 您可以自定义验证器错误消息
func (s Login) Messages() map[string]string {
	return validate.MS{
		"required": "字段 {field} 必填",
	}
}

// Translates 你可以自定义字段翻译
func (s Login) Translates() map[string]string {
	return validate.MS{
		"Username": "用户名",
		"Password": "密码",
	}
}
