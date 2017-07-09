package operations

// OperationType indicates the type of a given operation.
type OperationType string

const (
	// Insert represents an insert operation.
	InsertType OperationType = "insert"

	// Delete represents a delete operation.
	DeleteType OperationType = "delete"

	// Noop represents a noop operation.
	NoopType OperationType = "noop"
)
