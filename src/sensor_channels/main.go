package main

import (
	"context"
	"encoding/json"
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"os"
	"underwaterSensors/src/common/cqrs"
	"underwaterSensors/src/common/cqrs/query/connect_to_broker"
	"underwaterSensors/src/sensor_channels/cqrs/query/construct_sensors"
	"underwaterSensors/src/sensor_channels/cqrs/query/extract_config"
	"underwaterSensors/src/sensor_channels/di"
	"underwaterSensors/src/sensor_channels/domain/config"
	"underwaterSensors/src/sensor_channels/domain/faker"
)

func main() {
	diContainer := di.BuildChannelContainer()
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

	cfg, err := cqrs.ExecuteQuery[config.Config, extract_config.ConfigQuery](
		diContainer.CQRS,
		extract_config.ConfigQuery{
			File: os.Getenv("CONFIG_PATH"),
		},
	)
	fmt.Println(cfg.Fishes)

	fakeSensors, err := cqrs.ExecuteQuery[construct_sensors.ConstructResult, construct_sensors.ConstructQuery](
		diContainer.CQRS,
		construct_sensors.ConstructQuery{
			Config: cfg,
		},
	)

	for _, sensor := range fakeSensors.Sensors {
		go sensor.Sense()
		go func(fake *faker.SensorFaker) {
			for {
				sensedData := <-fake.Channel
				jsonSensedData, _ := json.Marshal(sensedData)
				err = mbq.Ch.PublishWithContext(context.Background(),
					"",         // exchange
					mbq.Q.Name, // routing key
					false,      // mandatory
					false,      // immediate
					amqp.Publishing{
						ContentType: "text/plain",
						Body:        jsonSensedData,
					})
			}
		}(sensor)
	}

	for {
	}
}

func failOnError(err error) {
	if err != nil {
		log.Panicf("%s", err)
	}
}
