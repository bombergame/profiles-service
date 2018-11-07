package errs

const (
	AccessDeniedErrorMessage = "access denied"
)

type AccessDeniedError struct {
	ServiceError
}

func NewAccessDeniedError() error {
	return NewAccessDeniedInnerError(nil)
}

func NewAccessDeniedInnerError(innerErr error) error {
	err := &AccessDeniedError{}
	err.message = AccessDeniedErrorMessage
	err.innerErr = innerErr
	return err
}
