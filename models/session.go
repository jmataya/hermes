package models

import "github.com/jmataya/hermes/errors"

// Session represents a connection between a server and client.
type Session struct {
	doc             *Document
	clientMessageID int
	outgoing        []*Message
}

// NewSession initializes a new Session for an existing document.
func NewSession(doc *Document) (*Session, error) {
	if doc == nil {
		return nil, errors.NewFieldIsNil("doc")
	}

	return &Session{
		doc:             doc,
		clientMessageID: 0,
		outgoing:        []*Message{},
	}, nil
}

// Receive handles a new message from the client and updates the document.
func (s *Session) Receive(msg *Message) error {
	if err := s.validateMessage(msg); err != nil {
		return err
	}

	outgoing := []*Message{}
	xformOp := msg.Operation

	for _, srvMsg := range s.outgoing {
		if srvMsg.ServerMessageID > msg.ServerMessageID {
			outgoing = append(outgoing, srvMsg)
			xformOp = srvMsg.Operation.Transform(xformOp)
		}
	}

	if err := s.doc.Apply(xformOp); err != nil {
		return err
	}

	s.outgoing = outgoing
	s.clientMessageID++

	return nil
}

func (s *Session) validateMessage(msg *Message) error {
	expectedClientID := s.clientMessageID + 1
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
