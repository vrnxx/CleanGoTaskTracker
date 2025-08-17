package services

import (
	"github.com/google/uuid"
	"github.com/vrnxx/CleanGoTaskTracker/tracker/internal/domain/task/aggregate"
	"github.com/vrnxx/CleanGoTaskTracker/tracker/internal/domain/task/events"
	"github.com/vrnxx/CleanGoTaskTracker/tracker/internal/domain/task/vo"
	"github.com/vrnxx/CleanGoTaskTracker/tracker/internal/domain/task/vo/task_priority"
)

type Service struct{}

func (Service) CreateTask(
	taskID vo.TaskID,
	title vo.Title,
	authorID uuid.UUID,
	assigneeID *uuid.UUID,
	priority task_priority.Priority,
	description vo.Description,
	number int,
) (aggregate.Task, error) {
	createdTask := aggregate.NewTask(
		taskID,
		title,
		authorID,
		assigneeID,
		priority,
		description,
		number,
	)
	createdTask.RecordEvent(
		events.NewTaskCreated(
			createdTask.ID.Value,
			createdTask.Title.Value,
			createdTask.AssigneeID,
			string(createdTask.Status),
			string(createdTask.Priority),
			createdTask.Description.Value,
			createdTask.Number,
		),
	)

	return createdTask, nil
}
