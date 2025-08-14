package exceptions

import (
	"fmt"
	"github.com/vrnxx/CleanGoTaskTracker/tracker/internal/domain/common/exceptions"
)

type TaskIsDeleted struct {
	exceptions.DomainException
}

func (t TaskIsDeleted) Exception(context string) TaskIsDeleted {
	return TaskIsDeleted{
		DomainException: exceptions.DomainException{
			Message: "Task already deleted",
			Ctx:     fmt.Sprintf("task id: %s", context),
		},
	}
}

type TaskIsCompleted struct {
	exceptions.DomainException
}

func (t TaskIsCompleted) Exception(context string) TaskIsCompleted {
	return TaskIsCompleted{
		DomainException: exceptions.DomainException{
			Message: "Task is completed",
			Ctx:     fmt.Sprintf("task id: %s", context),
		},
	}
}
