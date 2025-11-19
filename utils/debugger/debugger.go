package debugger

import (
	"gin/utils/ctx"
	"gin/utils/message"
)

type Debugger struct {
	Bus    *message.EventBus
	subIds map[string]uint64
}

func NewDebugger(bus *message.EventBus) *Debugger {
	return &Debugger{
		Bus:    bus,
		subIds: make(map[string]uint64),
	}
}

func (d *Debugger) Start() {
	d.subIds[TopicSql] = d.Bus.Subscribe(TopicSql, func(ev any) {
		if e, ok := ev.(SqlEvent); ok {
			ctx.AddSql(map[string]any{
				"sql":  e.Sql,
				"rows": e.Rows,
				"ms":   e.Ms,
			})
		}
	})
	d.subIds[TopicCache] = d.Bus.Subscribe(TopicCache, func(ev any) {
		if e, ok := ev.(CacheEvent); ok {
			ctx.AddCache(map[string]any{
				"driver": e.Driver,
				"name":   e.Name,
				"cmd":    e.Cmd,
				"args":   e.Args,
				"ms":     e.Ms,
			})
		}
	})
	d.subIds[TopicHttp] = d.Bus.Subscribe(TopicHttp, func(ev any) {
		if e, ok := ev.(HttpEvent); ok {
			ctx.AddHttp(map[string]any{
				"url":      e.Url,
				"method":   e.Method,
				"header":   e.Header,
				"body":     e.Body,
				"status":   e.Status,
				"response": e.Response,
				"ms":       e.Ms,
			})
		}
	})
	d.subIds[TopicMQ] = d.Bus.Subscribe(TopicMQ, func(ev any) {
		if e, ok := ev.(MqEvent); ok {
			ctx.AddMq(map[string]any{
				"driver":  e.Driver,
				"topic":   e.Topic,
				"message": e.Message,
				"key":     e.Key,
				"group":   e.Group,
				"ms":      e.Ms,
				"extra":   e.Extra,
			})
		}
	})
}

func (d *Debugger) Stop() {
	for topic, id := range d.subIds {
		d.Bus.Unsubscribe(topic, id)
	}
}
