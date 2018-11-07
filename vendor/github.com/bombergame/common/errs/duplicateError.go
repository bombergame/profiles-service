package errs

type DuplicateError struct {
	ServiceError
}

func NewDuplicateError(message string) error {
	return NewDuplicateInnerError(message, nil)
}

func NewDuplicateInnerError(message string, innerErr error) error {
	err := &DuplicateError{}
	err.message = message
	err.innerErr = innerErr
	return err
}
