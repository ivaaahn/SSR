package http

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"ssr/config"
	"ssr/internal/controller/http/middlewares"
	"ssr/pkg/logger"
	"strconv"
)

type supervisor struct {
	l              logger.Interface
	profileService ProfileService
	workService    WorkService
}

// ShowAccount godoc
// @Summary      Get supervisor's profile
// @Tags         supervisor
// @Produce      json
// @Param        supervisor_id path int  true  "Supervisor ID"
// @Success      200  {object}  dto.Supervisor
// @Router       /api/v1/supervisors/{supervisor_id}/profile [get]
// @Security	 OAuth2Password
func (ctrl *supervisor) getProfile(ctx echo.Context) error {
	supervisorID, _ := strconv.Atoi(ctx.Param("supervisor_id"))

	respDTO, err := ctrl.profileService.GetSupervisorProfile(supervisorID)
	if err != nil {
		ctrl.l.Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "TODO")
	}

	return ctx.JSON(http.StatusOK, respDTO)
}

// ShowAccount godoc
// @Summary      Get supervisor's works
// @Tags         supervisor
// @Param        supervisor_id path int  true  "Supervisor ID"
// @Produce      json
// @Success      200  {object}  dto.SupervisorViewWorkPlenty
// @Router       /api/v1/supervisors/{supervisor_id}/works [get]
// @Security	 	OAuth2Password
func (ctrl *supervisor) getWorks(ctx echo.Context) error {
	supervisorID, _ := strconv.Atoi(ctx.Param("supervisor_id"))

	respDTO, err := ctrl.workService.GetSupervisorWorks(supervisorID)
	if err != nil {
		ctrl.l.Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "TODO")
	}

	return ctx.JSON(http.StatusOK, respDTO)
}

func NewSupervisorRoutes(
	router *echo.Group,
	l logger.Interface,
	config *config.Config,
	profileService ProfileService,
	workService WorkService,
) {
	ctrl := &supervisor{l, profileService, workService}

	g := router.Group("/supervisors", middlewares.MakeAuthMiddleware(config), middlewares.CheckRole)

	{
		g.GET("/:supervisor_id/profile", ctrl.getProfile)
		g.GET("/:supervisor_id/works", ctrl.getWorks)
	}

}
