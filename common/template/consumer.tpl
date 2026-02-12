package {{.Package}}

import (
	"fmt"
	"gin/common/base"
	"gin/config"
)

type {{.Name}} struct {
	{{if eq .Type "kafka"}}
	*base.KafkaConsumer
	{{else}}
	*base.RabbitmqConsumer
	{{end}}
}

func New{{.Name}}() *{{.Name}} {
	c := &{{.Name}}{
		{{- if eq .Type "kafka"}}
		&base.KafkaConsumer{
			Reader:       base.NewReader(config.Conf.Kafka.Brokers, "{{.Topic}}", "{{.Group}}"),
			Topic:        "{{.Topic}}",
			Group:        "{{.Group}}",
			Retry:        {{.Retry}},
			IsDelayQueue: {{.IsDelay}},
		},
		{{- else}}
		&base.RabbitmqConsumer{
			Mq:           base.InitRabbitmq(),
			Queue:        "{{.Queue}}",
			Exchange:     "{{.Exchange}}",
			Routing:      "{{.Routing}}",
			Retry:        {{.Retry}},
			IsDelayQueue: {{.IsDelay}},
		},
		{{end}}
	}

	c.Start()

	return c
}

func (c *{{.Name}}) Start() {
	{{- if eq .Type "kafka"}}
	c.KafkaConsumer.Start(c)
	{{- else}}
	c.RabbitmqConsumer.Start(c)
	{{end}}
}

func (c *{{.Name}}) Handle(msg string) error {
	fmt.Println("Received:", msg)
	return nil
}

func init() {
	{{- if eq .Type "kafka"}}
	if config.GetConfig().Kafka.Enabled {
		New{{.Name}}()
	}
	{{- else}}
	if config.GetConfig().Rabbitmq.Enabled {
		New{{.Name}}()
	}
	{{end}}
}
