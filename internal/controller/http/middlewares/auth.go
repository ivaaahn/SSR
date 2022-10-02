package middlewares

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"ssr/config"
	"ssr/pkg/misc"
	"strings"
)

func CheckRole(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		split := strings.Split(c.Request().URL.Path, "/")

		// /api/<role>/... -> ["", "api", "vNo", "<role>" , ...]

		if len(split) < 4 {
			return echo.ErrForbidden
		}
		expRole := split[3]
		_, recRole := misc.ExtractCtx(c)

		mapping := map[string]string{"st": "students", "sv": "supervisors"}

		if mapping[recRole] != expRole {
			return echo.ErrForbidden
		}

		if err := next(c); err != nil {
			c.Error(err)
		}

		return nil
	}
}

func MakeAuthMiddleware(config *config.Config) echo.MiddlewareFunc {
	return middleware.JWTWithConfig(middleware.JWTConfig{
		Claims:     &misc.AppJWTClaims{},
		SigningKey: []byte(config.SigningKey),
		ContextKey: "ctx",
	})
}
