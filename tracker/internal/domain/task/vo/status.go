package vo

type TaskStatus string

const (
	StatusOpen       TaskStatus = "open"
	StatusInProgress TaskStatus = "in_progress"
	StatusClosed     TaskStatus = "closed"
)

var statusesMapping = map[TaskStatus]struct{}{
	StatusOpen:       {},
	StatusInProgress: {},
	StatusClosed:     {},
}

func (s TaskStatus) IsValid() bool {
	_, ok := statusesMapping[s]
	return ok
}
