package mw

import (
	"github.com/labstack/echo/v4"
)

func ServerHeader() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Response().Header().Set(echo.HeaderServer, "golang-echo-server")
			return next(c)
		}
	}
}
