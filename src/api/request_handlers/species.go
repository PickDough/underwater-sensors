package request_handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"underwaterSensors/src/api/cqrs/query/group_exists"
	"underwaterSensors/src/api/cqrs/query/species"
	"underwaterSensors/src/api/di/di_container"
	"underwaterSensors/src/common/cqrs"
)

type SpeciesHandler struct {
	container di_container.ApiContainer
}

func NewSpeciesHandler(container di_container.ApiContainer) *SpeciesHandler {
	return &SpeciesHandler{
		container: container,
	}
}

func (h SpeciesHandler) Handle(ctx *gin.Context) {
	existsQuery := group_exists.GroupExistsQuery{
		Group:     ctx.Param("groupName"),
		Container: h.container,
	}
	exists, err := cqrs.ExecuteQuery[bool, group_exists.GroupExistsQuery](h.container.CQRS, existsQuery)
	if err != nil {
		ctx.String(http.StatusInternalServerError, fmt.Sprintf("internal error"))
		return
	}
	if !*exists {
		ctx.String(http.StatusBadRequest, fmt.Sprintf("group with name \"%s\" doesn't exist", ctx.Param("groupName")))
		return
	}

	query := species.SpeciesQuery{
		Group:     ctx.Param("groupName"),
		Container: h.container,
	}

	res, err := cqrs.ExecuteQuery[species.SpeciesResult, species.SpeciesQuery](
		h.container.CQRS,
		query,
	)
	if err != nil {
		ctx.String(http.StatusInternalServerError, fmt.Sprintf("internal error"))
		fmt.Println(err)
	}

	ctx.JSONP(http.StatusOK, res.SensedFishes)
}
