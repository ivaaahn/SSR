package http

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"ssr/internal/misc"
	"ssr/internal/usecase"
	"ssr/pkg/logger"
)

type userRoutes struct {
	l  logger.Interface
	uc usecase.IUserUseCase
}

func (r *userRoutes) me(ctx echo.Context) error {
	email := misc.ExtractEmailFromContext(ctx)
	r.l.Debug(fmt.Sprintf("Email: %s", email))

	respDTO, err := r.uc.Me(email)
	if err != nil {
		r.l.Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "TODO")
	}

	return ctx.JSON(http.StatusOK, respDTO)
}

func NewUserRoutes(router *echo.Group, l logger.Interface, uc usecase.IUserUseCase) {
	r := &userRoutes{l, uc}

	g := router.Group("/user")

	{
		g.GET("/me", r.me)
	}

}
