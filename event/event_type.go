package event

import (
	"github.com/fernanlukban/cronus/common"
)

type EventType struct {
	id   int
	name string
}

type EventTyper interface {
	common.IderAndNamer
	IsEventType(a_et EventTyper) bool
}

func (et EventType) IsEventType(a_et EventTyper) bool {
	return et.id == a_et.Id()
}

func (et EventType) Id() int {
	return et.id
}

func (et EventType) Name() string {
	return et.name
}
