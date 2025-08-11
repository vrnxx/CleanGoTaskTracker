package task_priority

type Priority string

const (
	Critical Priority = "critical"
	High     Priority = "high"
	Medium   Priority = "medium"
	Low      Priority = "low"
	Minor    Priority = "minor"
)

var prioritiesMapping = map[Priority]int{
	Critical: 99,
	High:     10,
	Medium:   5,
	Low:      2,
	Minor:    0,
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
