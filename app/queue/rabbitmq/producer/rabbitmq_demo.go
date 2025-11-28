package producer

import (
	"gin/common/base"
)

type RabbitmqDemoPublisher struct {
	*base.RabbitmqProducer
}

func NewRabbitmqDemoPublisher() *RabbitmqDemoPublisher {
	return &RabbitmqDemoPublisher{
		&base.RabbitmqProducer{
			Mq:           base.InitRabbitmq(),
			Queue:        "rabbitmq_demo",
			Exchange:     "rabbitmq_demo_exchange",
			Routing:      "rabbitmq_demo",
			IsDelayQueue: false,
			DelayMs:      0,
			Headers:      nil,
		},
	}
}
