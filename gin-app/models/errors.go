package models

import "net/http"

type ApiError struct {
	Status  int
	Message string
}

func (e *ApiError) Error() string {
	return e.Message
}

func NewBadRequest(msg string) *ApiError {
	return &ApiError{Status: http.StatusBadRequest, Message: msg}
}

func NewNotFound(msg string) *ApiError {
	return &ApiError{Status: http.StatusNotFound, Message: msg}
}

func NewUnauthorized(msg string) *ApiError {
	return &ApiError{Status: http.StatusUnauthorized, Message: msg}
}

func NewConflict(msg string) *ApiError {
	return &ApiError{Status: http.StatusConflict, Message: msg}
}

func NewInternal(msg string) *ApiError {
	return &ApiError{Status: http.StatusInternalServerError, Message: msg}
}
