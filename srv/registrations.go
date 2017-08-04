package srv

import (
	"net/http"

	"github.com/jmataya/nile"
)

func createRegistration(c nile.Context) nile.Response {
	msg := map[string]string{"message": "registration!"}
	resp := nile.NewGenericResponse(http.StatusOK, msg)
	return resp
}
