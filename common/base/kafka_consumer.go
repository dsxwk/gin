package base

import (
	"context"
	"gin/config"
	"gin/pkg/queue"
	"github.com/goccy/go-json"
	"time"

	"github.com/segmentio/kafka-go"
)

type KafkaConsumer struct {
	Reader       *kafka.Reader
	Topic        string
	Group        string
	Retry        int
	IsDelayQueue bool
}

func NewReader(brokers []string, topic, group string) *kafka.Reader {
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:        brokers,
		Topic:          topic,
		GroupID:        group,
		MinBytes:       1,
		MaxBytes:       10e6,
		StartOffset:    kafka.LastOffset,
		CommitInterval: 0, // 关闭自动提交
	})
}

// Start 启动消费(循环读取+延迟消息+重试)
func (c *KafkaConsumer) Start(h queue.Handler) {
	go func() {
		for {
			msg, err := c.Reader.ReadMessage(context.Background())
			if err != nil {
				config.GetLogger().Error("kafka read error:" + err.Error())
				time.Sleep(time.Second)
				continue
			}

			var actualMsg string
			if c.IsDelayQueue {
				// 解析延迟消息
				var msgMap map[string]any
				if err = json.Unmarshal(msg.Value, &msgMap); err != nil {
					config.GetLogger().Error("kafka delay msg unmarshal error:" + err.Error())
					continue
				}

				publishAt := int64(msgMap["publishAt"].(float64))
				now := time.Now().UnixMilli()
				if now < publishAt {
					// 还没到消费时间,放回延迟或sleep
					sleepMs := publishAt - now
					time.Sleep(time.Millisecond * time.Duration(sleepMs))
				}

				actualMsg = msgMap["body"].(string)
			} else {
				actualMsg = string(msg.Value)
			}

			// 重试逻辑
			attempt := 0
			var handleErr error
			for {
				handleErr = h.Handle(actualMsg)
				if handleErr == nil {
					// 提交确认
					err = c.Reader.CommitMessages(context.Background(), msg)
					if err != nil {
						config.GetLogger().Error("kafka commit error:" + err.Error())
					}
					break
				}
				attempt++
				if attempt >= c.Retry {
					config.GetLogger().Error("kafka retry failed:" + actualMsg)
					break
				}
				time.Sleep(time.Second)
			}
		}
	}()
}
