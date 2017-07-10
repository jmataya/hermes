package errors

import (
	"fmt"
)

// GeneralError represents simple, application-wide errors.
type GeneralError struct {
	error
}

// NewFieldIsNil creates an error when a field with an expected value is nil.
func NewFieldIsNil(fieldName string) *GeneralError {
	const msg = "Error: field %s is nil"
	err := fmt.Errorf(msg, fieldName)
	return &GeneralError{err}
}
