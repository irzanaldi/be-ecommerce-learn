package response

import "net/http"

type AppError struct {
	Code    int    `json:"-"`
	Message string `json:"message"`
}

func (e *AppError) Error() string {
	return e.Message
}

func ErrBadRequest(msg string) *AppError {
	return &AppError{Code: http.StatusBadRequest, Message: msg}
}

func ErrUnauthorized(msg string) *AppError {
	return &AppError{Code: http.StatusUnauthorized, Message: msg}
}

func ErrForbidden(msg string) *AppError {
	return &AppError{Code: http.StatusForbidden, Message: msg}
}

func ErrNotFound(msg string) *AppError {
	return &AppError{Code: http.StatusNotFound, Message: msg}
}

func ErrConflict(msg string) *AppError {
	return &AppError{Code: http.StatusConflict, Message: msg}
}

func ErrInternal(msg string) *AppError {
	return &AppError{Code: http.StatusInternalServerError, Message: msg}
}
