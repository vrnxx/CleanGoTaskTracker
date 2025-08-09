package events

import (
	"encoding/json"
	"github.com/google/uuid"
	"time"
)

type BaseEvent struct {
	EventID        uuid.UUID
	EventType      string
	EventTimestamp time.Time
}

func NewBaseEvent(eventType string) BaseEvent {
	return BaseEvent{
		EventID:        uuid.New(),
		EventType:      eventType,
		EventTimestamp: time.Now().UTC(),
	}
}

func (e *BaseEvent) AggregateID() uuid.UUID {
	return e.EventID
}

func (e *BaseEvent) Bytes() ([]byte, error) {
	b, err := json.Marshal(e)
	if err != nil {
		return nil, err
	}
	return b, nil
}
