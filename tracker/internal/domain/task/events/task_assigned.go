package events

import (
	"github.com/google/uuid"
	"github.com/vrnxx/CleanGoTaskTracker/tracker/internal/domain/common/aggregate"
	"github.com/vrnxx/CleanGoTaskTracker/tracker/internal/domain/common/events"
)

type TaskAssigned struct {
	events.BaseEvent
	TaskID     uuid.UUID `json:"taskID"`
	AssigneeID uuid.UUID `json:"assigneeID"`
}

func NewTaskAssigned(taskID, assigneeID uuid.UUID) aggregate.Event {
	return &TaskAssigned{
		BaseEvent:  events.NewBaseEvent("TaskAssigned"),
		TaskID:     taskID,
		AssigneeID: assigneeID,
	}
}

func (e *TaskAssigned) Bytes() ([]byte, error) {
	return events.Bytes(e)
}

func (e *TaskAssigned) AggregateID() uuid.UUID {
	return e.TaskID
}
