package aggregate

import (
	"github.com/google/uuid"
	"github.com/vrnxx/CleanGoTaskTracker/tracker/internal/domain/task/events"
	delInfo "github.com/vrnxx/CleanGoTaskTracker/tracker/internal/domain/task/vo/task_delete"
)

func (t *Task) AssignTo(assigneeID uuid.UUID) error {
	if err := t.checkNotCompleted(); err != nil {
		return err
	}
	// record event: task assigned to
	t.RecordEvent(
		events.NewTaskAssigned(t.ID.Value, assigneeID),
	)
	t.AssigneeID = assigneeID
	return nil
}

func (t *Task) DeleteTask() error {
	if err := t.checkNotDeleted(); err != nil {
		return err
	}
	// record event: task delete
	t.RecordEvent(
		events.NewTaskDeleted(t.ID.Value),
	)
	t.TaskDeleteInfo = delInfo.DeletedNow()
	return nil
}
