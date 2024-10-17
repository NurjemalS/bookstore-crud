package errors

import "fmt"

type CustomError struct {
	Message    string
	StatusCode int
}

func (e *CustomError) Error() string {
	return fmt.Sprintf("%s (status %d)", e.Message, e.StatusCode)
}

func NewCustomError(message string, statusCode int) *CustomError {
	return &CustomError{
		Message:    message,
		StatusCode: statusCode,
	}
}
