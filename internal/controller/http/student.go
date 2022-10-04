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
}

// ShowAccount godoc
// @Summary      Get student's profile
// @Tags         student
// @Produce      json
// @Param        student_id path int  true  "Student ID"
// @Success      200  {object}  dto.Student
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
// @Success      200  {object}  dto.StudentViewWorkPlenty
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

func NewStudentRoutes(
	router *echo.Group,
	l logger.Interface,
	config *config.Config,
	profileService ProfileService,
	worksService WorkService,
	relationService RelationsService,
) {
	ctrl := &student{
		l,
		profileService,
		worksService,
		relationService,
	}

	student := router.Group("/students", middlewares.MakeAuthMiddleware(config), middlewares.CheckRole)

	{
		student.GET("/:student_id/profile", ctrl.getProfile)
		student.GET("/:student_id/works", ctrl.getWorks)
	}

}
