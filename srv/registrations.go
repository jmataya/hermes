package srv

import (
	"net/http"

	"github.com/jmataya/hermes/utils"
)

func createRegistration(c *utils.RequestContext) error {
	resp := map[string]string{
		"message": "okay",
	}

	return c.JSON(http.StatusOK, resp)
}
