package service

import (
	"errors"
	"fmt"
	"gin/app/event"
	"gin/app/middleware"
	"gin/app/model"
	"gin/app/queue/kafka/producer"
	p "gin/app/queue/rabbitmq/producer"
	"gin/common/base"
	"gin/pkg"
	"gin/pkg/container"
	"gin/pkg/eventbus"
	"gorm.io/gorm"
	"time"
)

type LoginService struct {
	base.BaseService
}

// Login 登录
func (s *LoginService) Login(username, password string) (err error, m model.User, accessToken, refreshToken string, tokenExpire, refreshTokenExpire int64) {
	containers := container.Get(s.GetContext())
	if err = containers.DB.
		Where("username = ?", username).
		First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("login.accountErr"), m, accessToken, refreshToken, tokenExpire, refreshTokenExpire
		}
	}

	check := pkg.BcryptCheck(password, m.Password)
	if !check {
		return errors.New("login.pwdErr"), m, accessToken, refreshToken, tokenExpire, refreshTokenExpire
	}

	if m.Status == 2 {
		return errors.New("login.accountDisabled"), m, accessToken, refreshToken, tokenExpire, refreshTokenExpire
	}

	jwt := middleware.Jwt{}
	accessToken, refreshToken, tokenExpire, refreshTokenExpire, err = jwt.WithRefresh(m.ID, containers.Config.Jwt.Exp, containers.Config.Jwt.RefreshExp)
	if err != nil {
		return errors.New(err.Error()), m, accessToken, refreshToken, tokenExpire, refreshTokenExpire
	}

	// 发布事件
	eventbus.Publish(s.GetContext(), event.UserLoginEvent{
		UserId:   m.ID,
		Username: m.Username,
	})

	_ = containers.RedisCache.Set("cache_test", 1, 100*time.Second)
	_ = containers.RedisCache.Set("cache_test1", 1, 100*time.Second)
	_ = containers.RedisCache.Set("redis_test", 1, 100*time.Second)
	_ = containers.RedisCache.Set("redis_test1", 1, 100*time.Second)
	_ = containers.DiskCache.Set("disk_test", 1, 100*time.Second)
	_ = containers.DiskCache.Set("disk_test1", 1, 100*time.Second)
	_ = containers.MemoryCache.Set("memory_test", 1, 100*time.Second)
	_ = containers.MemoryCache.Set("memory_test1", 1, 100*time.Second)
	_, _ = pkg.HttpRequestJson[map[string]interface{}](s.GetContext(), "GET", "http://127.0.0.1:8080/ping", nil)
	_, _ = pkg.HttpRequestJson[map[string]interface{}](s.GetContext(), "GET", "http://127.0.0.1:8080/ping", nil)
	_ = containers.RedisCache.Redis().Subscribe("testRedisChan", func(channel, payload string) {
		fmt.Println(channel, payload)
	})
	_ = containers.RedisCache.Redis().Publish("testRedisChan", map[string]interface{}{
		"test": "test",
	})

	rPub1 := p.NewRabbitmqDemoPublisher()
	defer func(rPub1 *p.RabbitmqDemoPublisher) {
		err = rPub1.Close()
		if err != nil {
			fmt.Println("kafka close error:", err)
			return
		}
	}(rPub1)
	err = rPub1.Publish(s.GetContext(), []byte(`{"orderId":333}`))
	if err != nil {
		fmt.Println("rabbitmq publish error:", err)
		return err, m, accessToken, refreshToken, tokenExpire, refreshTokenExpire
	}

	rPub2 := p.NewRabbitmqDelayDemoPublisher()
	defer func(rPub2 *p.RabbitmqDelayDemoPublisher) {
		err = rPub2.Close()
		if err != nil {
			fmt.Println("kafka close error:", err)
			return
		}
	}(rPub2)
	err = rPub2.Publish(s.GetContext(), []byte(`{"orderId":444}`))
	if err != nil {
		fmt.Println("rabbitmq publish error:", err)
		return err, m, accessToken, refreshToken, tokenExpire, refreshTokenExpire
	}

	kPub := producer.NewKafkaDemoProducer()
	defer func(kPub *producer.KafkaDemoProducer) {
		err = kPub.Close()
		if err != nil {
			fmt.Println("kafka close error:", err)
		}
	}(kPub)
	err = kPub.Publish(s.GetContext(), []byte(`{"orderId":111}`))
	if err != nil {
		fmt.Println("kafka publish error:", err)
		return err, m, accessToken, refreshToken, tokenExpire, refreshTokenExpire
	}

	kPub1 := producer.NewKafkaDelayDemoProducer()
	defer func(kPub1 *producer.KafkaDelayDemoProducer) {
		err = kPub1.Close()
		if err != nil {
			fmt.Println("kafka close error:", err)
			return
		}
	}(kPub1)
	err = kPub1.Publish(s.GetContext(), []byte(`{"orderId":222}`))
	if err != nil {
		fmt.Println("kafka publish error:", err)
		return err, m, accessToken, refreshToken, tokenExpire, refreshTokenExpire
	}

	return nil, m, accessToken, refreshToken, tokenExpire, refreshTokenExpire
}

// RefreshToken 刷新token
func (s *LoginService) RefreshToken(token string) (accessToken, refreshToken string, tExp, rExp int64, err error) {
	containers := container.Get(s.GetContext())
	j := middleware.Jwt{}
	claims, err := j.Decode(token)
	if err != nil || claims["typ"] != "refresh" {
		return accessToken, refreshToken, tExp, rExp, errors.New("login.invalidToken")
	}

	uid := int64(claims["id"].(float64))

	return j.WithRefresh(uid, containers.Config.Jwt.Exp, containers.Config.Jwt.RefreshExp)
}
