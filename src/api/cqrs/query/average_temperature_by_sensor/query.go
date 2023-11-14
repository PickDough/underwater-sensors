package average_temperature_by_sensor

import (
	"underwaterSensors/src/api/di/di_container"
	"underwaterSensors/src/common/cqrs"
)

type AverageTemperatureBySensorResult struct {
	TemperatureC float32 `json:"temperature_c"`
}

type AverageTemperatureBySensorQuery struct {
	Group         string
	Index         int
	FromDateTime  int64
	UntilDateTime int64
	Container     di_container.ApiContainer
	cqrs.Query[AverageTemperatureBySensorResult]
}
