package payloads

// Payload is the structure that all JSON payloads are expected to conform.
type Payload interface {
	Validate() error
}
