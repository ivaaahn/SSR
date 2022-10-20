// Package app configures and runs application.
package app

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"ssr/config"
	"ssr/internal/app/mw"
	ctrl "ssr/internal/controller/http"
	"ssr/internal/service"
	"ssr/internal/service/repo_pg"
	"ssr/pkg/logger"
	"ssr/pkg/postgres"
)

func setupMiddlewares(cfg *config.Config, server *echo.Echo) {
	server.Use(mw.ServerHeader())
	server.Use(middleware.CORS())

	if cfg.Performance == "false" {
		server.Use(middleware.Logger())
	}

	server.Use(middleware.Recover())
	server.Use(mw.HandleErrors())
}

func makeInjections(server *echo.Echo, pg *postgres.Postgres, l *logger.Logger, cfg *config.Config) {
	relationRepo := repo_pg.NewRelation(pg, l)
	workRepo := repo_pg.NewWork(pg, l)
	waypointRepo := repo_pg.NewWaypointRepo(pg, l)
	userRepo := repo_pg.NewUser(pg, l)
	studentRepo := repo_pg.NewStudent(pg, l)
	supervisorRepo := repo_pg.NewSupervisor(pg, l)

	authService := service.NewAuth(userRepo, l, cfg.Auth.TokenExp, []byte(cfg.Auth.SigningKey))
	profileService := service.NewProfile(studentRepo, supervisorRepo, l)
	workService := service.NewWork(workRepo, relationRepo, studentRepo, supervisorRepo, waypointRepo, l)
	relationService := service.NewRelation(relationRepo, l)

	ctrl.NewRouter(
		server,
		l,
		cfg,
		authService,
		profileService,
		workService,
		relationService,
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
	server.Use(mw.ServerHeader())
	setupMiddlewares(cfg, server)
	makeInjections(server, pg, loggerObject, cfg)

	if err := server.Start(cfg.HTTP.Port); err != http.ErrServerClosed {
		server.Logger.Fatal(err)
	}
}
