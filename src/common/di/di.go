package di

import (
	"underwaterSensors/src/common/cqrs"
	"underwaterSensors/src/common/cqrs/query/build_currently_query"
)
import "underwaterSensors/src/common/cqrs/query/connect_to_broker"

type CommonContainer struct {
	CQRS *cqrs.CQRS
}

func BuildCommonContainer() *CommonContainer {
	cq := cqrs.New()

	cqrs.RegisterQuery[connect_to_broker.MessageBroker](
		cq,
		connect_to_broker.MessageBrokerQuery{},
		connect_to_broker.Connect,
	)
	cqrs.RegisterQuery[build_currently_query.CurrentlyQueryResult](
		cq,
		build_currently_query.CurrentlyQuery{},
		build_currently_query.Handle,
	)

	return &CommonContainer{cq}
}
