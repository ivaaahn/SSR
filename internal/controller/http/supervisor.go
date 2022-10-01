package http

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"ssr/internal/controller/http/middlewares"
	"ssr/internal/dto"
	"ssr/pkg/logger"
	"ssr/pkg/misc"
	"strconv"
)

type supervisor struct {
	l              logger.Interface
	profileService SvProfileService
	bidService     SvBidService
	workService    SvWorkService
}

// ShowAccount godoc
// @Summary      GetUserByEmail supervisor's profile
// @Tags         supervisor
// @Produce      json
// @Success      200  {object}  dto.SvProfile
// @Router       /api/supervisor/profile [get]
// @Security	 Auth
func (ctrl *supervisor) getProfile(ctx echo.Context) error {
	email, _ := misc.ExtractCtx(ctx)
	ctrl.l.Debug(fmt.Sprintf("Email: %s", email))

	respDTO, err := ctrl.profileService.GetSupervisorProfile(email)
	if err != nil {
		ctrl.l.Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "TODO")
	}

	return ctx.JSON(http.StatusOK, respDTO)
}

// ShowAccount godoc
// @Summary      GetUserByEmail supervisor's works
// @Tags         supervisor
// @Param        supervisor_id query int  true  "Supervisor ID"
// @Produce      json
// @Success      200  {object}  dto.SvWorkPlenty
// @Router       /api/supervisor/work [get]
// @Security	 Auth
func (ctrl *supervisor) getWorks(ctx echo.Context) error {
	email, _ := misc.ExtractCtx(ctx)
	ctrl.l.Debug(fmt.Sprintf("Email: %s", email))

	supervisorID, _ := strconv.Atoi(ctx.QueryParam("supervisor_id"))

	respDTO, err := ctrl.workService.GetSupervisorWorks(supervisorID)
	if err != nil {
		ctrl.l.Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "TODO")
	}

	return ctx.JSON(http.StatusOK, respDTO)
}

// ShowAccount godoc
// @Summary      GetUserByEmail supervisor's bids
// @Tags         supervisor
// @Param        supervisor_id query int  true  "Supervisor ID"
// @Produce      json
// @Success      200  {object}  dto.SvBids
// @Router       /api/supervisor/bid [get]
// @Security	 Auth
func (ctrl *supervisor) getBids(ctx echo.Context) error {
	email, _ := misc.ExtractCtx(ctx)
	ctrl.l.Debug(fmt.Sprintf("Email: %s", email))

	supervisorID, _ := strconv.Atoi(ctx.QueryParam("supervisor_id"))

	respDTO, err := ctrl.bidService.GetSupervisorBids(supervisorID)
	if err != nil {
		ctrl.l.Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "TODO")
	}

	return ctx.JSON(http.StatusOK, respDTO)
}

// ShowAccount godoc
// @Summary      Accept or Decline student's bid
// @Tags         supervisor
// @Param 		 ResolveBid body dto.ResolveBid true "bid info"
// @Produce      json
// @Success      200  {object}  dto.ResolveBidResp
// @Router       /api/supervisor/bid/resolve [post]
// @Security	 Auth
func (ctrl *supervisor) resolveBid(ctx echo.Context) error {
	email, _ := misc.ExtractCtx(ctx)
	ctrl.l.Debug(fmt.Sprintf("Email: %s", email))

	reqDTO := &dto.ResolveBid{}
	if err := ctx.Bind(reqDTO); err != nil {
		return echo.ErrBadRequest
	}

	if err := ctrl.bidService.Resolve(reqDTO); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	newStatus := ""
	if reqDTO.Accept {
		newStatus = "accepted"
	} else {
		newStatus = "rejected"
	}

	resp := dto.ResolveBidResp{NewStatus: newStatus}

	return ctx.JSON(http.StatusOK, resp)
}

func NewSupervisorRoutes(
	router *echo.Group,
	l logger.Interface,
	profileService SvProfileService,
	bidService SvBidService,
	workService SvWorkService,
) {
	ctrl := &supervisor{l, profileService, bidService, workService}

	g := router.Group("/supervisor", middlewares.CheckRole)

	{
		g.GET("/profile", ctrl.getProfile)
		g.GET("/bid", ctrl.getBids)
		g.POST("/bid/resolve", ctrl.resolveBid)
		g.GET("/work", ctrl.getWorks)
		g.GET("/work", ctrl.getWorks)
	}

}
