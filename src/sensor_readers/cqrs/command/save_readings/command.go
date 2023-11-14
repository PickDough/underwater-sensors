package save_readings

import (
	"underwaterSensors/src/common/cqrs"
	"underwaterSensors/src/common/domain/sensor/dto"
	"underwaterSensors/src/sensor_readers/di/di_container"
)

type SaveReadingsCommand struct {
	Readings  *dto.SensedData
	Container di_container.SensorReadersContainer
	cqrs.Command
}
