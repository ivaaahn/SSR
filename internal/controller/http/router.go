package http

import (
	"github.com/labstack/echo/v4"
	"ssr/pkg/logger"
)

func NewRouter(echo *echo.Echo, l logger.Interface) {
	g := echo.Group("/api")

	newHelloRoutes(g, l)

}
