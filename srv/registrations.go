package srv

import "net/http"

func createRegistration(c *RequestContext) error {
	resp := map[string]string{
		"message": "okay",
	}

	return c.JSON(http.StatusOK, resp)
}
