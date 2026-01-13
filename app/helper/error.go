package helper

import "net/http"

type AppError struct {
	Code    int
	Message string
}

func (e *AppError) Error() string {
	return e.Message
}

func NewNotFound(msg string) error {
	return &AppError{Code: http.StatusNotFound, Message: msg}
}

func NewBadRequest(msg string) error {
	return &AppError{Code: http.StatusBadRequest, Message: msg}
}

func NewInternalServerError(msg string) error {
	return &AppError{Code: http.StatusInternalServerError, Message: msg}
}

func NewUnauthorizedError(msg string) error {
	return &AppError{Code: http.StatusUnauthorized, Message: msg}
}
