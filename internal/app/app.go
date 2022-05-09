// Package app configures and runs application.
package app

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"ssr/config"
	"ssr/internal/controller/http"
	"ssr/internal/repo"
	"ssr/internal/usecase"
	"ssr/pkg/logger"
	"ssr/pkg/postgres"
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

	authUC := usecase.NewAuthUC(repo.NewAuthPGRepo(pg, l), cfg.Auth.TokenExp, []byte(cfg.Auth.SigningKey))
	userUC := usecase.NewUserUC(repo.NewUserPGRepo(pg, l))

	http.NewRouter(e, l, authUC, userUC)

	e.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		Claims:     &jwt.StandardClaims{},
		SigningKey: []byte(cfg.SigningKey),
		ContextKey: "userEmail",
		Skipper: func(c echo.Context) bool {
			// Skip middleware if path is equal 'login'
			if c.Request().URL.Path == "/api/auth/login" {
				return true
			}
			return false
		},
	}))

	e.Logger.Fatal(e.Start(cfg.HTTP.Port))
}
