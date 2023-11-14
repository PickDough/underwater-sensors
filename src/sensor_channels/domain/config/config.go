package config

import (
	"underwaterSensors/src/common/domain/sensor/dto"
)

type Config struct {
	Sensors []*dto.Sensor
	Fishes  []*dto.Fish
}
