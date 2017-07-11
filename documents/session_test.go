package documents

import (
	"testing"

	"github.com/jmataya/hermes/operations"
)

func TestReceiveMessage(t *testing.T) {
	type action struct {
		opType   operations.OperationType
		position int
		value    string
		want     string
	}

	var tests = []struct {
		actions []action
	}{
		{[]action{
			action{operations.InsertType, 0, "hello", "hello"},
		}},
		{[]action{
			action{operations.InsertType, 0, "a", "a"},
			action{operations.InsertType, 1, "b", "ab"},
			action{operations.InsertType, 2, "c", "abc"},
			action{operations.InsertType, 3, "d", "abcd"},
			action{operations.InsertType, 4, "e", "abcde"},
		}},
		{[]action{
			action{operations.InsertType, 0, "hold", "hold"},
			action{operations.InsertType, 1, "ello, w", "hello, wold"},
			action{operations.DeleteType, 8, "", "hello, wld"},
			action{operations.InsertType, 8, "e", "hello, weld"},
		}},
		{[]action{
			action{operations.InsertType, 0, "hold", "hold"},
			action{operations.DeleteType, 3, "", "hol"},
			action{operations.DeleteType, 2, "", "ho"},
			action{operations.DeleteType, 1, "", "h"},
			action{operations.DeleteType, 0, "", ""},
			action{operations.InsertType, 0, "hold", "hold"},
		}},
	}

	for _, test := range tests {
		doc := NewDocument()
		session, _ := NewSession(doc)

		for idx, action := range test.actions {
			var op operations.Operation

			switch action.opType {
			case operations.InsertType:
				op = operations.NewInsert(action.position, action.value)
			case operations.DeleteType:
				op = operations.NewDelete(action.position)
			}

			msg := &Message{
				Operation:       op,
				ClientMessageID: idx + 1,
				ServerMessageID: 0,
			}

			if err := session.Receive(msg); err != nil {
				t.Error(err.Error())
				return
			}

			if doc.String() != action.want {
				t.Errorf(
					"Error receiving %s: wanted %s, got %s",
					action.opType,
					action.want,
					doc.String())
			}
		}
	}
}
