package connect_to_broker

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"underwaterSensors/src/common/cqrs"
)

type MessageBroker struct {
	Cn *amqp.Connection
	Ch *amqp.Channel
	Q  amqp.Queue
}

type MessageBrokerQuery struct {
	Url   string
	Queue string
	cqrs.Query[MessageBroker]
}
