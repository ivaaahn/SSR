package http

import (
	"github.com/labstack/echo/v4"
	"ssr/pkg/logger"
)

func NewRouter(
	echo *echo.Echo,
	l logger.Interface,
	auth AuthService,
	stProfile StProfileService,
	svProfile SvProfileService,
	stBids StBidService,
	stWorks StWorkService,
	stRelations StRelationService,
	feedback FeedbackService,
) {
	g := echo.Group("/api")

	{
		NewAuthRoutes(g, l, auth)
		NewStudentRoutes(g, l, stProfile, stBids, stWorks, stRelations, feedback)
		NewSupervisorRoutes(g, l, svProfile)
	}
}
