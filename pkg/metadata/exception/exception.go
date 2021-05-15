package exception

type NotImplementedError struct {
	Err error
}

func (e *NotImplementedError) Error() string {
	return "not implemented"
}
