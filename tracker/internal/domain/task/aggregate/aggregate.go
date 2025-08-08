package aggregate

import (
	"github.com/google/uuid"
	"github.com/vrnxx/CleanGoTaskTracker/tracker/internal/domain/task/vo"
	"github.com/vrnxx/CleanGoTaskTracker/tracker/internal/domain/task/vo/priority"
)

type Task struct {
	vo.TaskID
	vo.Title
	AuthorID   uuid.UUID
	AssigneeID uuid.UUID
	vo.TaskStatus
	priority.Priority
	vo.Description
}

func NewTask(
	id vo.TaskID,
	title vo.Title,
	authorID uuid.UUID,
	priority priority.Priority,
	description vo.Description,
) Task {
	return Task{
		TaskID:      id,
		Title:       title,
		AuthorID:    authorID,
		TaskStatus:  vo.StatusOpen,
		Priority:    priority,
		Description: description,
	}
}
