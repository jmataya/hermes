package payloads

import "github.com/jmataya/hermes/errors"
import "github.com/jmataya/nile"

const regPayloadCode = "70001"

// CreateRegistration is the payload needed to register a user in the system.
type CreateRegistration struct {
	Email     string  `json:"email"`
	Password  string  `json:"password"`
	FirstName *string `json:"first_name"`
	LastName  *string `json:"last_name"`
}

// Validate ensures that the payload is in a valid state.
func (c CreateRegistration) Validate() *nile.ErrorResponse {
	if c.Email == "" {
		return c.nilFieldError("email")
	} else if c.Password == "" {
		return c.nilFieldError("password")
	} else if c.FirstName != nil && *c.FirstName == "" {
		return c.nilFieldError("first_name")
	} else if c.LastName != nil && *c.LastName == "" {
		return c.nilFieldError("last_name")
	}

	return nil
}

func (c CreateRegistration) nilFieldError(field string) *nile.ErrorResponse {
	return nile.NewBadRequest(regPayloadCode, errors.NewFieldIsNil(field))
}
