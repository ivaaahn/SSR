package http

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"ssr/config"
	"ssr/internal/controller/http/middlewares"
	"ssr/pkg/logger"
	"strconv"
)

type student struct {
	l                logger.Interface
	profileService   ProfileService
	workService      WorkService
	relationsService RelationsService
	feedbackService  FeedbackService
}

// ShowAccount godoc
// @Summary      Get student's profile
// @Tags         student
// @Produce      json
// @Param        student_id path int  true  "Student ID"
// @Success      200  {object}  dto.StProfile
// @Failure      404
// @Router       /api/v1/students/{student_id}/profile [get]
// @Security	 OAuth2Password
func (ctrl *student) getProfile(ctx echo.Context) error {
	studentID, _ := strconv.Atoi(ctx.Param("student_id"))

	profileDTO, err := ctrl.profileService.GetStudentProfile(studentID)
	if err != nil {
		return echo.ErrNotFound
	}

	return ctx.JSON(http.StatusOK, profileDTO)
}

// ShowAccount godoc
// @Summary      Get student's works
// @Tags         student
// @Param        student_id path int  true  "Student ID"
// @Produce      json
// @Success      200  {object}  dto.StWorkPlenty
// @Router       /api/v1/students/{student_id}/works [get]
// @Security	 OAuth2Password
func (ctrl *student) getWorks(ctx echo.Context) error {
	studentID, _ := strconv.Atoi(ctx.Param("student_id"))

	respDTO, err := ctrl.workService.GetStudentWorks(studentID)
	if err != nil {
		return echo.ErrNotFound
	}

	return ctx.JSON(http.StatusOK, respDTO)
}

// ShowAccount godoc
// @Summary      Get student's bids
// @Tags         student
// @Produce      json
// @Param        student_id path int  true  "Student ID"
// @Success      200  {object}  dto.StRelationPlenty
// @Failure      404
// @Router       /api/v1/students/{student_id}/relations [get]
// @Security	 OAuth2Password
func (ctrl *student) getRelations(ctx echo.Context) error {
	studentID, _ := strconv.Atoi(ctx.Param("student_id"))

	respDTO, err := ctrl.relationsService.GetStudentRelations(studentID)
	if err != nil {
		return echo.ErrNotFound
	}

	return ctx.JSON(http.StatusOK, respDTO)
}

//// ShowAccount godoc
//// @Summary      Start SSR
//// @Tags         student
//// @Accept		 json
//// @Param 		 ApplyBid body dto.CreateSSR true "ssr info"
//// @Produce      json
//// @Success      200  {object}  dto.StViewRelation
//// @Router       /api/student/ssr [post]
//// @Security	 Auth
//func (ctrl *student) createSSR(ctx echo.Context) error {
//	email, _ := misc.ExtractCtx(ctx)
//	ctrl.l.Debug(fmt.Sprintf("Email: %s", email))
//
//	reqDTO := &dto.CreateSSR{}
//	if err := ctx.Bind(reqDTO); err != nil {
//		return echo.ErrBadRequest
//	}
//
//	respDTO, err := ctrl.relationService.Accept(reqDTO)
//	if err != nil {
//		return echo.ErrInternalServerError
//	}
//
//	return ctx.JSON(http.StatusCreated, respDTO)
//}

//// ShowAccount godoc
//// @Summary      Provide a feedback
//// @Tags         student
//// @Accept		 json
//// @Param 		 Feedback body dto.FeedbackReq true "feedback info"
//// @Produce      json
//// @Success      201  {object}  dto.FeedbackAddResp
//// @Failure      500
//// @Router       /api/student/feedback [put]
//// @Security	 Authorization
//func (ctrl *student) provideFeedback(ctx echo.Context) error {
//	email, _ := misc.ExtractCtx(ctx)
//	ctrl.l.Debug(fmt.Sprintf("Email: %s", email))
//
//	reqDTO := &dto.FeedbackReq{}
//	if err := ctx.Bind(reqDTO); err != nil {
//		return echo.ErrBadRequest
//	}
//
//	id, err := ctrl.feedbackService.Add(reqDTO)
//	if err != nil {
//		return echo.NewHTTPError(http.StatusConflict)
//	}
//
//	return ctx.JSON(http.StatusCreated, dto.FeedbackAddResp{FeedbackID: id})
//}

//// ShowAccount godoc
//// @Summary      Get feedbacks on the supervisor.
//// @Tags         student
//// @Param        supervisor_id path integer true "Supervisor ID"
//// @Produce      json
//// @Success      200  {object}  dto.FeedbackPlenty
//// @Router       /api/student/feedback [get]
//// @Security	 Auth
//func (ctrl *student) getFeedback(ctx echo.Context) error {
//	email, _ := misc.ExtractCtx(ctx)
//	ctrl.l.Debug(fmt.Sprintf("Email: %s", email))
//
//	supervisorID, _ := strconv.Atoi(ctx.Param("supervisor_id"))
//
//	respDTO, err := ctrl.feedbackService.GetOnSupervisor(supervisorID)
//	if err != nil {
//		return echo.ErrInternalServerError
//	}
//
//	return ctx.JSON(http.StatusOK, respDTO)
//}

func NewStudentRoutes(
	router *echo.Group,
	l logger.Interface,
	config *config.Config,
	profileService ProfileService,
	worksService WorkService,
	relationService RelationsService,
	feedbackService FeedbackService,
) {
	ctrl := &student{
		l,
		profileService,
		worksService,
		relationService,
		feedbackService,
	}

	student := router.Group("/students", middlewares.MakeAuthMiddleware(config), middlewares.CheckRole)

	{
		student.GET("/:student_id/profile", ctrl.getProfile)
		student.GET("/:student_id/works", ctrl.getWorks)
		student.GET("/:student_id/relations", ctrl.getRelations)
		//student.PUT("/bid", ctrl.applyBid)
		//student.POST("/ssr", ctrl.createSSR)
		//student.GET("/work/supervisor_id", ctrl.getSupervisors)
		//student.GET("/feedback/:supervisor_id", ctrl.getFeedback)
		//student.PUT("/feedback", ctrl.provideFeedback)
	}

}
