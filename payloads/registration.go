package payloads

import "github.com/jmataya/hermes/errors"

// CreateRegistration is the payload needed to register a user in the system.
type CreateRegistration struct {
	Email     string  `json:"email"`
	Password  string  `json:"password"`
	FirstName *string `json:"first_name"`
	LastName  *string `json:"last_name"`
}

// Validate ensures that the payload is in a valid state.
func (c CreateRegistration) Validate() error {
	if c.Email == "" {
		return errors.NewFieldIsNil("email")
	} else if c.Password == "" {
		return errors.NewFieldIsNil("password")
	} else if c.FirstName != nil && *c.FirstName == "" {
		return errors.NewFieldIsNil("first_name")
	} else if c.LastName != nil && *c.LastName == "" {
		return errors.NewFieldIsNil("last_name")
	}

	return nil
}
