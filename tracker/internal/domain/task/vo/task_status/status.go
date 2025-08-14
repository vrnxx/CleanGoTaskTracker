package task_status

type TaskStatus string

const (
	Open       TaskStatus = "open"
	InProgress TaskStatus = "in_progress"
	Completed  TaskStatus = "completed"
)

var statusesMapping = map[TaskStatus]struct{}{
	Open:       {},
	InProgress: {},
	Completed:  {},
}

func (s TaskStatus) IsValid() bool {
	_, ok := statusesMapping[s]
	return ok
}
