package aggregate

import (
	"github.com/google/uuid"
	common "github.com/vrnxx/CleanGoTaskTracker/tracker/internal/domain/common/aggregate"
	"github.com/vrnxx/CleanGoTaskTracker/tracker/internal/domain/task/vo"
	"github.com/vrnxx/CleanGoTaskTracker/tracker/internal/domain/task/vo/task_delete"
	"github.com/vrnxx/CleanGoTaskTracker/tracker/internal/domain/task/vo/task_info"
	"github.com/vrnxx/CleanGoTaskTracker/tracker/internal/domain/task/vo/task_priority"
	"github.com/vrnxx/CleanGoTaskTracker/tracker/internal/domain/task/vo/task_status"
)

type Task struct {
	common.AggregateRoot
	ID          vo.TaskID
	Title       vo.Title
	AuthorID    uuid.UUID
	AssigneeID  *uuid.UUID
	Status      task_status.TaskStatus
	Priority    task_priority.Priority
	Description vo.Description
	task_info.TaskInfo
	task_delete.TaskDeleteInfo
}

func NewTask(
	id vo.TaskID,
	title vo.Title,
	authorID uuid.UUID,
	assigneeID *uuid.UUID,
	priority task_priority.Priority,
	description vo.Description,
	number int,
) Task {
	return Task{
		ID:             id,
		Title:          title,
		AuthorID:       authorID,
		AssigneeID:     assigneeID,
		Status:         task_status.Open,
		Priority:       priority,
		Description:    description,
		TaskInfo:       task_info.NewTaskInfo(number),
		TaskDeleteInfo: task_delete.NewTaskDeleteInfo(),
	}
}
