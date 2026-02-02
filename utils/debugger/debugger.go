package debugger

import (
	"gin/common/trace"
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
			trace.AddSql(e.TraceId, map[string]any{
				"sql":  e.Sql,
				"rows": e.Rows,
				"ms":   e.Ms,
			})
		}
	})
	d.subIds[TopicCache] = d.Bus.Subscribe(TopicCache, func(ev any) {
		if e, ok := ev.(CacheEvent); ok {
			trace.AddCache(e.TraceId, map[string]any{
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
			trace.AddHttp(e.TraceId, map[string]any{
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
	d.subIds[TopicMq] = d.Bus.Subscribe(TopicMq, func(ev any) {
		if e, ok := ev.(MqEvent); ok {
			trace.AddMq(e.TraceId, map[string]any{
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
	d.subIds[TopicListener] = d.Bus.Subscribe(TopicListener, func(ev any) {
		if e, ok := ev.(ListenerEvent); ok {
			trace.AddListener(e.TraceId, map[string]any{
				"name":  e.Name,
				"topic": e.Description,
				"data":  e.Data,
			})
		}
	})
}

func (d *Debugger) Stop() {
	for topic, id := range d.subIds {
		d.Bus.Unsubscribe(topic, id)
	}
}
