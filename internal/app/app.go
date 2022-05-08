// Package app configures and runs application.
package app

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"ssr/config"
	"ssr/internal/controller/http"
	"ssr/pkg/logger"
	"ssr/pkg/postgres"
)

// Run creates objects via constructors.
func Run(cfg *config.Config) {
	l := logger.New(cfg.Log.Level)

	pg, err := postgres.New(cfg.Pg.DSN)

	if err != nil {
		l.Fatal(fmt.Errorf("app - Run - postgres.New: %w", err))
	}

	defer pg.Close()

	e := echo.New()

	e.Use(middleware.Logger())

	http.NewRouter(e, l)

	e.Logger.Fatal(e.Start(cfg.HTTP.Port))
}
