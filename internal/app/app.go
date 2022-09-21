// Package app configures and runs application.
package app

import (
	"fmt"
	"github.com/labstack/echo/v4"
	m "github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	"ssr/config"
	"ssr/internal/controller/http"
	uc "ssr/internal/usecase"
	repo "ssr/internal/usecase/repo_pg"
	"ssr/pkg/logger"
	"ssr/pkg/misc"
	"ssr/pkg/postgres"
	_ "ssr/swagger"
	"strings"
)

func setupMiddlewares(server *echo.Echo, cfg *config.Config) {
	server.Use(m.CORS())
	server.Use(m.Logger())
	server.Use(m.Recover())
	server.Use(m.JWTWithConfig(m.JWTConfig{
		Claims:     &misc.AppJWTClaims{},
		SigningKey: []byte(cfg.SigningKey),
		ContextKey: "ctx",
		Skipper: func(c echo.Context) bool {
			// Skip middleware if 'login' or 'swagger'
			path := c.Request().URL.Path
			split := strings.Split(path, "/")
			return split[2] == "auth" || split[1] == "swagger"
		},
	}))
}

func setupUC(server *echo.Echo, pg *postgres.Postgres, l *logger.Logger, cfg *config.Config) {
	authUC := uc.NewAuth(repo.NewAuthPgRepo(pg, l), l, cfg.Auth.TokenExp, []byte(cfg.Auth.SigningKey))
	profileUC := uc.NewProfile(repo.NewProfilePgRepo(pg, l), l)
	bidUC := uc.NewBid(repo.NewSSRPgRepo(pg, l), l)
	workUC := uc.NewWork(repo.NewWorkPgRepo(pg, l), repo.NewSSRPgRepo(pg, l), l)
	ssrUC := uc.NewSSR(repo.NewSSRPgRepo(pg, l), l)
	feedBackUC := uc.NewFeedback(repo.NewFeedback(pg, l), l)

	http.NewRouter(server, l, authUC, profileUC, bidUC, bidUC, workUC, workUC, ssrUC, feedBackUC)
}

func Run(cfg *config.Config) {
	loggerObject := logger.New(cfg.Log.Level)

	pg, err := postgres.New(cfg.Pg.DSN)
	if err != nil {
		loggerObject.Fatal(fmt.Errorf("app - Run - postgres.NewAuth: %w", err))
	}

	defer pg.Close()

	server := echo.New()
	setupMiddlewares(server, cfg)
	setupUC(server, pg, loggerObject, cfg)

	server.GET("/swagger*", echoSwagger.WrapHandler)
	server.Logger.Fatal(server.Start(cfg.HTTP.Port))
}
