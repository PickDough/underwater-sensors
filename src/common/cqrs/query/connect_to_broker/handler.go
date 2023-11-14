package connect_to_broker

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

func Connect(query MessageBrokerQuery) (*MessageBroker, error) {
	conn, err := amqp.Dial(query.Url)
	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	q, err := ch.QueueDeclare(
		query.Queue, // name
		false,       // durable
		false,       // delete when unused
		false,       // exclusive
		false,       // no-wait
		nil,         // arguments
	)
	if err != nil {
		return nil, err
	}

	return &MessageBroker{conn, ch, q}, nil
}
