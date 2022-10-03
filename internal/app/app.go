// Package app configures and runs application.
package app

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"ssr/config"
	ctrl "ssr/internal/controller/http"
	"ssr/internal/service"
	"ssr/internal/service/repo_pg"
	"ssr/pkg/logger"
	"ssr/pkg/postgres"
)

func setupMiddlewares(server *echo.Echo) {
	server.Use(middleware.CORS())
	server.Use(middleware.Logger())
	server.Use(middleware.Recover())
}

func makeInjections(server *echo.Echo, pg *postgres.Postgres, l *logger.Logger, cfg *config.Config) {
	relationRepo := repo_pg.NewRelation(pg, l)
	//profileRepo := repo_pg.NewProfile(pg, l)
	workRepo := repo_pg.NewWork(pg, l)
	waypointRepo := repo_pg.NewWaypointRepo(pg, l)
	feedbackRepo := repo_pg.NewFeedback(pg, l)
	userRepo := repo_pg.NewUser(pg, l)
	studentRepo := repo_pg.NewStudent(pg, l)
	supervisorRepo := repo_pg.NewSupervisor(pg, l)

	authService := service.NewAuth(userRepo, l, cfg.Auth.TokenExp, []byte(cfg.Auth.SigningKey))
	profileService := service.NewProfile(studentRepo, supervisorRepo, l)
	//bidService := service.NewBid(relationRepo, l)
	workService := service.NewWork(workRepo, relationRepo, studentRepo, supervisorRepo, waypointRepo, l)
	relationService := service.NewRelation(relationRepo, l)
	feedbackService := service.NewFeedback(feedbackRepo, l)

	ctrl.NewRouter(
		server,
		l,
		cfg,
		authService,
		profileService,
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
	setupMiddlewares(server)
	makeInjections(server, pg, loggerObject, cfg)

	if err := server.Start(cfg.HTTP.Port); err != http.ErrServerClosed {
		server.Logger.Fatal(err)
	}
}
