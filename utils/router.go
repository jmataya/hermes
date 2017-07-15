package utils

import "github.com/labstack/echo"

// Router is a wrapper for the Echo router.
type Router struct {
	e *echo.Echo
}

// NewRouter initializes a new Router with appropriate middleware.
func NewRouter() *Router {
	e := echo.New()
	e.Use(RegisterContext)
	return &Router{e}
}

// Start begins serving the application.
func (r *Router) Start(host string) error {
	return r.e.Start(host)
}

// GET registers a new GET route for a path with a matching handler.
func (r *Router) GET(path string, handler RequestHandler) {
	r.e.GET(path, mapContext(handler))
}

// POST registers a new POST route for a path with a matching handler.
func (r *Router) POST(path string, handler RequestHandler) {
	r.e.POST(path, mapContext(handler))
}

func mapContext(rh RequestHandler) echo.HandlerFunc {
	return func(ec echo.Context) error {
		rc := ec.(*RequestContext)
		return rh(rc)
	}
}
