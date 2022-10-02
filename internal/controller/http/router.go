package http

import (
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	"ssr/config"
	"ssr/pkg/logger"
	_ "ssr/swagger"
)

func NewRouter(
	server *echo.Echo,
	l logger.Interface,
	config *config.Config,
	auth AuthService,
	profiles ProfileService,
	stBids StBidService,
	works WorkService,
	stRelations StRelationService,
	feedback FeedbackService,
) {
	g := server.Group("/api/v1")
	g.GET("/swagger/*", echoSwagger.WrapHandler)

	{
		NewAuthRoutes(g, l, auth)
		NewStudentRoutes(g, l, config, profiles, stBids, works, stRelations, feedback)
		NewSupervisorRoutes(g, l, config, profiles, works)
		NewWorksRoutes(g, l, config, works)
	}
}
