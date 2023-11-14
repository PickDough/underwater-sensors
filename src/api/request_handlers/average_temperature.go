package request_handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"underwaterSensors/src/api/cqrs/query/average_temperature"
	"underwaterSensors/src/api/cqrs/query/group_exists"
	"underwaterSensors/src/api/di/di_container"
	"underwaterSensors/src/common/cqrs"
)

type AverageTemperatureHandler struct {
	container di_container.ApiContainer
}

func NewAverageTemperatureHandler(container di_container.ApiContainer) *AverageTemperatureHandler {
	return &AverageTemperatureHandler{
		container: container,
	}
}

func (h AverageTemperatureHandler) Handle(ctx *gin.Context) {
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

	query := average_temperature.AverageTemperatureQuery{
		Group:     ctx.Param("groupName"),
		Container: h.container,
	}

	res, err := cqrs.ExecuteQuery[average_temperature.AverageTemperatureResult, average_temperature.AverageTemperatureQuery](
		h.container.CQRS,
		query,
	)
	if err != nil {
		ctx.String(http.StatusInternalServerError, fmt.Sprintf("internal error"))
		return
	}

	ctx.JSONP(http.StatusOK, res)
}
