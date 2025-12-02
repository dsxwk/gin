package {{.Package}}

import (
	"gin/common/base"
	"gin/config"
)

type {{.Name}} struct {
	{{- if eq .Type "kafka"}}
	*base.KafkaProducer
	{{- else}}
	*base.RabbitmqProducer
	{{end}}
}

func New{{.Name}}() *{{.Name}} {
	return &{{.Name}}{
		{{- if eq .Type "kafka"}}
		&base.KafkaProducer{
			Writer:       base.NewWriter(config.Conf.Kafka.Brokers, "{{.Topic}}"),
			Topic:        "{{.Topic}}",
			Key:          "{{.Key}}",
			IsDelayQueue: {{.IsDelay}},
			DelayMs:      {{.DelayMs}},
		},
		{{- else}}
		&base.RabbitmqProducer{
			Mq:           base.InitRabbitmq(),
			Queue:        "{{.Queue}}",
			Exchange:     "{{.Exchange}}",
			Routing:      "{{.Routing}}",
			IsDelayQueue: {{.IsDelay}},
			DelayMs:      {{.DelayMs}},
			Headers:      nil,
		},
		{{end}}
	}
}
