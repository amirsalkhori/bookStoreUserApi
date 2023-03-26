package errors

import "net/http"

type RestError struct {
	Message string `json:"message"`
	Status  int32  `json:"status"`
	Error   string `json:"error"`
}

func NewBadRequestError(message string) *RestError {
	return &RestError{
		Message: message,
		Status:  http.StatusBadRequest,
		Error:   "Bad request",
	}
}
