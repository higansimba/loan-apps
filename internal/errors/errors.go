package errors

import (
	"fmt"
	"net/http"
)

// ErrorType represents the type of error
type ErrorType string

const (
	// ErrorTypeValidation represents validation errors
	ErrorTypeValidation ErrorType = "VALIDATION_ERROR"
	// ErrorTypeNotFound represents not found errors
	ErrorTypeNotFound ErrorType = "NOT_FOUND"
	// ErrorTypeUnauthorized represents unauthorized access errors
	ErrorTypeUnauthorized ErrorType = "UNAUTHORIZED"
	// ErrorTypeForbidden represents forbidden access errors
	ErrorTypeForbidden ErrorType = "FORBIDDEN"
	// ErrorTypeInternal represents internal server errors
	ErrorTypeInternal ErrorType = "INTERNAL_ERROR"
)

// AppError represents an application error
type AppError struct {
	Type    ErrorType `json:"type"`
	Message string    `json:"message"`
	Err     error     `json:"error,omitempty"`
}

// Error returns the error message
func (e *AppError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %v", e.Message, e.Err)
	}
	return e.Message
}

// StatusCode returns the HTTP status code for the error type
func (e *AppError) StatusCode() int {
	switch e.Type {
	case ErrorTypeValidation:
		return http.StatusBadRequest
	case ErrorTypeNotFound:
		return http.StatusNotFound
	case ErrorTypeUnauthorized:
		return http.StatusUnauthorized
	case ErrorTypeForbidden:
		return http.StatusForbidden
	default:
		return http.StatusInternalServerError
	}
}

// NewValidationError creates a new validation error
func NewValidationError(message string, err error) *AppError {
	return &AppError{
		Type:    ErrorTypeValidation,
		Message: message,
		Err:     err,
	}
}

// NewNotFoundError creates a new not found error
func NewNotFoundError(message string) *AppError {
	return &AppError{
		Type:    ErrorTypeNotFound,
		Message: message,
	}
}

// NewUnauthorizedError creates a new unauthorized error
func NewUnauthorizedError(message string) *AppError {
	return &AppError{
		Type:    ErrorTypeUnauthorized,
		Message: message,
	}
}

// NewForbiddenError creates a new forbidden error
func NewForbiddenError(message string) *AppError {
	return &AppError{
		Type:    ErrorTypeForbidden,
		Message: message,
	}
}

// NewInternalError creates a new internal error
func NewInternalError(message string, err error) *AppError {
	if err == nil {
		fmt.Println(err.Error())
	}
	return &AppError{
		Type:    ErrorTypeInternal,
		Message: message,
		Err:     err,
	}
}
