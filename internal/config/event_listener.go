package config

import (
	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/kv"
)

type EventListener struct {
	Url     string `fig:"url"`
	Address string `fig:"address"`
	Abi     string `fig:"abi"`
}

type EventListenerer interface {
	EventListener() *EventListener
}

type eventlistener struct {
	getter kv.Getter
	once   comfig.Once
}

func NewEventListenerer(getter kv.Getter) EventListenerer {
	return &eventlistener{
		getter: getter,
	}
}

func (e *eventlistener) EventListener() *EventListener {
	return e.once.Do(func() interface{} {
		var el EventListener
		if err := figure.Out(&el).From(kv.MustGetStringMap(e.getter, "event_listener")).Please(); err != nil {
			panic(err)
		}
		return &el
	}).(*EventListener)
}
