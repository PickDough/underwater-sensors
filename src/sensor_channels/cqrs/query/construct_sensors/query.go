package construct_sensors

import (
	"underwaterSensors/src/common/cqrs"
	"underwaterSensors/src/sensor_channels/domain/config"
	"underwaterSensors/src/sensor_channels/domain/faker"
)

type ConstructResult struct {
	Sensors []*faker.SensorFaker
}

type ConstructQuery struct {
	Config *config.Config
	cqrs.Query[config.Config]
}
