package middlewares

import (
	"github.com/labstack/echo/v4"
	"ssr/pkg/misc"
	"strings"
)

func CheckRole(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		split := strings.Split(c.Request().URL.Path, "/")

		// /api/<role>/... -> ["", "api", "<role>" , ...]
		expRole := split[2]
		_, recRole := misc.ExtractInfoFromContext(c)

		println(recRole, expRole)

		if recRole != expRole {
			return echo.ErrForbidden
		}

		if err := next(c); err != nil {
			c.Error(err)
		}

		return nil
	}
}
