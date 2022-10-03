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

//// ShowAccount godoc
//// @Summary      Get supervisors of the work
//// @Tags         works
//// @Param        work_id path int true "Work ID"
//// @Produce      json
//// @Success      200  {object}  dto.WorkPlenty
//// @Router       /api/v1/works/ [get]
//// @Security	 OAuth2Password
//func (ctrl *works) getPlenty(ctx echo.Context) error {
//	respDTO, err := ctrl.workService.GetPlenty()
//	if err != nil {
//		return echo.ErrNotFound
//	}
//
//	return ctx.JSON(http.StatusOK, respDTO)
//}

// ShowAccount godoc
// @Summary      Get supervisors of the work
// @Tags         works
// @Param        work_id path int true "Work ID"
// @Produce      json
// @Success      200  {object}  dto.WorkFullResp
// @Router       /api/v1/works/{work_id} [get]
// @Security	 OAuth2Password
func (ctrl *works) get(ctx echo.Context) error {
	workID, _ := strconv.Atoi(ctx.Param("work_id"))

	respDTO, err := ctrl.workService.Get(workID)
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
		//works.GET("/", ctrl.getPlenty)
		works.GET("/:work_id", ctrl.get)
	}

}
