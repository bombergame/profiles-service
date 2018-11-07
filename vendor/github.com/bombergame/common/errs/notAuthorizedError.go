package errs

const (
	NotAuthorizedErrorMessage = "not authorized"
)

type NotAuthorizedError struct {
	ServiceError
}

func NewNotAuthorizedError() error {
	return NewNotAuthorizedInnerError(nil)
}

func NewNotAuthorizedInnerError(innerErr error) error {
	err := &NotAuthorizedError{}
	err.message = NotAuthorizedErrorMessage
	err.innerErr = innerErr
	return err
}
