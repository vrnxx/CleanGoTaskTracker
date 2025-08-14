package events

import (
	"github.com/google/uuid"
	"github.com/vrnxx/CleanGoTaskTracker/tracker/internal/domain/common/aggregate"
	"github.com/vrnxx/CleanGoTaskTracker/tracker/internal/domain/common/events"
)

type TaskCreated struct {
	events.BaseEvent
	TaskID      uuid.UUID  `json:"taskID"`
	Title       string     `json:"title"`
	AssigneeID  *uuid.UUID `json:"assigneeID"`
	Status      string     `json:"status"`
	Priority    string     `json:"priority"`
	Description string     `json:"description"`
	Number      int        `json:"number"`
}

func NewTaskCreated(
	taskID uuid.UUID,
	title string,
	assigneeID *uuid.UUID,
	status,
	priority,
	description string,
	number int,
) aggregate.Event {
	return &TaskCreated{
		BaseEvent:   events.NewBaseEvent("TaskCreated"),
		TaskID:      taskID,
		Title:       title,
		AssigneeID:  assigneeID,
		Status:      status,
		Priority:    priority,
		Description: description,
		Number:      number,
	}
}

func (e *TaskCreated) Bytes() ([]byte, error) {
	return events.Bytes(e)
}

func (e *TaskCreated) AggregateID() uuid.UUID {
	return e.TaskID
}
