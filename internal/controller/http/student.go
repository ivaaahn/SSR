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

type studentRoutes struct {
	l         logger.Interface
	profileUC usecase.IProfileUC
	bidsUC    usecase.IStudentBidUC
	worksUC   usecase.IStudentWorkUC
	ssrUC     usecase.IStudentRelUC
}

func (r *studentRoutes) getProfile(ctx echo.Context) error {
	email, _ := misc.ExtractInfoFromContext(ctx)
	r.l.Debug(fmt.Sprintf("Email: %s", email))

	respDTO, err := r.profileUC.GetStudentProfile(email)
	if err != nil {
		r.l.Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "TODO")
	}

	return ctx.JSON(http.StatusOK, respDTO)
}

func (r *studentRoutes) getBids(ctx echo.Context) error {
	email, _ := misc.ExtractInfoFromContext(ctx)
	r.l.Debug(fmt.Sprintf("Email: %s", email))

	studentID, _ := strconv.Atoi(ctx.QueryParam("student_id"))

	respDTO, err := r.bidsUC.GetStudentBids(studentID)
	if err != nil {
		r.l.Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "TODO")
	}

	return ctx.JSON(http.StatusOK, respDTO)
}

func (r *studentRoutes) getWorks(ctx echo.Context) error {
	email, _ := misc.ExtractInfoFromContext(ctx)
	r.l.Debug(fmt.Sprintf("Email: %s", email))

	studentID, _ := strconv.Atoi(ctx.QueryParam("student_id"))

	respDTO, err := r.worksUC.GetStudentWorks(studentID)
	if err != nil {
		r.l.Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "TODO")
	}

	return ctx.JSON(http.StatusOK, respDTO)
}

func (r *studentRoutes) getSupervisorsOfWork(ctx echo.Context) error {
	email, _ := misc.ExtractInfoFromContext(ctx)
	r.l.Debug(fmt.Sprintf("Email: %s", email))

	workID, _ := strconv.Atoi(ctx.QueryParam("work_id"))

	respDTO, err := r.worksUC.GetWorkSupervisors(workID)
	if err != nil {
		r.l.Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "TODO")
	}

	return ctx.JSON(http.StatusOK, respDTO)
}

func (r *studentRoutes) applyBid(ctx echo.Context) error {
	email, _ := misc.ExtractInfoFromContext(ctx)
	r.l.Debug(fmt.Sprintf("Email: %s", email))

	reqDTO := &dto.ApplyBid{}
	if err := ctx.Bind(reqDTO); err != nil {
		return echo.ErrBadRequest
	}

	respDTO, err := r.bidsUC.Apply(reqDTO)
	if err != nil {
		r.l.Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "TODO")
	}

	return ctx.JSON(http.StatusCreated, respDTO)
}

func (r *studentRoutes) createSSR(ctx echo.Context) error {
	email, _ := misc.ExtractInfoFromContext(ctx)
	r.l.Debug(fmt.Sprintf("Email: %s", email))

	reqDTO := &dto.CreateSSR{}
	if err := ctx.Bind(reqDTO); err != nil {
		return echo.ErrBadRequest
	}

	respDTO, err := r.ssrUC.Create(reqDTO)
	if err != nil {
		r.l.Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "TODO")
	}

	return ctx.JSON(http.StatusCreated, respDTO)
}

func NewStudentRoutes(
	router *echo.Group,
	l logger.Interface,
	profileUC usecase.IProfileUC,
	bidsUC usecase.IStudentBidUC,
	worksUC usecase.IStudentWorkUC,
	ssrUC usecase.IStudentRelUC,
) {
	r := &studentRoutes{l, profileUC, bidsUC, worksUC, ssrUC}

	g := router.Group("/student", middlewares.CheckRole)

	{
		g.GET("/profile", r.getProfile)
		g.GET("/bid", r.getBids)
		g.PUT("/bid", r.applyBid)
		g.POST("/ssr", r.createSSR)
		g.GET("/work", r.getWorks)
		g.GET("/work/supervisor", r.getSupervisorsOfWork)
	}

}
