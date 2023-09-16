package error

import "fmt"

const (
	INVALID_REQUEST      = 1
	NO_DATA_FOUND        = 2
	UNAUTHORIZED_REQUEST = 3
	PROCESSING_ERROR     = 4
)

type ServiceError struct {
	Err  error
	Code int32
}

func (e *ServiceError) Error() string {
	return fmt.Sprintf("[Code:%v, Error:%v]", e.Code, e.Err.Error())
}
