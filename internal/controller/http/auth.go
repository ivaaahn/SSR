package http

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"ssr/internal/usecase"
	"ssr/pkg/logger"
)

type authRoutes struct {
	l  logger.Interface
	uc usecase.IAuthUC
}

func (r *authRoutes) login(ctx echo.Context) error {
	email := ctx.FormValue("username")
	password := ctx.FormValue("password")

	r.l.Debug(fmt.Sprintf("Email %s; Password: %s", email, password))

	respDTO, err := r.uc.Login(email, password)

	if err != nil {
		r.l.Error(err)
	}

	return ctx.JSON(http.StatusOK, respDTO)
}

func NewAuthRoutes(router *echo.Group, l logger.Interface, uc usecase.IAuthUC) {
	ar := &authRoutes{l, uc}

	g := router.Group("/auth")

	{
		g.POST("/login", ar.login)
	}

}
