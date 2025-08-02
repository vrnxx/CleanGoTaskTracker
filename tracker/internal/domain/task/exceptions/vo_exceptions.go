package exceptions

import (
	"fmt"
	"github.com/vrnxx/CleanGoTaskTracker/tracker/internal/domain/common/exceptions"
	"github.com/vrnxx/CleanGoTaskTracker/tracker/internal/domain/task/consts"
)

type InvalidTitleLengthCreation struct {
	exceptions.DomainException
}

func (e InvalidTitleLengthCreation) Exception(context string) *InvalidTitleLengthCreation {
	return &InvalidTitleLengthCreation{
		DomainException: exceptions.DomainException{
			Message: fmt.Sprintf(
				"Title length must be gt %d and lt %d",
				consts.MinTitleLength,
				consts.MaxTitleLength,
			),
			Ctx: context,
		},
	}
}
