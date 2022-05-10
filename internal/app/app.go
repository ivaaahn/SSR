// Package app configures and runs application.
package app

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	"ssr/config"
	_ "ssr/docs/swagger"
	"ssr/internal/controller/http"
	"ssr/internal/usecase"
	"ssr/internal/usecase/repo_pg"
	"ssr/pkg/logger"
	"ssr/pkg/misc"
	"ssr/pkg/postgres"
	"strings"
)

// Run creates objects via constructors.
func Run(cfg *config.Config) {
	l := logger.New(cfg.Log.Level)

	pg, err := postgres.New(cfg.Pg.DSN)

	if err != nil {
		l.Fatal(fmt.Errorf("app - Run - postgres.NewAuthUC: %w", err))
	}

	defer pg.Close()

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	authUC := usecase.NewAuthUC(repo_pg.NewAuthPGRepo(pg, l), cfg.Auth.TokenExp, []byte(cfg.Auth.SigningKey))
	profileUC := usecase.NewProfileUC(repo_pg.NewProfilePGRepo(pg, l))
	bidUC := usecase.NewBidUC(repo_pg.NewSSRPGRepo(pg, l))

	http.NewRouter(e, l, authUC, profileUC, bidUC, bidUC)

	e.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		Claims:     &misc.JWTClaimsSSR{},
		SigningKey: []byte(cfg.SigningKey),
		ContextKey: "userEmail",
		Skipper: func(c echo.Context) bool {
			// Skip middleware if 'login' or 'swagger'
			path := c.Request().URL.Path

			split := strings.Split(path, "/")
			if split[2] == "auth" || split[1] == "swagger" {
				return true
			}
			return false
		},
	}))

	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.Logger.Fatal(e.Start(cfg.HTTP.Port))
}
