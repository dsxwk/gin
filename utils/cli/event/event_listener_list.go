package event

import (
	"gin/common/base"
	"gin/utils/cli"
	"gin/utils/eventbus"
)

type EventListenerList struct{}

func (s *EventListenerList) Name() string {
	return "event-listener:list"
}

func (s *EventListenerList) Description() string {
	return "事件监听列表"
}

func (s *EventListenerList) Help() []base.CommandOption {
	return []base.CommandOption{}
}

func (s *EventListenerList) Execute(args []string) {
	eventbus.DebugPrint()
}

func init() {
	cli.Register(&EventListenerList{})
}
