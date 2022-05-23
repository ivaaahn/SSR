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

// ShowAccount godoc
// @Summary      Get student's profile
// @Tags         student
// @Produce      json
// @Success      200  {object}  dto.StudentProfile
// @Router       /api/student/profile [get]
// @Security	 Auth
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

// ShowAccount godoc
// @Summary      Get student's bids
// @Tags         student
// @Produce      json
// @Param        student_id query int  true  "Student ID"
// @Success      200  {object}  dto.StudentBids
// @Router       /api/student/bid [get]
// @Security	 Auth
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

// ShowAccount godoc
// @Summary      Get student's works
// @Tags         student
// @Param        student_id query int  true  "Student ID"
// @Produce      json
// @Success      200  {object}  dto.StudentWorkPlenty
// @Router       /api/student/work [get]
// @Security	 Auth
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

// ShowAccount godoc
// @Summary      Get supervisors of the work
// @Tags         student
// @Param        work_id query int  true  "Work ID"
// @Produce      json
// @Success      200  {object}  dto.WorkSupervisorPlenty
// @Router       /api/student/work/supervisor [get]
// @Security	 Auth
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

// ShowAccount godoc
// @Summary      Apply bid
// @Tags         student
// @Accept		 json
// @Param 		 ApplyBid body dto.ApplyBid true "bid info"
// @Produce      json
// @Success      200  {object}  dto.ApplyBidResponse
// @Router       /api/student/bid [put]
// @Security	 Auth
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

// ShowAccount godoc
// @Summary      Start SSR
// @Tags         student
// @Accept		 json
// @Param 		 ApplyBid body dto.CreateSSR true "ssr info"
// @Produce      json
// @Success      200  {object}  dto.StudentViewSSR
// @Router       /api/student/ssr [post]
// @Security	 Auth
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
