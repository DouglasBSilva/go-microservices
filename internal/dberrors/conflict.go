package dberrors

import "net/http"

type ConflictError struct{}

func (e *ConflictError) Error() string {
	return "attempted to create a record with duplicate key"
}

func (e *ConflictError) Code() int {
	return http.StatusConflict
}
