package rest

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (c RestController) Index(e echo.Context) error {
	err := e.String(http.StatusAccepted, "Hello World")
	if err != nil {
		return err
	}
	return nil
}
