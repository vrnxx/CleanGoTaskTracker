package exceptions

import (
	"fmt"

	"github.com/vrnxx/CleanGoTaskTracker/tracker/internal/domain/common/exceptions"
	"github.com/vrnxx/CleanGoTaskTracker/tracker/internal/domain/task/consts"
)

type InvalidTitleLength struct {
	exceptions.DomainException
}

func (e InvalidTitleLength) Exception(context string) *InvalidTitleLength {
	return &InvalidTitleLength{
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

type InvalidDescriptionLength struct {
	exceptions.DomainException
}

func (e InvalidDescriptionLength) Exception(context string) *InvalidDescriptionLength {
	return &InvalidDescriptionLength{
		DomainException: exceptions.DomainException{
			Message: fmt.Sprintf(
				"Description length must be lt %d",
				consts.MaxDescriptionLength,
			),
			Ctx: context,
		},
	}
}
