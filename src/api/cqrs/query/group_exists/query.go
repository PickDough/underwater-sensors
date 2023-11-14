package group_exists

import (
	"underwaterSensors/src/api/di/di_container"
	"underwaterSensors/src/common/cqrs"
)

type GroupExistsQuery struct {
	Group     string
	Container di_container.ApiContainer
	cqrs.Query[bool]
}
