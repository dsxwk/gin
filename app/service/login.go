package service

import (
	"errors"
	"fmt"
	"gin/app/model"
	"gin/app/queue/kafka/producer"
	p "gin/app/queue/rabbitmq/producer"
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
	kPub := producer.NewKafkaDemoProducer()
	defer func(kPub *producer.KafkaDemoProducer) {
		err = kPub.Close()
		if err != nil {

		}
	}(kPub)
	err = kPub.Publish([]byte(`{"orderId":111}`))
	if err != nil {
		fmt.Println("kafka publish error:", err)
		return m, err
	}

	kPub1 := producer.NewKafkaDelayDemoProducer()
	defer func(kPub1 *producer.KafkaDelayDemoProducer) {
		err = kPub1.Close()
		if err != nil {

		}
	}(kPub1)
	err = kPub1.Publish([]byte(`{"orderId":222}`))
	if err != nil {
		fmt.Println("kafka publish error:", err)
		return m, err
	}

	rPub1 := p.NewRabbitmqDemoPublisher()
	defer func(rPub1 *p.RabbitmqDemoPublisher) {
		err = rPub1.Close()
		if err != nil {

		}
	}(rPub1)
	err = rPub1.Publish([]byte(`{"orderId":333}`))
	if err != nil {
		fmt.Println("rabbitmq publish error:", err)
		return m, err
	}

	rPub2 := p.NewRabbitmqDelayDemoPublisher()
	defer func(rPub2 *p.RabbitmqDelayDemoPublisher) {
		err = rPub2.Close()
		if err != nil {

		}
	}(rPub2)
	err = rPub2.Publish([]byte(`{"orderId":444}`))
	if err != nil {
		fmt.Println("rabbitmq publish error:", err)
		return m, err
	}

	return m, nil
}
