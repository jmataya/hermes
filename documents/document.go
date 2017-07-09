package documents

import "github.com/jmataya/ot-editor/operations"

// Document represents the structure, operations, and server state for a file.
type Document struct {
	Operations      []operations.Operation
	ServerMessageID int

	str string
}

// NewDocument creates a new empty document.
func NewDocument() *Document {
	return &Document{
		Operations:      []operations.Operation{},
		ServerMessageID: 0,
		str:             "",
	}
}
