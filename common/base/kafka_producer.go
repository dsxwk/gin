package base

import (
	"context"
	"encoding/json"
	"gin/common/ctxkey"
	"gin/pkg/debugger"
	"gin/pkg/message"
	"time"

	"github.com/segmentio/kafka-go"
)

type KafkaProducer struct {
	Writer       *kafka.Writer
	Topic        string
	Key          string
	IsDelayQueue bool
	DelayMs      int64
}

func NewWriter(brokers []string, topic string) *kafka.Writer {
	return &kafka.Writer{
		Addr:         kafka.TCP(brokers...),
		Topic:        topic,
		Balancer:     &kafka.LeastBytes{},
		RequiredAcks: kafka.RequireAll,
	}
}

// Publish 发送消息
func (p *KafkaProducer) Publish(ctx context.Context, msg []byte) error {
	start := time.Now()

	// 延迟队列模拟
	if p.IsDelayQueue {
		msgMap := map[string]any{
			"body":      string(msg),
			"publishAt": time.Now().Add(time.Millisecond * time.Duration(p.DelayMs)).UnixMilli(),
		}
		msg, _ = json.Marshal(msgMap)
	}

	kmsg := kafka.Message{
		Value: msg, // 不指定Topic
	}

	if p.Key != "" {
		kmsg.Key = []byte(p.Key)
	}

	err := p.Writer.WriteMessages(context.Background(), kmsg)

	message.MsgEventBus.Publish(debugger.TopicMq, debugger.MqEvent{
		TraceId: ctx.Value(ctxkey.TraceIdKey).(string),
		Driver:  "kafka",
		Topic:   p.Topic,
		Message: string(msg),
		Ms:      float64(time.Since(start).Milliseconds()),
		Extra:   map[string]interface{}{"err": err},
	})

	return err
}

// Close 关闭Writer
func (p *KafkaProducer) Close() error {
	return p.Writer.Close()
}
