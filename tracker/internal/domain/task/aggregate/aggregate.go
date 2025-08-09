package aggregate

import (
	"github.com/google/uuid"
	"github.com/vrnxx/CleanGoTaskTracker/tracker/internal/domain/task/vo"
	"github.com/vrnxx/CleanGoTaskTracker/tracker/internal/domain/task/vo/priority"
)

type Task struct {
	ID          vo.TaskID
	Title       vo.Title
	AuthorID    uuid.UUID
	AssigneeID  uuid.UUID
	Status      vo.TaskStatus
	Priority    priority.Priority
	Description vo.Description
}

func NewTask(
	id vo.TaskID,
	title vo.Title,
	authorID uuid.UUID,
	priority priority.Priority,
	description vo.Description,
) Task {
	return Task{
		ID:          id,
		Title:       title,
		AuthorID:    authorID,
		Status:      vo.StatusOpen,
		Priority:    priority,
		Description: description,
	}
}
