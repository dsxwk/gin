package service

import (
	"errors"
	"gin/app/model"
	"gin/common/base"
	"gin/common/global"
	"gin/utils"
	"gorm.io/gorm"
)

type LoginService struct {
	base.BaseService
}

// Login 登录
func (s *LoginService) Login(username, password string) (m model.User, err error) {
	if err = global.DB.
		Where("username = ?", username).
		First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return m, errors.New("login.accountErr")
		}
	}

	check := utils.BcryptCheck(password, m.Password)
	if !check {
		return m, errors.New("login.pwdErr")
	}

	if m.Status == 2 {
		return m, errors.New("login.accountDisabled")
	}

	return m, nil
}
