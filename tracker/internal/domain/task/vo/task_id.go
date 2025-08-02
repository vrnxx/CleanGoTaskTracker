package vo

import "github.com/google/uuid"

type TaskID struct {
	Value uuid.UUID
}

func (t TaskID) AsGenericType() string {
	return t.Value.String()
}
