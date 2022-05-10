package http

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"ssr/internal/controller/http/middlewares"
	"ssr/internal/usecase"
	"ssr/pkg/logger"
	"ssr/pkg/misc"
	"strconv"
)

type supervisorRoutes struct {
	l         logger.Interface
	profileUC usecase.IProfileUC
	bidsUC    usecase.ISupervisorBidUC
}

func (r *supervisorRoutes) getProfile(ctx echo.Context) error {
	email, _ := misc.ExtractInfoFromContext(ctx)
	r.l.Debug(fmt.Sprintf("Email: %s", email))

	respDTO, err := r.profileUC.GetSupervisorProfile(email)
	if err != nil {
		r.l.Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "TODO")
	}

	return ctx.JSON(http.StatusOK, respDTO)
}

func (r *supervisorRoutes) getBids(ctx echo.Context) error {
	email, _ := misc.ExtractInfoFromContext(ctx)
	r.l.Debug(fmt.Sprintf("Email: %s", email))

	supervisorID, _ := strconv.Atoi(ctx.QueryParam("supervisor_id"))

	respDTO, err := r.bidsUC.GetSupervisorBids(supervisorID)
	if err != nil {
		r.l.Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "TODO")
	}

	return ctx.JSON(http.StatusOK, respDTO)
}

func NewSupervisorRoutes(
	router *echo.Group,
	l logger.Interface,
	profileUC usecase.IProfileUC,
	bidsUC usecase.ISupervisorBidUC,

) {
	r := &supervisorRoutes{l, profileUC, bidsUC}

	g := router.Group("/supervisor", middlewares.CheckRole)

	{
		g.GET("/profile", r.getProfile)
		g.GET("/bid", r.getBids)
	}

}
