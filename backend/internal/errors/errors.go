package errors

import "errors"

var (
	ErrNotFound       = errors.New("not found")
	ErrUnauthorized   = errors.New("unauthorized")
	ErrInternalServer = errors.New("internal server error")
	// Add more common error types as needed
)

// NewValidationError creates a validation error with a specific message.
func NewValidationError(msg string) error {
	return errors.New("validation error: " + msg)
}
