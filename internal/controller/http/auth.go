package http

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"ssr/pkg/logger"
)

type auth struct {
	l       logger.Interface
	service AuthService
}

// ShowAccount godoc
// @Summary      Login into account
// @Tags         auth
// @Accept       x-www-form-urlencoded
// @Produce      json
// @Param        username formData  string  true  "UserFull email"
// @Param        password formData  string  true  "UserFull password"
// @Success      200  {object}  dto.LoginResponse
// @Failure      401
// @Failure      500
// @Router       /api/v1/auth/login [post]
func (ctrl *auth) login(ctx echo.Context) error {
	email := ctx.FormValue("username")
	password := ctx.FormValue("password")

	ctrl.l.Debug(fmt.Sprintf("Email %s; Password: %s", email, password))

	respDTO, err := ctrl.service.Login(email, password)
	if err != nil {
		return echo.ErrUnauthorized
	}

	return ctx.JSON(http.StatusOK, respDTO)
}

func NewAuthRoutes(router *echo.Group, l logger.Interface, service AuthService) {
	ctrl := &auth{l, service}

	g := router.Group("/auth")
	{
		g.POST("/login", ctrl.login)
	}

}
