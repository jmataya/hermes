package errors

import (
	"errors"
	"fmt"
)

// InvalidOperation represents error states when there is a problem when
// interacting with or applying an operation.
type InvalidOperation struct {
	error
}

// NewNoValue is an error that occurs when trying to access the value of an
// operation that shouldn't have a value.
func NewNoValue() *InvalidOperation {
	err := errors.New("Operation does not contain a value")
	return &InvalidOperation{err}
}

// NewNonExistentPosition is an error that occurs when attempting to modify a
// document at a position that doesn't exist.
func NewNonExistentPosition(expectedMax, actual int) *InvalidOperation {
	const msg = "Invalid operation: expected position between 0 and %d, got %d"
	err := fmt.Errorf(msg, expectedMax, actual)
	return &InvalidOperation{err}
}
