package http

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"ssr/config"
	"ssr/internal/controller/http/middlewares"
	"ssr/internal/dto"
	"ssr/pkg/logger"
	"strconv"
)

type relation struct {
	l               logger.Interface
	relationService RelationsService
}

// ShowAccount godoc
// @Summary      Create relation
// @Tags         relation
// @Accept		 json
// @Param 		 CreateRelation body dto.RelationCreateReq true "Relation data"
// @Produce      json
// @Success      201  {object}  dto.RelationCreateResp
// @Router       /api/v1/relations/ [post]
// @Security	 OAuth2Password
func (ctrl *relation) create(ctx echo.Context) error {
	reqDTO := &dto.RelationCreateReq{}

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
// @Summary      Update relation
// @Tags         relation
// @Accept		 json
// @Param 		 CreateRelation body dto.RelationUpdateReq true "Relation data"
// @Produce      json
// @Success      200  {object}  dto.RelationResp
// @Router       /api/v1/relations/ [patch]
// @Security	 OAuth2Password
func (ctrl *relation) update(ctx echo.Context) error {
	reqDTO := &dto.RelationUpdateReq{}

	if err := ctx.Bind(reqDTO); err != nil {
		return echo.ErrBadRequest
	}

	respDTO, err := ctrl.relationService.Update(reqDTO)
	if err != nil {
		return echo.ErrInternalServerError
	}

	return ctx.JSON(http.StatusCreated, respDTO)
}

// ShowAccount godoc
// @Summary      Get relations
// @Tags         relation
// @Produce      json
// @Param        student_id query int  false  "Student ID"
// @Param        supervisor_id query int  false  "Supervisor ID"
// @Success      200  {object}  dto.RelationPlenty
// @Failure      404
// @Router       /api/v1/relations/ [get]
// @Security	 OAuth2Password
func (ctrl *relation) getPlenty(ctx echo.Context) error {
	studentID, _ := strconv.Atoi(ctx.QueryParam("student_id"))
	supervisorID, _ := strconv.Atoi(ctx.QueryParam("supervisor_id"))

	if studentID < 0 || supervisorID < 0 {
		return echo.NewHTTPError(400, map[string]string{"err": "Bad request", "msg": "Parameter (student_id, supervisor_id) must be positive"})
	}

	if studentID == 0 && supervisorID == 0 {
		return echo.NewHTTPError(400, map[string]string{"err": "Bad request", "msg": "Must be passed at least one parameter (student_id, supervisor_id)"})
	}

	if studentID != 0 && supervisorID != 0 {
		return echo.NewHTTPError(400, map[string]string{"err": "Bad request", "msg": "Must be passed only on parameter (student_id, supervisor_id)"})
	}

	respDTO, err := ctrl.relationService.GetPlenty(studentID, supervisorID)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, respDTO)
}

func NewRelationRoutes(
	router *echo.Group,
	l logger.Interface,
	config *config.Config,
	relationsService RelationsService,
) {
	ctrl := &relation{
		l,
		relationsService,
	}

	relations := router.Group("/relations", middlewares.MakeAuthMiddleware(config))

	{
		relations.GET("/", ctrl.getPlenty)
		relations.POST("/", ctrl.create)
		relations.PATCH("/", ctrl.update)
	}

}
