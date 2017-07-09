package operations

// Operation describes any action performed by the server or a client that
// either changes the document or communicates a message between the parties.
type Operation interface {
	ID() int
	Position() int
	Type() OperationType
	Apply(string) (string, error)
	Transform(Operation) Operation
	UpdatePosition(int) Operation
}

type operation struct {
	id       int
	opType   OperationType
	position int
}

func (op operation) ID() int {
	return op.id
}

func (op operation) Position() int {
	return op.position
}

func (op operation) Type() OperationType {
	return op.opType
}
