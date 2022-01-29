package rest

import (
	"github.com/AmirrezaKN/examinator/pkg/log"
	"github.com/AmirrezaKN/examinator/service"
	"github.com/labstack/echo/v4"
)

type RestController struct {
	log     log.Logger
	server  *echo.Echo
	service service.ExaminatorService
}

func NewRestController(logger log.Logger, server *echo.Echo, service service.ExaminatorService) RestController {
	return RestController{
		log:     logger,
		server:  server,
		service: service,
	}
}

func (c RestController) Serve() error {
	return c.server.Start(":6789")
}
