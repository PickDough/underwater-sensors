package sensor_exists

import (
	"underwaterSensors/src/api/di/di_container"
	"underwaterSensors/src/common/cqrs"
)

type SensorExistsQuery struct {
	Group     string
	Index     int
	Container di_container.ApiContainer
	cqrs.Query[bool]
}
