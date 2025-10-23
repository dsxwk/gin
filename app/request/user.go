package request

import (
	"errors"
	"github.com/gookit/validate"
)

// UserCreate 用户创建验证
type UserCreate struct {
	Username string `json:"username" validate:"required" label:"用户名"`
	FullName string `json:"fullName" validate:"required" label:"姓名"`
	Nickname string `json:"nickname" validate:"required" label:"昵称"`
	Gender   int    `json:"gender" validate:"required|int" label:"性别"`
	Password string `json:"password" validate:"required" label:"密码"`
}

// UserUpdate 用户更新验证
type UserUpdate struct {
	UserDetail
	UserCreate
}

// UserDetail 用户详情验证
type UserDetail struct {
	ID int64 `json:"id" validate:"required|int|gt:0" label:"ID"`
}

// UserSearch 用户搜索
type UserSearch struct {
	Username string `form:"username" validate:"required" label:"用户名"`
	FullName string `form:"fullName" validate:"required" label:"姓名"`
	Nickname string `form:"nickname" validate:"required" label:"昵称"`
	Gender   int    `form:"gender" validate:"required|int" label:"性别"`
}

// User 用户请求验证
type User struct {
	UserDetail
	UserCreate
	PageListValidate
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
		// 列表
		"List": []string{
			"PageListValidate.Page",
			"PageListValidate.PageSize",
		},
		// 创建
		"Create": []string{
			"UserCreate.Username",
			"UserCreate.FullName",
			"UserCreate.Nickname",
			"UserCreate.Gender",
			"UserCreate.Password",
		},
		// 更新
		"Update": []string{
			"UserUpdate.UserDetail.ID",
			"UserCreate.Username",
			"UserCreate.FullName",
			"UserCreate.Nickname",
			"UserCreate.Gender",
		},
		// 详情
		"Detail": []string{
			"UserDetail.ID",
		},
		// 删除
		"Delete": []string{
			"UserDetail.ID",
		},
	})
}

// Messages 验证器错误消息
func (s User) Messages() map[string]string {
	return validate.MS{
		"required":    "字段 {field} 必填",
		"int":         "字段 {field} 必须为整数",
		"Page.gt":     "字段 {field} 需大于 0",
		"PageSize.gt": "字段 {field} 需大于 0",
	}
}

// Translates 字段翻译
func (s User) Translates() map[string]string {
	return validate.MS{
		"Page":                "页码",
		"PageSize":            "每页数量",
		"ID":                  "ID",
		"UserCreate.Username": "用户名",
		"UserCreate.FullName": "姓名",
		"UserCreate.Nickname": "昵称",
		"UserCreate.Gender":   "性别",
		"UserCreate.Password": "密码",
	}
}
