package operations

import "github.com/jmataya/ot-editor/errors"

// Insert is an operation that represents inserting a string into a document.
type Insert struct {
	operation
	value string
}

// Apply returns a modified copy of the current document string with the insert
// operation applied.
func (i Insert) Apply(current string) (string, error) {
	if i.position > len(current) {
		return current, errors.NewNonExistentPosition(len(current), i.position)
	}

	head := current[0:i.position]
	tail := current[i.position:]
	return head + i.value + tail, nil
}

// Transform modifies the supplied operation if this insertion affects the
// result of the operation. If not, the original operation is returned.
// A key assumption is that the insert before the supplied operation.
func (i Insert) Transform(op Operation) Operation {
	if op.Type() == NoopType {
		return op
	}

	if i.Position() > op.Position() {
		return op
	}

	newPos := op.Position() + len(i.value)
	return op.UpdatePosition(newPos)
}

// UpdatePosition returns a clone of the current operation with a new position.
func (i Insert) UpdatePosition(pos int) Operation {
	op := operation{id: i.id, opType: i.opType, position: pos}
	return &Insert{op, i.value}
}
