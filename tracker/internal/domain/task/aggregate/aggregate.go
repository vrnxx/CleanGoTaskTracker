package aggregate

import (
	"github.com/google/uuid"
	"github.com/vrnxx/CleanGoTaskTracker/tracker/internal/domain/task/vo"
)

type Task struct {
	ID         vo.TaskID
	Title      vo.Title
	AuthorID   uuid.UUID
	AssigneeID uuid.UUID
	Status     vo.TaskStatus
	Priority   vo.Priority
}
