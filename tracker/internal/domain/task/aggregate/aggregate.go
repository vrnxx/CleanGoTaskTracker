package aggregate

import (
	"github.com/google/uuid"
	"github.com/vrnxx/CleanGoTaskTracker/tracker/internal/domain/task/vo"
	"github.com/vrnxx/CleanGoTaskTracker/tracker/internal/domain/task/vo/priority"
	"github.com/vrnxx/CleanGoTaskTracker/tracker/internal/domain/task/vo/status"
)

type Task struct {
	ID          vo.TaskID
	Title       vo.Title
	AuthorID    uuid.UUID
	AssigneeID  uuid.UUID
	Status      status.TaskStatus
	Priority    priority.Priority
	Description vo.Description
	vo.TaskInfo
}

func NewTask(
	id vo.TaskID,
	title vo.Title,
	authorID uuid.UUID,
	priority priority.Priority,
	description vo.Description,
	number int,
) Task {
	return Task{
		ID:          id,
		Title:       title,
		AuthorID:    authorID,
		Status:      status.Open,
		Priority:    priority,
		Description: description,
		TaskInfo:    vo.NewTaskInfo(number),
	}
}
