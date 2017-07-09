package documents

import (
	"fmt"

	"github.com/jmataya/ot-editor/errors"
)

// Session represents a connection between a server and client.
type Session struct {
	doc             Document
	clientMessageID int
	outgoing        []*Message
}

// Receive handles a new message from the client and updates the document.
func (s *Session) Receive(msg *Message) error {
	if err := s.validateMessage(msg); err != nil {
		return err
	}

	outgoing := []*Message{}

	for _, srvMsg := range s.outgoing {
		if srvMsg.ServerMessageID > msg.ServerMessageID {
			outgoing = append(outgoing, srvMsg)

			xformClient := srvMsg.Operation.Transform(msg.Operation)
			fmt.Printf("%+v\n", xformClient)
		}
	}

	s.outgoing = outgoing

	return nil
}

func (s *Session) validateMessage(msg *Message) error {
	expectedClientID := msg.ClientMessageID + 1
	if msg.ClientMessageID != expectedClientID {
		return errors.NewInvalidClientMessageID(expectedClientID, msg.ClientMessageID)
	}

	serverStartID, serverEndID := s.serverMessageRange()
	if msg.ServerMessageID < serverStartID || msg.ServerMessageID > serverEndID {
		return errors.NewInvalidServerMessageID(serverStartID, serverEndID, msg.ServerMessageID)
	}

	return nil
}

func (s *Session) serverMessageRange() (rangeStart int, rangeEnd int) {
	outgoingLen := len(s.outgoing)

	if outgoingLen == 0 {
		return s.doc.ServerMessageID, s.doc.ServerMessageID
	}

	return s.outgoing[0].ServerMessageID, s.outgoing[outgoingLen-1].ServerMessageID
}
