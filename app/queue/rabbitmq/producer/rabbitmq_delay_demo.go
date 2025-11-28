package producer

import (
	"gin/common/base"
)

type RabbitmqDelayDemoPublisher struct {
	*base.RabbitmqProducer
}

func NewRabbitmqDelayDemoPublisher() *RabbitmqDelayDemoPublisher {
	return &RabbitmqDelayDemoPublisher{
		&base.RabbitmqProducer{
			Mq:           base.InitRabbitmq(),
			Queue:        "rabbitmq_delay_demo",
			Exchange:     "rabbitmq_delay_demo_exchange",
			Routing:      "rabbitmq_delay_demo",
			IsDelayQueue: true,
			DelayMs:      10000,
			Headers:      nil,
		},
	}
}
