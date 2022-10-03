package http

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"ssr/config"
	"ssr/internal/controller/http/middlewares"
	"ssr/pkg/logger"
	"strconv"
)

type works struct {
	l           logger.Interface
	workService WorkService
}

// ShowAccount godoc
// @Summary      Get supervisors of the work
// @Tags         works
// @Param        work_id path int true "Work ID"
// @Produce      json
// @Success      200  {object}  dto.WorkSupervisorPlenty
// @Router       /api/v1/works/{work_id}/supervisors [get]
// @Security	 OAuth2Password
func (ctrl *works) getSupervisors(ctx echo.Context) error {
	workID, _ := strconv.Atoi(ctx.Param("work_id"))

	respDTO, err := ctrl.workService.GetWorkSupervisors(workID)
	if err != nil {
		return echo.ErrNotFound
	}

	return ctx.JSON(http.StatusOK, respDTO)
}

func NewWorksRoutes(
	router *echo.Group,
	l logger.Interface,
	config *config.Config,
	worksService WorkService,
) {
	ctrl := &works{
		l,
		worksService,
	}

	works := router.Group("/works", middlewares.MakeAuthMiddleware(config))

	{
		works.GET("/:work_id/supervisors", ctrl.getSupervisors)
	}

}
