package http

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"ssr/pkg/logger"
)

type helloRoutes struct {
	l logger.Interface
}

func (r *helloRoutes) helloHandler(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "Hello, World!")
}

func (r *helloRoutes) ep1(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "EP1")
}

func (r *helloRoutes) ep2(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "EP2")
}

func newHelloRoutes(router *echo.Group, l logger.Interface) {
	r := &helloRoutes{l}

	g := router.Group("/hello")

	{
		g.GET("/ep1", r.ep1)
		g.GET("/ep2", r.ep2)
		g.GET("/epta", r.helloHandler)
	}

}
