package errors

func CommonError(text string) error {
	return &commonError{text}
}

type commonError struct {
	s string
}

func (e *commonError) Error() string {
	return e.s
}
