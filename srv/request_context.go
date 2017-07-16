package srv

import "github.com/labstack/echo"

// RequestContext is a wrapper around echo's context that provides error
// handling and simplified access to payloads and responses.
// The lifetime of a context should only for a single request, since the
// internal error makes the object stateful.
type RequestContext struct {
	echo.Context
	err error
}

// RequestHandler is the function signature needed to use the RequestContext.
type RequestHandler func(*RequestContext) error

// RegisterContext is a middleware handler for echo that registers the
// usage of the RequestContext. It should be called before any other
// middleware if it is to be used.
func RegisterContext(h echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		rc := &RequestContext{c, nil}
		return h(rc)
	}
}
