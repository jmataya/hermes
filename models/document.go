package models

import "github.com/jmataya/hermes/operations"

// Document represents the structure, operations, and server state for a file.
type Document struct {
	ServerMessageID int

	history []operations.Operation
	str     string
}

// NewDocument creates a new empty document.
func NewDocument() *Document {
	return &Document{
		ServerMessageID: 0,
		history:         []operations.Operation{},
		str:             "",
	}
}

// Apply updates the document based on an operation and records that
// operation in its history.
func (d *Document) Apply(op operations.Operation) error {
	// TODO: Make this work in a multi-threaded environment.
	newStr, err := op.Apply(d.str)
	if err != nil {
		return err
	}

	d.str = newStr
	d.history = append(d.history, op)
	return nil
}

// String returns the contents of the document as a string.
func (d *Document) String() string {
	return d.str
}
