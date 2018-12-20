package event

import (
	"time"
)

type Event struct {
	eventDetails EventDetails
	eventTypes   map[int]EventTyper
	owner        int
}

type EventDetails struct {
	StartTime   time.Time
	EndTime     time.Time
	Title       string
	Description string
	Location    string
}

type Eventer interface {
	TotalDuration() time.Time
	IsEventType(a_et EventTyper) bool
	StartTime() time.Time
	EndTime() time.Time
	Title() string
	Description() string
	Location() string
	EventTypes() map[int]EventTyper
}

func (e Event) TotalDuration() time.Time {
	var time_difference time.Duration = e.EndTime().Sub(e.StartTime())

	// Translates time_difference into a time.Time object
	total_duration := time.Time{}.Add(time_difference)

	return total_duration
}

func (e Event) IsEventType(a_et EventTyper) bool {
	eventTypes := e.EventTypes()
	if et, ok := eventTypes[a_et.Id()]; ok {
		return et.IsEventType(a_et)
	} else {
		return false
	}
}

func (e Event) StartTime() time.Time {
	return e.eventDetails.StartTime
}

func (e Event) EndTime() time.Time {
	return e.eventDetails.EndTime
}

func (e Event) Title() string {
	return e.eventDetails.Title
}

func (e Event) Description() string {
	return e.eventDetails.Description
}

func (e Event) Location() string {
	return e.eventDetails.Location
}

func (e Event) EventTypes() map[int]EventTyper {
	return e.eventTypes
}
