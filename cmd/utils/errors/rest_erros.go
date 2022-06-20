package errors

import (
	"net/http"
)

type RestErr struct {
	ErrMessage string `json:"message"`
	ErrStatus  int    `json:"status"`
	ErrError   string `json:"error"`
}

func NewBadRequestError(message string) *RestErr {
	return &RestErr{
		ErrMessage: message,
		ErrStatus:  http.StatusBadRequest,
		ErrError:   "bad_request",
	}
}

func NewRestError(message string, status int, err string) *RestErr {
	return &RestErr{
		ErrMessage: message,
		ErrStatus:  status,
		ErrError:   err,
	}
}

func NewNotFoundError(message string) *RestErr {
	return &RestErr{
		ErrMessage: message,
		ErrStatus:  http.StatusNotFound,
		ErrError:   "not_found",
	}
}
