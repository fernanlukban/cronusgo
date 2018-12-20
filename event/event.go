package event

import (
    "time"
)

type Event struct {
    startTime time.Time
    endTime time.Time
    title string
    description string
    location string
    eventTypes map[int]EventTyper
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
    return e.startTime
}

func (e Event) EndTime() time.Time {
    return e.endTime
}

func (e Event) Title() string {
    return e.title
}

func (e Event) Description() string {
    return e.description
}

func (e Event) Location() string {
    return e.location
}

func (e Event) EventTypes() map[int]EventTyper {
    return e.eventTypes
}