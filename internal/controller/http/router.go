package http

import (
	"github.com/labstack/echo/v4"
	"ssr/internal/usecase"
	"ssr/pkg/logger"
)

func NewRouter(
	echo *echo.Echo,
	l logger.Interface,
	authUC usecase.IUsecaseAuth,
	profileUC usecase.IUsecaseProfile,
	stBidUC usecase.IUsecaseStudentBid,
	svBidUC usecase.IUseCaseSupervisorBid,
	stWorkUC usecase.IStudentWorkUC,
	svWorkUC usecase.ISupervisorWorkUC,
	stSsrUC usecase.IUseCaseStudentRelation,
) {
	g := echo.Group("/api")

	{
		NewAuthRoutes(g, l, authUC)
		NewStudentRoutes(g, l, profileUC, stBidUC, stWorkUC, stSsrUC)
		NewSupervisorRoutes(g, l, profileUC, svBidUC, svWorkUC)
	}
}
