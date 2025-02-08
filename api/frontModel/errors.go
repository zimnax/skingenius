package frontModel

import "fmt"

var ProcessingServerError = NewSkingeniusError(500, "processing server error")

type SkingeniusError struct {
	Code    int
	Message string
}

func NewSkingeniusError(code int, message string) *SkingeniusError {
	return &SkingeniusError{
		Code:    code,
		Message: message,
	}
}

func (e *SkingeniusError) Error() string {
	return fmt.Sprintf("error [%d],  %s", e.Code, e.Message)
}
