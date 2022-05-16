package app

import (
	"net/http"
)

type Error struct {
	StatusCode int         `json:"statusCode"`
	RootErr    error       `json:"-"`
	Message    string      `json:"message"`
	Log        string      `json:"log"`
	Key        string      `json:"key"`
	Data       interface{} `json:"data"`
}

func NewErrorResponse(message string, statusCode int, data interface{}, err error) *Error {

	return &Error{
		Message:    message,
		StatusCode: statusCode,
		Data:       data,
		RootErr:    err,
	}
}

func (e *Error) RootError() error {
	if err, ok := e.RootErr.(*Error); ok {
		return err.RootError()
	}

	return e.RootErr
}

func (e *Error) Error() string {
	return e.RootError().Error()
}

func ErrInvalidRequest(err error) *Error {
	return NewErrorResponse(err.Error(), http.StatusBadRequest, nil, err)
}

func ErrInvalidRequestWithMessage(err error, message string) *Error {
	return NewErrorResponse(message, http.StatusBadRequest, nil, err)
}

func ErrInternalServer(err error) *Error {
	return NewErrorResponse(err.Error(), http.StatusInternalServerError, nil, err)
}

func ErrEntityExisted(err error, message string) *Error {
	return NewErrorResponse(message, http.StatusConflict, nil, err)
}

func ErrNoPermission(err error) *Error {
	return NewErrorResponse("Don't have permission", http.StatusForbidden, nil, err)
}

func ErrEntityNotFound(err error, message string) *Error {
	return NewErrorResponse(message, http.StatusNotFound, nil, err)
}

func ErrConflict(err error, message string) *Error {
	return NewErrorResponse(message, http.StatusConflict, nil, err)
}
