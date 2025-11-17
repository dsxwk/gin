package message

import "sync"

type Event interface{}

type Subscriber func(event Event)

type EventBus struct {
	subscribers map[string][]Subscriber
	lock        sync.RWMutex
}

func New() *EventBus {
	return &EventBus{
		subscribers: make(map[string][]Subscriber),
	}
}

func (b *EventBus) Subscribe(topic string, sub Subscriber) {
	b.lock.Lock()
	defer b.lock.Unlock()
	b.subscribers[topic] = append(b.subscribers[topic], sub)
}

func (b *EventBus) Publish(topic string, event Event) {
	b.lock.RLock()
	defer b.lock.RUnlock()
	if subs, ok := b.subscribers[topic]; ok {
		for _, sub := range subs {
			go sub(event) // 异步
		}
	}
}
