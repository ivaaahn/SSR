package http

import (
	"github.com/labstack/echo/v4"
	"ssr/internal/usecase"
	"ssr/pkg/logger"
)

func NewRouter(
	echo *echo.Echo,
	l logger.Interface,
	authUC usecase.IAuthUC,
	profileUC usecase.IProfileUC,
	stBidUC usecase.IStudentBidUC,
	svBidUC usecase.ISupervisorBidUC,
) {
	g := echo.Group("/api")

	{
		NewAuthRoutes(g, l, authUC)
		NewStudentRoutes(g, l, profileUC, stBidUC)
		NewSupervisorRoutes(g, l, profileUC, svBidUC)
	}
}
