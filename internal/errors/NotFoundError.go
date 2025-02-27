package errors

import "net/http"

type NotFoundError struct {
	Message string
	Code    int
}

func (e *NotFoundError) New() *BaseError {
	return &BaseError{
		Message: "Not Found",
		Code:    http.StatusNotFound,
	}
}
