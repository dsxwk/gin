package service

import (
	"errors"
	"fmt"
	"gin/app/model"
	"gin/common/base"
	"gin/common/global"
	"gin/utils"
	"gorm.io/gorm"
	"time"
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

	_ = global.RedisCache.Set("test", 1, 100*time.Second)
	_ = global.RedisCache.Set("test1", 1, 100*time.Second)
	_ = global.DiskCache.Set("test", 1, 100*time.Second)
	_ = global.DiskCache.Set("test1", 1, 100*time.Second)
	_ = global.MemoryCache.Set("test", 1, 100*time.Second)
	_ = global.MemoryCache.Set("test1", 1, 100*time.Second)
	_, _ = utils.HttpRequestJson[map[string]interface{}]("GET", "http://127.0.0.1:8080/ping", nil)
	_, _ = utils.HttpRequestJson[map[string]interface{}]("GET", "http://127.0.0.1:8080/ping", nil)
	_ = global.RedisCache.Subscribe("testChan", func(channel, payload string) {
		fmt.Println(channel, payload)
	})
	_ = global.RedisCache.Publish("testChan", map[string]interface{}{
		"test": "test",
	})

	return m, nil
}
