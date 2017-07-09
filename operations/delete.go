package operations

import "github.com/jmataya/ot-editor/errors"

// Delete is an operation that removes a single character from the document.
type Delete struct {
	operation
}

// Apply returns a modified copy of the current document string with the delete
// operation applied.
func (d Delete) Apply(current string) (string, error) {
	if d.position > len(current) {
		return current, errors.NewNonExistentPosition(len(current), d.position)
	}

	head := current[0 : d.position-1]
	tail := current[d.position:]
	return head + tail, nil
}

// Transform modifies the supplied operation if this deletion affects the result
// of the operation. If not, the original operation is returned. A key
// assumption i s that the delete occurs before the supplied operation.
func (d Delete) Transform(op Operation) Operation {
	if op.Type() == NoopType {
		return op
	}

	if op.Position() == 0 {
		return op
	}

	if d.Position() > op.Position() {
		return op
	}

	if d.Position() == op.Position() && op.Type() == DeleteType {
		return NewNoop(op.ID())
	}

	newPos := op.Position() - 1
	return op.UpdatePosition(newPos)
}

// UpdatePosition returns a clone of the current operation with a new position.
func (d Delete) UpdatePosition(pos int) Operation {
	op := operation{id: d.id, opType: d.opType, position: pos}
	return &Delete{op}
}
