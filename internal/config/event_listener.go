package config

import (
	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/kv"
)

type EventListener struct {
	Url     string
	Address string
	Abi     string
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

func (t *eventlistener) EventListener() *EventListener {
	return t.once.Do(func() interface{} {
		var config struct {
			Url     string `fig:"url"`
			Address string `fig:"address"`
			Abi     string `fig:"abi"`
		}
		if err := figure.Out(&config).From(kv.MustGetStringMap(t.getter, "event_listener")).Please(); err != nil {
			panic(err)
		}
		return &EventListener{Address: config.Address, Url: config.Url, Abi: config.Abi}
	}).(*EventListener)
}
