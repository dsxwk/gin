package producer

import (
	"gin/common/base"
	"gin/config"
)

type KafkaDemoProducer struct {
	*base.KafkaProducer
}

func NewKafkaDemoProducer() *KafkaDemoProducer {
	return &KafkaDemoProducer{
		&base.KafkaProducer{
			Writer:       base.NewWriter(config.Conf.Kafka.Brokers, "kafka_demo"),
			Topic:        "kafka_demo",
			Key:          "kafka_demo_key",
			IsDelayQueue: false,
			DelayMs:      0,
		},
	}
}
