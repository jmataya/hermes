package errors

import "fmt"

// DatabaseError represents errors occurred working with the database or
// data that may be stored in the database.
type DatabaseError struct {
	error
}

// NewDeleteFailed is an error to throw when deleting a row fails.
func NewDeleteFailed(modelName string, id int, rowsAffected int64) *DatabaseError {
	const msg = "Error: unable to delete %s with id %d, expected 1 affected row, got %d"
	err := fmt.Errorf(msg, modelName, id, rowsAffected)
	return &DatabaseError{err}
}

// NewModelHasID occurs when a model is expected to have no ID but one is found.
func NewModelHasID(modelName string, id int) *DatabaseError {
	const msg = "Error: to create, model %s must have ID=0, got %d"
	err := fmt.Errorf(msg, modelName, id)
	return &DatabaseError{err}
}

// NewRequiredFieldIsEmpty when a model's field must be set.
func NewRequiredFieldIsEmpty(modelName, fieldName string) *DatabaseError {
	const msg = "Error: field %s on %s must be non-empty"
	err := fmt.Errorf(msg, modelName, fieldName)
	return &DatabaseError{err}
}
