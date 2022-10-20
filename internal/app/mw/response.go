package mw

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func ServerHeader() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Response().Header().Set(echo.HeaderServer, "golang-echo-server")
			return next(c)
		}
	}
}

func HandleErrors() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			err := next(c)
			if err == middleware.ErrJWTMissing {
				return echo.NewHTTPError(http.StatusUnauthorized, "Token is not provided")
			}
			return err
		}
	}
}
