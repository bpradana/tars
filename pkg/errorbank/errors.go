package errorbank

import (
	"fmt"
)

// MessageError represents errors that occur during message operations
type MessageError struct {
	Operation string
	Message   string
	Cause     error
}

// Error returns the formatted error message
func (e *MessageError) Error() string {
	if e.Cause != nil {
		return fmt.Sprintf("[%s] %s: %v", e.Operation, e.Message, e.Cause)
	}
	return fmt.Sprintf("[%s] %s", e.Operation, e.Message)
}

// Unwrap returns the underlying error
func (e *MessageError) Unwrap() error {
	return e.Cause
}

// TemplateError represents errors that occur during template variable substitution
type TemplateError struct {
	Variable string
	Message  string
	Cause    error
}

// Error returns the formatted error message
func (e *TemplateError) Error() string {
	if e.Cause != nil {
		return fmt.Sprintf("[Template] variable '%s': %s: %v", e.Variable, e.Message, e.Cause)
	}
	return fmt.Sprintf("[Template] variable '%s': %s", e.Variable, e.Message)
}

// Unwrap returns the underlying error
func (e *TemplateError) Unwrap() error {
	return e.Cause
}

// ValidationError represents errors that occur during message validation
type ValidationError struct {
	Field   string
	Message string
	Value   any
}

// Error returns the formatted error message
func (e *ValidationError) Error() string {
	return fmt.Sprintf("[Validation] field '%s': %s (value: %v)", e.Field, e.Message, e.Value)
}

// Common error constructors

// NewMessageError creates a new MessageError
func NewMessageError(operation, message string, cause error) *MessageError {
	return &MessageError{
		Operation: operation,
		Message:   message,
		Cause:     cause,
	}
}

// NewTemplateError creates a new TemplateError
func NewTemplateError(variable, message string, cause error) *TemplateError {
	return &TemplateError{
		Variable: variable,
		Message:  message,
		Cause:    cause,
	}
}

// NewValidationError creates a new ValidationError
func NewValidationError(field, message string, value any) *ValidationError {
	return &ValidationError{
		Field:   field,
		Message: message,
		Value:   value,
	}
}

// IsMessageError checks if an error is a MessageError
func IsMessageError(err error) bool {
	_, ok := err.(*MessageError)
	return ok
}

// IsTemplateError checks if an error is a TemplateError
func IsTemplateError(err error) bool {
	_, ok := err.(*TemplateError)
	return ok
}

// IsValidationError checks if an error is a ValidationError
func IsValidationError(err error) bool {
	_, ok := err.(*ValidationError)
	return ok
}
