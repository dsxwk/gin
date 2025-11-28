package base

import (
	"gin/common/queue"
	"gin/config"
	"github.com/streadway/amqp"
	"time"
)

type RabbitmqConsumer struct {
	Mq           *RabbitMq
	Queue        string
	Exchange     string
	Routing      string
	IsDelayQueue bool
	Retry        int
}

func (c *RabbitmqConsumer) Start(h queue.Handler) {
	go func() {
		for {
			if c.Mq == nil || c.Mq.Channel == nil {
				time.Sleep(time.Second)
				continue
			}

			// 声明队列和交换机(延迟)
			args := amqp.Table{}
			exchangeType := "direct"
			if c.IsDelayQueue {
				exchangeType = "x-delayed-message"
				args["x-delayed-type"] = "direct"
			}

			if err := c.Mq.Channel.ExchangeDeclare(c.Exchange, exchangeType, true, false, false, false, args); err != nil {
				config.ZapLogger.Error("[RabbitMq] ExchangeDeclare error:" + err.Error())
				time.Sleep(time.Second)
				continue
			}

			if _, err := c.Mq.Channel.QueueDeclare(c.Queue, true, false, false, false, nil); err != nil {
				config.ZapLogger.Error("[RabbitMq] QueueDeclare error:" + err.Error())
				time.Sleep(time.Second)
				continue
			}

			if err := c.Mq.Channel.QueueBind(c.Queue, c.Routing, c.Exchange, false, nil); err != nil {
				config.ZapLogger.Error("[RabbitMq] QueueBind error:" + err.Error())
				time.Sleep(time.Second)
				continue
			}

			msgs, err := c.Mq.Channel.Consume(c.Queue, "", false, false, false, false, nil)
			if err != nil {
				config.ZapLogger.Error("[RabbitMq] Consume error:" + err.Error())
				time.Sleep(time.Second)
				continue
			}

			for msg := range msgs {
				go func(msg amqp.Delivery) {
					retry := 0
					for {
						err = h.Handle(string(msg.Body))
						if err == nil {
							err = msg.Ack(false)
							if err != nil {
								config.ZapLogger.Error("[RabbitMq] Ack error:" + err.Error())
							}
							break
						}
						retry++
						if retry >= c.Retry {
							config.ZapLogger.Error("[RabbitMq] Retry failed:" + string(msg.Body))
							err = msg.Ack(false)
							if err != nil {
								config.ZapLogger.Error("[RabbitMq] Ack error:" + err.Error())
							}
							break
						}
						time.Sleep(time.Second)
					}
				}(msg)
			}

			time.Sleep(time.Second)
		}
	}()
}
