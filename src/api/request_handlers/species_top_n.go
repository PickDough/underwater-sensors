package request_handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"underwaterSensors/src/api/cqrs/query/group_exists"
	"underwaterSensors/src/api/cqrs/query/species"
	"underwaterSensors/src/api/di/di_container"
	"underwaterSensors/src/common/cqrs"
)

type SpeciesTopNHandler struct {
	container di_container.ApiContainer
}

func NewSpeciesTopNHandler(container di_container.ApiContainer) *SpeciesTopNHandler {
	return &SpeciesTopNHandler{
		container: container,
	}
}

func (h SpeciesTopNHandler) Handle(ctx *gin.Context) {
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

	limit, err := strconv.Atoi(ctx.Param("N"))
	if err != nil {
		ctx.String(http.StatusBadRequest, fmt.Sprintf("N is not a number: %s", ctx.Param("N")))
		return
	}
	query := species.SpeciesQuery{
		Group:     ctx.Param("groupName"),
		Limit:     limit,
		Container: h.container,
	}

	res, err := cqrs.ExecuteQuery[species.SpeciesResult, species.SpeciesQuery](
		h.container.CQRS,
		query,
	)
	if err != nil {
		ctx.String(http.StatusInternalServerError, fmt.Sprintf("internal error"))
		fmt.Println(err)
		return
	}

	ctx.JSONP(http.StatusOK, res.SensedFishes)
}
