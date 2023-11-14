package average_temperature

import (
	"underwaterSensors/src/api/di/di_container"
	"underwaterSensors/src/common/cqrs"
)

type AverageTemperatureResult struct {
	TemperatureC float32 `json:"temperature_c"`
}

type AverageTemperatureQuery struct {
	Group     string
	Container di_container.ApiContainer
	cqrs.Query[AverageTemperatureResult]
}
