package srv

func routes() *Router {
	r := NewRouter()
	r.POST("/registrations", createRegistration)
	return r
}
