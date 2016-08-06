package markdone

type errorString struct {
	message string
}

func (e *errorString) Error() string {
	return e.message
}
