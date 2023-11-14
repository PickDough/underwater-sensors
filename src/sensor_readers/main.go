package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"underwaterSensors/src/common/cqrs"
	"underwaterSensors/src/common/cqrs/query/connect_to_broker"
	"underwaterSensors/src/common/domain/sensor/dto"
	"underwaterSensors/src/sensor_readers/cqrs/command/save_readings"
	"underwaterSensors/src/sensor_readers/di"
)

func main() {
	diContainer := di.BuildReaderContainer()
	mbq, err := cqrs.ExecuteQuery[connect_to_broker.MessageBroker, connect_to_broker.MessageBrokerQuery](
		diContainer.CQRS,
		connect_to_broker.MessageBrokerQuery{
			Url:   os.Getenv("RABBITMQ_CONN"),
			Queue: os.Getenv("RABBITMQ_QUEUE"),
		},
	)
	if err != nil {
		failOnError(err)
	}
	defer mbq.Ch.Close()
	defer mbq.Ch.Close()

	msgs, err := mbq.Ch.Consume(
		mbq.Q.Name, // queue
		"",         // consumer
		true,       // auto-ack
		false,      // exclusive
		false,      // no-local
		false,      // no-wait
		nil,        // args
	)
	if err != nil {
		failOnError(err)
	}

	go func() {
		for d := range msgs {
			var data dto.SensedData
			err := json.Unmarshal(d.Body, &data)
			if err != nil {
				return
			}
			fmt.Println(data)

			err = cqrs.ExecuteCommand[save_readings.SaveReadingsCommand](diContainer.CQRS, save_readings.SaveReadingsCommand{
				Readings:  &data,
				Container: diContainer,
			})
			if err != nil {
				fmt.Println(err)
				continue
			}
		}
	}()

	for {
	}
}

func failOnError(err error) {
	if err != nil {
		log.Panicf("%s", err)
	}
}
