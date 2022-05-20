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

// ShowAccount godoc
// @Summary      Login into account
// @Tags         auth
// @Accept       x-www-form-urlencoded
// @Produce      json
// @Param        username formData  string  true  "User email"
// @Param        password formData  string  true  "User password"
// @Success      200  {object}  dto.LoginResponse
// @Failure      401
// @Failure      500
// @Router       /api/auth/login [post]
func (r *authRoutes) login(ctx echo.Context) error {
	email := ctx.FormValue("username")
	password := ctx.FormValue("password")

	r.l.Debug(fmt.Sprintf("Email %s; Password: %s", email, password))

	respDTO, err := r.uc.Login(email, password)

	if err != nil {
		r.l.Error(err)
		return echo.ErrUnauthorized
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
