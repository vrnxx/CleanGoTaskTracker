package events

import (
	"github.com/google/uuid"
	"github.com/vrnxx/CleanGoTaskTracker/tracker/internal/domain/common/aggregate"
	"github.com/vrnxx/CleanGoTaskTracker/tracker/internal/domain/common/events"
)

type TaskDeleted struct {
	events.BaseEvent
	TaskID uuid.UUID `json:"taskID"`
}

func NewTaskDeleted(taskID uuid.UUID) aggregate.Event {
	return &TaskDeleted{
		BaseEvent: events.NewBaseEvent("TaskDeleted"),
		TaskID:    taskID,
	}
}

func (e *TaskDeleted) Bytes() ([]byte, error) {
	return events.Bytes(e)
}

func (e *TaskDeleted) AggregateID() uuid.UUID {
	return e.TaskID
}
