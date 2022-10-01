// Package app configures and runs application.
package app

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	"net/http"
	"ssr/config"
	ctrl "ssr/internal/controller/http"
	"ssr/internal/service"
	"ssr/internal/service/repo_pg"
	"ssr/pkg/logger"
	"ssr/pkg/misc"
	"ssr/pkg/postgres"
	_ "ssr/swagger"
	"strings"
)

func setupMiddlewares(server *echo.Echo, cfg *config.Config) {
	server.Use(middleware.CORS())
	server.Use(middleware.Logger())
	server.Use(middleware.Recover())
	server.Use(middleware.JWTWithConfig(middleware.JWTConfig{
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

func makeInjections(server *echo.Echo, pg *postgres.Postgres, l *logger.Logger, cfg *config.Config) {
	relationRepo := repo_pg.NewRelation(pg, l)
	authRepo := repo_pg.NewAuth(pg, l)
	profileRepo := repo_pg.NewProfile(pg, l)
	workRepo := repo_pg.NewWork(pg, l)
	feedbackRepo := repo_pg.NewFeedback(pg, l)

	authService := service.NewAuth(authRepo, l, cfg.Auth.TokenExp, []byte(cfg.Auth.SigningKey))
	profileService := service.NewProfile(profileRepo, l)
	bidService := service.NewBid(relationRepo, l)
	workService := service.NewWork(workRepo, relationRepo, l)
	relationService := service.NewRelation(relationRepo, l)
	feedbackService := service.NewFeedback(feedbackRepo, l)

	ctrl.NewRouter(
		server,
		l,
		authService,
		profileService,
		profileService,
		bidService,
		bidService,
		workService,
		workService,
		relationService,
		feedbackService,
	)
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
	makeInjections(server, pg, loggerObject, cfg)

	server.GET("/swagger*", echoSwagger.WrapHandler)
	if err := server.Start(cfg.HTTP.Port); err != http.ErrServerClosed {
		server.Logger.Fatal(err)
	}
}
