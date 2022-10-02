package http

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"ssr/config"
	"ssr/internal/controller/http/middlewares"
	"ssr/internal/dto"
	"ssr/pkg/logger"
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
// @Success      200  {object}  dto.RelationCreateResp
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
		relations.POST("/", ctrl.create)
	}

}
