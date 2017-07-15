package srv

import "github.com/jmataya/hermes/utils"

func routes() *utils.Router {
	r := utils.NewRouter()
	r.POST("/registrations", createRegistration)
	return r
}
