package di_container

import (
	"github.com/uptrace/bun"
	"underwaterSensors/src/common/cqrs"
)

type SensorReadersContainer struct {
	CQRS *cqrs.CQRS
	DB   *bun.DB
}
