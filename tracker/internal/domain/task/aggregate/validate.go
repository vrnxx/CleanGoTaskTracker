package aggregate

import (
	"github.com/vrnxx/CleanGoTaskTracker/tracker/internal/domain/task/exceptions"
	"github.com/vrnxx/CleanGoTaskTracker/tracker/internal/domain/task/vo/task_status"
)

func (t *Task) checkNotDeleted() error {
	if t.Deleted {
		exc := exceptions.TaskIsDeleted{}.Exception(t.ID.AsGenericType())
		return &exc
	}
	return nil
}

func (t *Task) checkNotCompleted() error {
	if t.Status == task_status.Completed {
		exc := exceptions.TaskIsCompleted{}.Exception(t.ID.AsGenericType())
		return &exc
	}
	return nil
}
