package event

import (
    "time"
)

type Event struct {
    StartTime time.Time
    EndTime time.Time
    Title string
    Description string
    Location string
    EventTypes map[int]EventTyper
}

type Eventer interface {
    TotalDuration() time.Time
    IsEventType(a_et EventTyper) bool
}

func (e Event) TotalDuration() time.Time {
    var time_difference time.Duration = e.EndTime.Sub(e.StartTime)

    // Translates time_difference into a time.Time object
    total_duration := time.Time{}.add(time_difference)

    return total_duration
}

func (e Event) IsEventType(a_et EventTyper) bool {
    if et, ok := e.EventTypes[a_et.Id]; ok {
        return et.isEventType(a_et)
    } else {
        return false
    }
}