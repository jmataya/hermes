package errors

import "fmt"

// InvalidMessage represents error states when there is a problem with the
// message sent from a client to the server.
type InvalidMessage struct {
	error
}

// NewInvalidClientMessageID is an error that's displayed when a client message
// ID is invalid.
func NewInvalidClientMessageID(expectedID, actualID int) *InvalidMessage {
	const msg = "Invalid message: expected client_message_id %d got %d"
	err := fmt.Errorf(msg, expectedID, actualID)
	return &InvalidMessage{err}
}

// NewInvalidServerMessageID occurs when the message's server ID is not the
// latest server message ID or within the range of outgoing message IDs.
func NewInvalidServerMessageID(expectedStartID, expectedEndID, actualID int) *InvalidMessage {
	var err error

	if expectedStartID == expectedEndID {
		const msg = "Invalid message: expected server_message_id %d got %d"
		err = fmt.Errorf(msg, expectedStartID, actualID)
	} else {
		const msg = "Invalid message: expected server_message_id between %d and %d, got %d"
		err = fmt.Errorf(msg, expectedStartID, expectedEndID, actualID)
	}

	return &InvalidMessage{err}
}
