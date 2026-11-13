package errors

import (
	"errors"
	"fmt"
)

// ErrNotFound is a sentinel error.
var ErrNotFound = errors.New("item not found")

// ErrValidation is another sentinel error.
var ErrValidation = errors.New("validation failed")

// ValidationError adds a field and reason.
type ValidationError struct {
	Field  string
	Reason string
}

func (e *ValidationError) Error() string {
	// TODO: return "validation failed: field: reason"
	return ""
}

func (e *ValidationError) Unwrap() error {
	// TODO: return ErrValidation
	return nil
}

// FindUser simulates a lookup.
func FindUser(id int) (string, error) {
	// TODO: return ErrNotFound when id <= 0
	return "", nil
}

// CreateUser validates input and wraps validation errors with context.
func CreateUser(name, email string) (string, error) {
	// TODO:
	// - empty name => ValidationError{Name: "name", Reason: "required"}
	// - email missing "@" => ValidationError{Field: "email", Reason: "invalid"}
	// - otherwise return "created: " + name
	return "", nil
}

// Handle demonstrates errors.Is and errors.As.
func Handle(err error) string {
	// TODO:
	// - if errors.Is(err, ErrNotFound), return "not found"
	// - if errors.As(err, &ValidationError{}), return "invalid input"
	// - otherwise return "unknown error"
	return ""
}

// WrapError wraps an error with useful context.
func WrapError(err error, context string) error {
	// TODO: use fmt.Errorf with %w
	return nil
}

// MustNotBeEmpty returns ErrValidation for empty input.
func MustNotBeEmpty(value, field string) error {
	// TODO
	return nil
}

var _ = fmt.Sprintf
