package request_handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"underwaterSensors/src/api/cqrs/query/average_temperature_by_sensor"
	"underwaterSensors/src/api/cqrs/query/sensor_exists"
	"underwaterSensors/src/api/di/di_container"
	"underwaterSensors/src/common/cqrs"
)

type AverageTemperatureBySensorHandler struct {
	container di_container.ApiContainer
	regex     *regexp.Regexp
}

func NewAverageTemperatureBySensorHandler(container di_container.ApiContainer) *AverageTemperatureBySensorHandler {
	r, _ := regexp.Compile("[a-zA-z]+-[0-9]+")
	return &AverageTemperatureBySensorHandler{
		container: container,
		regex:     r,
	}
}

func (h AverageTemperatureBySensorHandler) Handle(ctx *gin.Context) {
	param := ctx.Param("codename")
	if !h.regex.Match([]byte(param)) {
		ctx.String(http.StatusBadRequest, fmt.Sprintf("codename is incorrect"))
		return
	}

	from := ctx.Query("from")
	till := ctx.Query("till")
	if from == "" || till == "" {
		ctx.String(http.StatusBadRequest, fmt.Sprintf("time boundaries are not specified"))
		return
	}
	fromUnix, err := strconv.ParseInt(from, 10, 64)
	if err != nil {
		ctx.String(http.StatusBadRequest, fmt.Sprintf("incorrect from format: %s", from))
		return
	}
	tillUnix, err := strconv.ParseInt(till, 10, 64)
	if err != nil {
		ctx.String(http.StatusBadRequest, fmt.Sprintf("incorrect till format: %s", till))
		return
	}

	strs := strings.Split(param, "-")
	group := strs[0]
	index, _ := strconv.Atoi(strs[1])

	existsQuery := sensor_exists.SensorExistsQuery{
		Group:     group,
		Index:     index,
		Container: h.container,
	}
	exists, err := cqrs.ExecuteQuery[bool, sensor_exists.SensorExistsQuery](h.container.CQRS, existsQuery)
	if err != nil {
		ctx.String(http.StatusInternalServerError, fmt.Sprintf("internal error"))
		return
	}
	if !*exists {
		ctx.String(http.StatusBadRequest, fmt.Sprintf("sensor \"%s-%d\" doesn't exist", group, index))
		return
	}

	averageQuery := average_temperature_by_sensor.AverageTemperatureBySensorQuery{
		Group:         group,
		Index:         index,
		FromDateTime:  fromUnix,
		UntilDateTime: tillUnix,
		Container:     h.container,
	}

	temp, err := cqrs.ExecuteQuery[
		average_temperature_by_sensor.AverageTemperatureBySensorResult,
		average_temperature_by_sensor.AverageTemperatureBySensorQuery,
	](h.container.CQRS, averageQuery)
	if err != nil {
		ctx.String(http.StatusInternalServerError, fmt.Sprintf("internal error"))
		return
	}

	ctx.JSONP(http.StatusOK, temp)
}
