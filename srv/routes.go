package srv

import "github.com/jmataya/nile"

func routes() nile.Router {
	r := nile.New()
	r.POST("/registrations", createRegistration)
	return r
}
