package errs

const (
	ConstraintViolationErrorMessagePrefix = "constraint violation: "
)

type ConstraintViolationError struct {
	ServiceError
}

func NewConstraintViolationError(message string) error {
	return NewConstraintViolationInnerError(message, nil)
}

func NewConstraintViolationInnerError(message string, innerErr error) error {
	err := &ConstraintViolationError{}
	err.message = ConstraintViolationErrorMessagePrefix + message
	err.innerErr = innerErr
	return err
}
