package errs

const (
	ServiceErrorMessage = "internal service error"
)

type ServiceError struct {
	message  string
	innerErr error
}

func NewServiceError(err error) error {
	return &ServiceError{
		message:  ServiceErrorMessage,
		innerErr: err,
	}
}

func (err *ServiceError) Error() string {
	return err.message
}

func (err *ServiceError) InnerError() string {
	if err.innerErr == nil {
		return ""
	} else {
		return err.innerErr.Error()
	}
}
