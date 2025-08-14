package events

import (
	"github.com/google/uuid"
	"github.com/vrnxx/CleanGoTaskTracker/tracker/internal/domain/common/events"
)

type TaskDeleted struct {
	events.BaseEvent
	TaskID uuid.UUID
}

func NewTaskDeleted(taskID uuid.UUID) *TaskDeleted {
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
