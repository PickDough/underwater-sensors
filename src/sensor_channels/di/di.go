package di

import (
	"underwaterSensors/src/common/cqrs"
	"underwaterSensors/src/common/di"
	"underwaterSensors/src/sensor_channels/cqrs/query/construct_sensors"
	"underwaterSensors/src/sensor_channels/cqrs/query/extract_config"
	"underwaterSensors/src/sensor_channels/domain/config"
)

type SensorChannelsContainer struct {
	CQRS *cqrs.CQRS
}

func BuildChannelContainer() SensorChannelsContainer {
	cq := di.BuildCommonContainer().CQRS

	cqrs.RegisterQuery[config.Config](
		cq,
		extract_config.ConfigQuery{},
		extract_config.Handle,
	)
	cqrs.RegisterQuery[construct_sensors.ConstructResult](
		cq,
		construct_sensors.ConstructQuery{},
		construct_sensors.ConstructSensors,
	)
	return SensorChannelsContainer{cq}
}
