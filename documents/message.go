package documents

import "github.com/jmataya/ot-editor/operations"

// Message is the structure of communication between a client and the server.
// It describes both the operation that occurred, as well as the indexes for
// both the client and server, which is necessary to understand if and how the
// operation should be transformed.
type Message struct {
	Operation       operations.Operation `json:"operation"`
	ClientMessageID int                  `json:"client_message_id"`
	ServerMessageID int                  `json:"server_message_id"`
}

// Transform determines if a change is needed to an incoming message. If one is
// needed, a new transformed message is created. Otherwise the original message
// is returned
func (m Message) Transform(msg *Message) *Message {

	return msg
}
