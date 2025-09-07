package apperror

import (
	"fmt"
)

// AppError provides structured error information for the application
type AppError struct {
	Type    error  // The error type
	Message string // Descriptive message
	Cause   error  // Underlying error, if any
}

// Error implements the error interface for AppErr.
func (err *AppError) Error() string {
	if err.Cause != nil {
		return fmt.Sprintf("%s: %s", err.Message, err.Cause.Error())
	}

	return err.Message
}

// New creates a new AppErr with the specified type, message, and optional cause.
func New(errorType error, message string, cause error) *AppError {
	// Prevent double-wrapping of the same error type.
	if err, ok := cause.(*AppError); ok && err.Type == errorType {
		return err
	}

	return &AppError{
		Type:    errorType,
		Message: message,
		Cause:   cause,
	}
}
