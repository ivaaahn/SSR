package http

import (
	"github.com/labstack/echo/v4"
	"ssr/internal/usecase"
	"ssr/pkg/logger"
)

func NewRouter(
	echo *echo.Echo,
	l logger.Interface,
	auth usecase.IUsecaseAuth,
	profile usecase.IUsecaseProfile,
	studentBids usecase.IUsecaseStudentBid,
	supervisorBids usecase.IUseCaseSupervisorBid,
	studentWorks usecase.IStudentWorkUC,
	supervisorWorks usecase.ISupervisorWorkUC,
	studentRelations usecase.IUseCaseStudentRelation,
	feedback usecase.IUsecaseFeedback,
) {
	g := echo.Group("/api")

	{
		NewAuthRoutes(g, l, auth)
		NewStudentRoutes(g, l, profile, studentBids, studentWorks, studentRelations, feedback)
		NewSupervisorRoutes(g, l, profile, supervisorBids, supervisorWorks)
	}
}
