package errors

import (
	"net/http"
)

// UniError stsruct
type UniError struct {
	Code    int16
	Message string
}

// BadRequest indicates request not valid
func (e *UniError) BadRequest(key string) *UniError {
	return &UniError{
		Code : http.StatusBadRequest,
		Message : "Error Request Parameters " + key,
	}
}
