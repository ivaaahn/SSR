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
	stProfile StProfileService,
	svProfile SvProfileService,
	stBids StBidService,
	stWorks StWorkService,
	svWorks SvWorkService,
	stRelations StRelationService,
	feedback FeedbackService,
) {
	g := server.Group("/api/v1")
	g.GET("/swagger/*", echoSwagger.WrapHandler)

	{
		NewAuthRoutes(g, l, auth)
		NewStudentRoutes(g, l, config, stProfile, stBids, stWorks, stRelations, feedback)
		NewSupervisorRoutes(g, l, config, svProfile, svWorks)
	}
}
