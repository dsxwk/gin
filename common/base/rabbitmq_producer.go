package base

import (
	"context"
	"gin/common/ctxkey"
	"gin/config"
	"gin/pkg/debugger"
	"gin/pkg/message"
	"github.com/streadway/amqp"
	"time"
)

type RabbitMq struct {
	Conn    *amqp.Connection
	Channel *amqp.Channel
}

func NewRabbitMq(url string) (*RabbitMq, error) {
	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, err
	}
	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}
	return &RabbitMq{Conn: conn, Channel: ch}, nil
}

// Close 关闭连接
func (r *RabbitMq) Close() error {
	var err1, err2 error
	if r.Channel != nil {
		err1 = r.Channel.Close()
	}
	if r.Conn != nil {
		err2 = r.Conn.Close()
	}
	if err1 != nil {
		return err1
	}
	return err2
}

type RabbitmqProducer struct {
	Mq           *RabbitMq
	Queue        string
	Exchange     string
	Routing      string
	IsDelayQueue bool
	DelayMs      int64
	Headers      amqp.Table
}

func InitRabbitmq() *RabbitMq {
	rmq, err := NewRabbitMq(config.Conf.Rabbitmq.Url)
	if err != nil {
		config.ZapLogger.Error("RabbitMq连接失败: " + err.Error())
	}

	return rmq
}

func (p *RabbitmqProducer) initQueue() error {
	// 使用插件处理延迟
	args := amqp.Table{}
	if p.IsDelayQueue {
		// 延迟插件参数
		args["x-delayed-type"] = "direct"
	}

	// 声明交换机
	exchangeType := "direct"
	if p.IsDelayQueue {
		exchangeType = "x-delayed-message"
	}

	if err := p.Mq.Channel.ExchangeDeclare(
		p.Exchange,
		exchangeType,
		true,  // durable
		false, // autoDelete
		false,
		false,
		args,
	); err != nil {
		return err
	}

	// 声明队列并绑定
	if _, err := p.Mq.Channel.QueueDeclare(p.Queue, true, false, false, false, nil); err != nil {
		return err
	}
	return p.Mq.Channel.QueueBind(p.Queue, p.Routing, p.Exchange, false, nil)
}

func (p *RabbitmqProducer) Publish(ctx context.Context, msg []byte) error {
	start := time.Now()
	if err := p.initQueue(); err != nil {
		return err
	}

	headers := p.Headers
	if headers == nil {
		headers = amqp.Table{}
	}
	if p.IsDelayQueue && p.DelayMs > 0 {
		headers["x-delay"] = p.DelayMs
	}

	pub := amqp.Publishing{
		ContentType: "application/json",
		Body:        msg,
		Headers:     headers,
	}

	err := p.Mq.Channel.Publish(p.Exchange, p.Routing, false, false, pub)

	message.MsgEventBus.Publish(debugger.TopicMq, debugger.MqEvent{
		TraceId: ctx.Value(ctxkey.TraceIdKey).(string),
		Driver:  "rabbitmq",
		Topic:   p.Exchange + ":" + p.Routing,
		Message: string(msg),
		Key:     "",
		Group:   "",
		Ms:      float64(time.Since(start).Milliseconds()),
		Extra: map[string]any{
			"exchange": p.Exchange,
			"routing":  p.Routing,
			"err":      err,
		},
	})

	return err
}

func (p *RabbitmqProducer) Close() error {
	if p.Mq != nil {
		return p.Mq.Close()
	}
	return nil
}
