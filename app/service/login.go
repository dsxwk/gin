package service

import (
	"errors"
	"fmt"
	"gin/app/middleware"
	"gin/app/model"
	"gin/app/queue/kafka/producer"
	p "gin/app/queue/rabbitmq/producer"
	"gin/common/base"
	"gin/common/global"
	"gin/pkg"
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

	check := pkg.BcryptCheck(password, m.Password)
	if !check {
		return m, errors.New("login.pwdErr")
	}

	if m.Status == 2 {
		return m, errors.New("login.accountDisabled")
	}

	_ = global.RedisCache.Set("cache_test", 1, 100*time.Second)
	_ = global.RedisCache.Set("cache_test1", 1, 100*time.Second)
	_ = global.RedisCache.Set("redis_test", 1, 100*time.Second)
	_ = global.RedisCache.Set("redis_test1", 1, 100*time.Second)
	_ = global.DiskCache.Set("disk_test", 1, 100*time.Second)
	_ = global.DiskCache.Set("disk_test1", 1, 100*time.Second)
	_ = global.MemoryCache.Set("memory_test", 1, 100*time.Second)
	_ = global.MemoryCache.Set("memory_test1", 1, 100*time.Second)
	_, _ = pkg.HttpRequestJson[map[string]interface{}](s.Context.Get(), "GET", "http://127.0.0.1:8080/ping", nil)
	_, _ = pkg.HttpRequestJson[map[string]interface{}](s.Context.Get(), "GET", "http://127.0.0.1:8080/ping", nil)
	_ = global.RedisCache.Redis().Subscribe("testRedisChan", func(channel, payload string) {
		fmt.Println(channel, payload)
	})
	_ = global.RedisCache.Redis().Publish("testRedisChan", map[string]interface{}{
		"test": "test",
	})
	kPub := producer.NewKafkaDemoProducer()
	defer func(kPub *producer.KafkaDemoProducer) {
		err = kPub.Close()
		if err != nil {
			fmt.Println("kafka close error:", err)
		}
	}(kPub)
	err = kPub.Publish(s.Context.Get(), []byte(`{"orderId":111}`))
	if err != nil {
		fmt.Println("kafka publish error:", err)
		return m, err
	}

	kPub1 := producer.NewKafkaDelayDemoProducer()
	defer func(kPub1 *producer.KafkaDelayDemoProducer) {
		err = kPub1.Close()
		if err != nil {
			fmt.Println("kafka close error:", err)
			return
		}
	}(kPub1)
	err = kPub1.Publish(s.Context.Get(), []byte(`{"orderId":222}`))
	if err != nil {
		fmt.Println("kafka publish error:", err)
		return m, err
	}

	rPub1 := p.NewRabbitmqDemoPublisher()
	defer func(rPub1 *p.RabbitmqDemoPublisher) {
		err = rPub1.Close()
		if err != nil {
			fmt.Println("kafka close error:", err)
			return
		}
	}(rPub1)
	err = rPub1.Publish(s.Context.Get(), []byte(`{"orderId":333}`))
	if err != nil {
		fmt.Println("rabbitmq publish error:", err)
		return m, err
	}

	rPub2 := p.NewRabbitmqDelayDemoPublisher()
	defer func(rPub2 *p.RabbitmqDelayDemoPublisher) {
		err = rPub2.Close()
		if err != nil {
			fmt.Println("kafka close error:", err)
			return
		}
	}(rPub2)
	err = rPub2.Publish(s.Context.Get(), []byte(`{"orderId":444}`))
	if err != nil {
		fmt.Println("rabbitmq publish error:", err)
		return m, err
	}

	return m, nil
}

// RefreshToken 刷新token
func (s *LoginService) RefreshToken(token string) (accessToken, refreshToken string, tExp, rExp int64, err error) {
	j := middleware.Jwt{}
	claims, err := j.Decode(token)
	if err != nil || claims["typ"] != "refresh" {
		return accessToken, refreshToken, tExp, rExp, errors.New("login.invalidToken")
	}

	uid := int64(claims["id"].(float64))

	return j.WithRefresh(uid, global.Config.Jwt.Exp, global.Config.Jwt.RefreshExp)
}
