package errs

type NotFoundError struct {
	ServiceError
}

func NewNotFoundError(message string) error {
	return NewNotFoundInnerError(message, nil)
}

func NewNotFoundInnerError(message string, innerErr error) error {
	err := &NotFoundError{}
	err.message = message
	err.innerErr = innerErr
	return err
}
