package extract_config

import (
	"underwaterSensors/src/common/cqrs"
	"underwaterSensors/src/sensor_channels/domain/config"
)

type ConfigQuery struct {
	File string
	cqrs.Query[config.Config]
}
