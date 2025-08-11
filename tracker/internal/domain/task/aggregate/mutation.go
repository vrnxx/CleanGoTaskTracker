package aggregate

import (
	delInfo "github.com/vrnxx/CleanGoTaskTracker/tracker/internal/domain/task/vo/task_delete"
)

func (t *Task) DeleteTask() error {
	if err := t.checkNotDeleted(); err != nil {
		return err
	}
	t.TaskDeleteInfo = delInfo.DeletedNow()
	return nil
}
