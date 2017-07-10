package operations

// Noop is an operation that has no effect on the document contents.
// It is used if a transformation renders another operation unneeded or if the
// server is attempting to communicate server message position.
type Noop struct {
	operation
}

// NewNoop creates a new Noop operation.
func NewNoop() *Noop {
	op := operation{opType: NoopType, position: 0}
	return &Noop{op}
}

// Apply returns a modified copy of the current document string. Since this is
// a Noop, the original is returned.
func (n Noop) Apply(current string) (string, error) {
	return current, nil
}

// Transform modifies the supplied operation based on the current opration.
// Since this is a Noop, the original operation is returned.
func (n Noop) Transform(op Operation) Operation {
	return op
}

// UpdatePosition returns a clone of the original operation with a new position.
// Since this is a Noop, the original operation is returned.
func (n Noop) UpdatePosition(pos int) Operation {
	return n
}
