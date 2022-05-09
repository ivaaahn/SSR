package http

import (
	"github.com/labstack/echo/v4"
	"ssr/internal/usecase"
	"ssr/pkg/logger"
)

func NewRouter(echo *echo.Echo, l logger.Interface, authUC usecase.IAuthUseCase, userUC usecase.IUserUseCase) {
	g := echo.Group("/api")

	{
		NewAuthRoutes(g, l, authUC)
		NewUserRoutes(g, l, userUC)
	}
}
