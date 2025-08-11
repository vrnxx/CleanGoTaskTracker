package vo

import (
	"github.com/vrnxx/CleanGoTaskTracker/tracker/internal/domain/task/consts"
	"github.com/vrnxx/CleanGoTaskTracker/tracker/internal/domain/task/exceptions"
)

type Description struct {
	Value string
}

func NewDescription(text string) (Description, error) {
	if consts.MaxDescriptionLength < len(text) {
		return Description{}, exceptions.InvalidDescriptionLength{}.Exception(text)
	}
	return Description{Value: text}, nil
}
