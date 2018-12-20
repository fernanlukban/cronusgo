package event

type EventType struct {
    Id int
    Name string
}

type EventTyper interface {
    IsEventType(a_et EventTyper) bool
}

func (et EventType) IsEventType(a_et EventTyper) bool {
    return et.Id == a_et.Id
}