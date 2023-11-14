package build_currently_query

import (
	"github.com/uptrace/bun"
	"underwaterSensors/src/common/cqrs"
)

type CurrentlyQueryResult struct {
	Query *bun.SelectQuery
}
type CurrentlyQuery struct {
	DB *bun.DB
	cqrs.Query[CurrentlyQueryResult]
}
