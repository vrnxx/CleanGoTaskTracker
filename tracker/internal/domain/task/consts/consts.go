package consts

type TaskStatus string

const (
	MinTitleLength              = 1
	MaxTitleLength              = 128
	StatusOpen       TaskStatus = "open"
	StatusInProgress TaskStatus = "in_progress"
	StatusClosed     TaskStatus = "closed"
)
