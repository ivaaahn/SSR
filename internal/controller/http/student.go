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

type student struct {
	l               logger.Interface
	profileService  StProfileService
	bidService      StBidService
	workService     StWorkService
	relationService StRelationService
	feedbackService FeedbackService
}

// ShowAccount godoc
// @Summary      GetUserByEmail student's profile
// @Tags         student
// @Produce      json
// @Success      200  {object}  dto.StProfile
// @Failure      404
// @Router       /api/student/profile [get]
// @Security	 Auth
func (ctrl *student) getProfile(ctx echo.Context) error {
	rawUserID, _ := misc.ExtractCtx(ctx)
	userID, _ := strconv.Atoi(rawUserID)

	profileDTO, err := ctrl.profileService.GetStudentProfile(userID)
	if err != nil {
		return echo.ErrNotFound
	}

	return ctx.JSON(http.StatusOK, profileDTO)
}

// ShowAccount godoc
// @Summary      GetUserByEmail student's bids
// @Tags         student
// @Produce      json
// @Param        student_id query int  true  "Student ID"
// @Success      200  {object}  dto.StBids
// @Failure      404
// @Router       /api/student/bid [get]
// @Security	 Auth
func (ctrl *student) getBids(ctx echo.Context) error {
	email, _ := misc.ExtractCtx(ctx)
	ctrl.l.Debug(fmt.Sprintf("Email: %s", email))

	studentID, _ := strconv.Atoi(ctx.QueryParam("student_id"))

	respDTO, err := ctrl.bidService.GetStudentBids(studentID)
	if err != nil {
		return echo.ErrNotFound
	}

	return ctx.JSON(http.StatusOK, respDTO)
}

// ShowAccount godoc
// @Summary      GetUserByEmail student's works
// @Tags         student
// @Param        student_id query int  true  "Student ID"
// @Produce      json
// @Success      200  {object}  dto.StWorkPlenty
// @Router       /api/student/work [get]
// @Security	 Auth
func (ctrl *student) getWorks(ctx echo.Context) error {
	email, _ := misc.ExtractCtx(ctx)
	ctrl.l.Debug(fmt.Sprintf("Email: %s", email))

	studentID, _ := strconv.Atoi(ctx.QueryParam("student_id"))

	respDTO, err := ctrl.workService.GetStudentWorks(studentID)
	if err != nil {
		return echo.ErrNotFound
	}

	return ctx.JSON(http.StatusOK, respDTO)
}

// ShowAccount godoc
// @Summary      GetUserByEmail supervisors of the work
// @Tags         student
// @Param        work_id query int  true  "WorkResp ID"
// @Produce      json
// @Success      200  {object}  dto.WorkSvPlenty
// @Router       /api/student/work/supervisor [get]
// @Security	 Auth
//func (ctrl *student) getSupervisorsOfWork(ctx echo.Context) error {
//	email, _ := misc.ExtractCtx(ctx)
//	ctrl.l.Debug(fmt.Sprintf("Email: %s", email))
//
//	workID, _ := strconv.Atoi(ctx.QueryParam("work_id"))
//
//	respDTO, err := ctrl.workService.GetWorkSupervisors(workID)
//	if err != nil {
//		return echo.ErrNotFound
//	}
//
//	return ctx.JSON(http.StatusOK, respDTO)
//}

// ShowAccount godoc
// @Summary      Apply bid
// @Tags         student
// @Accept		 json
// @Param 		 ApplyBid body dto.ApplyBid true "bid info"
// @Produce      json
// @Success      200  {object}  dto.ApplyBidResp
// @Router       /api/student/bid [put]
// @Security	 Auth
func (ctrl *student) applyBid(ctx echo.Context) error {
	email, _ := misc.ExtractCtx(ctx)
	ctrl.l.Debug(fmt.Sprintf("Email: %s", email))

	reqDTO := &dto.ApplyBid{}
	if err := ctx.Bind(reqDTO); err != nil {
		return echo.ErrBadRequest
	}

	respDTO, err := ctrl.bidService.Apply(reqDTO)
	if err != nil {
		return echo.ErrInternalServerError
	}

	return ctx.JSON(http.StatusCreated, respDTO)
}

// ShowAccount godoc
// @Summary      Start SSR
// @Tags         student
// @Accept		 json
// @Param 		 ApplyBid body dto.CreateSSR true "ssr info"
// @Produce      json
// @Success      200  {object}  dto.StViewRelation
// @Router       /api/student/ssr [post]
// @Security	 Auth
func (ctrl *student) createSSR(ctx echo.Context) error {
	email, _ := misc.ExtractCtx(ctx)
	ctrl.l.Debug(fmt.Sprintf("Email: %s", email))

	reqDTO := &dto.CreateSSR{}
	if err := ctx.Bind(reqDTO); err != nil {
		return echo.ErrBadRequest
	}

	respDTO, err := ctrl.relationService.Create(reqDTO)
	if err != nil {
		return echo.ErrInternalServerError
	}

	return ctx.JSON(http.StatusCreated, respDTO)
}

// ShowAccount godoc
// @Summary      Provide a feedback
// @Tags         student
// @Accept		 json
// @Param 		 Feedback body dto.FeedbackReq true "feedback info"
// @Produce      json
// @Success      201  {object}  dto.FeedbackAddResp
// @Failure      500
// @Router       /api/student/feedback [put]
// @Security	 Auth
func (ctrl *student) provideFeedback(ctx echo.Context) error {
	email, _ := misc.ExtractCtx(ctx)
	ctrl.l.Debug(fmt.Sprintf("Email: %s", email))

	reqDTO := &dto.FeedbackReq{}
	if err := ctx.Bind(reqDTO); err != nil {
		return echo.ErrBadRequest
	}

	id, err := ctrl.feedbackService.Add(reqDTO)
	if err != nil {
		return echo.NewHTTPError(http.StatusConflict)
	}

	return ctx.JSON(http.StatusCreated, dto.FeedbackAddResp{FeedbackID: id})
}

// ShowAccount godoc
// @Summary      Get feedbacks on the supervisor.
// @Tags         student
// @Param        supervisor_id path integer true "Supervisor ID"
// @Produce      json
// @Success      200  {object}  dto.FeedbackPlenty
// @Router       /api/student/feedback [get]
// @Security	 Auth
func (ctrl *student) getFeedback(ctx echo.Context) error {
	email, _ := misc.ExtractCtx(ctx)
	ctrl.l.Debug(fmt.Sprintf("Email: %s", email))

	supervisorID, _ := strconv.Atoi(ctx.Param("supervisor_id"))

	respDTO, err := ctrl.feedbackService.GetOnSupervisor(supervisorID)
	if err != nil {
		return echo.ErrInternalServerError
	}

	return ctx.JSON(http.StatusOK, respDTO)
}

func NewStudentRoutes(
	router *echo.Group,
	l logger.Interface,
	profileService StProfileService,
	bidsService StBidService,
	worksService StWorkService,
	relationService StRelationService,
	feedbackService FeedbackService,
) {
	ctrl := &student{
		l,
		profileService,
		bidsService,
		worksService,
		relationService,
		feedbackService,
	}

	student := router.Group("/student", middlewares.CheckRole)

	{
		student.GET("/profile", ctrl.getProfile)
		student.GET("/bid", ctrl.getBids)
		student.PUT("/bid", ctrl.applyBid)
		student.POST("/ssr", ctrl.createSSR)
		student.GET("/work", ctrl.getWorks)
		//student.GET("/work/supervisor_id", ctrl.getSupervisorsOfWork)
		student.GET("/feedback/:supervisor_id", ctrl.getFeedback)
		student.PUT("/feedback", ctrl.provideFeedback)
	}

}
