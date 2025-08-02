package vo

type Priority string

const (
	PriorityCritical Priority = "critical"
	PriorityHigh     Priority = "high"
	PriorityMedium   Priority = "medium"
	PriorityLow      Priority = "low"
	PriorityMinor    Priority = "minor"
)

var prioritiesMapping = map[Priority]int{
	PriorityCritical: 99,
	PriorityHigh:     10,
	PriorityMedium:   5,
	PriorityLow:      2,
	PriorityMinor:    0,
}

func (p Priority) IsValid() bool {
	_, ok := prioritiesMapping[p]
	return ok
}

func (p Priority) IsHigherThan(other Priority) bool {
	if p.IsValid() && other.IsValid() {
		return prioritiesMapping[other] < prioritiesMapping[p]
	}
	return false
}
