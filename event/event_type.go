package event

type EventType struct {
    id int
    name string
}

type EventTyper interface {
    IsEventType(a_et EventTyper) bool
    Id() int
    Name() string
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