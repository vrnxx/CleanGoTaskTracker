package task_status

type TaskStatus string

const (
	Open       TaskStatus = "open"
	InProgress TaskStatus = "in_progress"
	Closed     TaskStatus = "closed"
)

var statusesMapping = map[TaskStatus]struct{}{
	Open:       {},
	InProgress: {},
	Closed:     {},
}

func (s TaskStatus) IsValid() bool {
	_, ok := statusesMapping[s]
	return ok
}
