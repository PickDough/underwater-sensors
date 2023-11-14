package species

import (
	"underwaterSensors/src/api/di/di_container"
	"underwaterSensors/src/common/cqrs"
	"underwaterSensors/src/common/domain/sensor/dto"
)

type SpeciesResult struct {
	SensedFishes []*dto.SensedFish
}

type SpeciesQuery struct {
	Group     string
	Container di_container.ApiContainer
	Limit     int
	cqrs.Query[SpeciesResult]
}
