package http

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"ssr/internal/controller/http/middlewares"
	"ssr/internal/dto"
	"ssr/internal/usecase"
	"ssr/pkg/logger"
	"ssr/pkg/misc"
	"strconv"
)

type supervisorRoutes struct {
	l         logger.Interface
	profileUC usecase.IProfileUC
	bidUC     usecase.ISupervisorBidUC
	workUC    usecase.ISupervisorWorkUC
}

// ShowAccount godoc
// @Summary      GetUserInfo supervisor's profile
// @Tags         supervisor
// @Produce      json
// @Success      200  {object}  dto.SupervisorProfile
// @Router       /api/supervisor/profile [get]
// @Security	 Auth
func (r *supervisorRoutes) getProfile(ctx echo.Context) error {
	email, _ := misc.ExtractCtx(ctx)
	r.l.Debug(fmt.Sprintf("Email: %s", email))

	respDTO, err := r.profileUC.GetSupervisorProfile(email)
	if err != nil {
		r.l.Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "TODO")
	}

	return ctx.JSON(http.StatusOK, respDTO)
}

// ShowAccount godoc
// @Summary      GetUserInfo supervisor's works
// @Tags         supervisor
// @Param        supervisor_id query int  true  "Supervisor ID"
// @Produce      json
// @Success      200  {object}  dto.SupervisorWorkPlenty
// @Router       /api/supervisor/work [get]
// @Security	 Auth
func (r *supervisorRoutes) getWorks(ctx echo.Context) error {
	email, _ := misc.ExtractCtx(ctx)
	r.l.Debug(fmt.Sprintf("Email: %s", email))

	supervisorID, _ := strconv.Atoi(ctx.QueryParam("supervisor_id"))

	respDTO, err := r.workUC.GetSupervisorWorks(supervisorID)
	if err != nil {
		r.l.Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "TODO")
	}

	return ctx.JSON(http.StatusOK, respDTO)
}

// ShowAccount godoc
// @Summary      GetUserInfo supervisor's bids
// @Tags         supervisor
// @Param        supervisor_id query int  true  "Supervisor ID"
// @Produce      json
// @Success      200  {object}  dto.SupervisorBids
// @Router       /api/supervisor/bid [get]
// @Security	 Auth
func (r *supervisorRoutes) getBids(ctx echo.Context) error {
	email, _ := misc.ExtractCtx(ctx)
	r.l.Debug(fmt.Sprintf("Email: %s", email))

	supervisorID, _ := strconv.Atoi(ctx.QueryParam("supervisor_id"))

	respDTO, err := r.bidUC.GetSupervisorBids(supervisorID)
	if err != nil {
		r.l.Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "TODO")
	}

	return ctx.JSON(http.StatusOK, respDTO)
}

// ShowAccount godoc
// @Summary      Accept or Decline student's bid
// @Tags         supervisor
// @Param 		 ResolveBid body dto.ResolveBid true "bid info"
// @Produce      json
// @Success      200
// @Router       /api/supervisor/bid/resolve [post]
// @Security	 Auth
func (r *supervisorRoutes) resolveBid(ctx echo.Context) error {
	email, _ := misc.ExtractCtx(ctx)
	r.l.Debug(fmt.Sprintf("Email: %s", email))

	reqDTO := &dto.ResolveBid{}
	if err := ctx.Bind(reqDTO); err != nil {
		return echo.ErrBadRequest
	}

	err := r.bidUC.Resolve(reqDTO)
	if err != nil {
		r.l.Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "TODO")
	}

	return ctx.NoContent(http.StatusOK)
}

func NewSupervisorRoutes(
	router *echo.Group,
	l logger.Interface,
	profileUC usecase.IProfileUC,
	bidUC usecase.ISupervisorBidUC,
	workUC usecase.ISupervisorWorkUC,
) {
	r := &supervisorRoutes{l, profileUC, bidUC, workUC}

	g := router.Group("/supervisor", middlewares.CheckRole)

	{
		g.GET("/profile", r.getProfile)
		g.GET("/bid", r.getBids)
		g.POST("/bid/resolve", r.resolveBid)
		g.GET("/work", r.getWorks)
		g.GET("/work", r.getWorks)
	}

}
