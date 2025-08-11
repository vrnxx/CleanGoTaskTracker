package task_delete

import "time"

type TaskDeleteInfo struct {
	Deleted   bool
	DeletedAt *time.Time
}

func NewTaskDeleteInfo() TaskDeleteInfo {
	return TaskDeleteInfo{
		Deleted:   false,
		DeletedAt: nil,
	}
}

func DeletedNow() TaskDeleteInfo {
	now := time.Now().UTC()
	return TaskDeleteInfo{
		Deleted:   true,
		DeletedAt: &now,
	}
}
