package vo

import (
	"github.com/vrnxx/CleanGoTaskTracker/tracker/internal/domain/task/consts"
	"github.com/vrnxx/CleanGoTaskTracker/tracker/internal/domain/task/exceptions"
)

type Title struct {
	Value string
}

func (Title) Create(val string) (Title, error) {
	if len(val) < consts.MinTitleLength || consts.MaxTitleLength < len(val) {
		return Title{}, exceptions.InvalidTitleLengthCreation{}.Exception(val)
	}
	return Title{Value: val}, nil
}

func (vo Title) AsGenericType() string {
	return vo.Value
}
