package aggregate

import "github.com/google/uuid"

type Event interface {
	Bytes() ([]byte, error)
	AggregateID() uuid.UUID
}

type AggregateRoot struct {
	events []Event
}

func (r *AggregateRoot) RecordEvent(event Event) {
	r.events = append(r.events, event)
}

func (r *AggregateRoot) Events() []Event {
	return r.events
}

func (r *AggregateRoot) UnloadEvents() []Event {
	unloaded := make([]Event, len(r.Events()))
	copy(unloaded, r.Events())
	return unloaded
}
