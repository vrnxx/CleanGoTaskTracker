package vo

import (
	"github.com/vrnxx/CleanGoTaskTracker/tracker/internal/domain/task/consts"
	"github.com/vrnxx/CleanGoTaskTracker/tracker/internal/domain/task/exceptions"
)

type Title struct {
	Value string
}

func NewTitle(val string) (Title, error) {
	if len(val) < consts.MinTitleLength || consts.MaxTitleLength < len(val) {
		exc := exceptions.InvalidTitleLength{}.Exception(val)
		return Title{}, &exc
	}
	return Title{Value: val}, nil
}

func (vo Title) AsGenericType() string {
	return vo.Value
}
