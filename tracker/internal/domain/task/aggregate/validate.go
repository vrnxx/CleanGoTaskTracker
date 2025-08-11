package aggregate

import "github.com/vrnxx/CleanGoTaskTracker/tracker/internal/domain/task/exceptions"

func (t *Task) checkNotDeleted() error {
	if t.Deleted {
		exc := exceptions.TaskIsDeleted{}.Exception(t.ID.AsGenericType())
		return &exc
	}
	return nil
}
