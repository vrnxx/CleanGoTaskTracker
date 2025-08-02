package exceptions

type DomainException struct {
	Message string
	Ctx     string
}

func (e *DomainException) Error() string {
	return e.Message + " " + e.Ctx
}
