package vo

import "time"

type TaskInfo struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	Number    int
}

func NewTaskInfo(num int) TaskInfo {
	return TaskInfo{
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Number:    num,
	}
}
