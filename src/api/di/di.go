package di

import (
	"database/sql"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"
	"os"
	"underwaterSensors/src/api/cqrs/query/average_temperature"
	"underwaterSensors/src/api/cqrs/query/average_temperature_by_sensor"
	"underwaterSensors/src/api/cqrs/query/group_exists"
	"underwaterSensors/src/api/cqrs/query/sensor_exists"
	"underwaterSensors/src/api/cqrs/query/species"
	"underwaterSensors/src/api/di/di_container"
	"underwaterSensors/src/common/cqrs"
	"underwaterSensors/src/common/di"
)

func BuildApiContainer() di_container.ApiContainer {
	cq := di.BuildCommonContainer().CQRS

	cqrs.RegisterQuery[average_temperature.AverageTemperatureResult](
		cq,
		average_temperature.AverageTemperatureQuery{},
		average_temperature.Handle,
	)
	cqrs.RegisterQuery[species.SpeciesResult](
		cq,
		species.SpeciesQuery{},
		species.Handle,
	)
	cqrs.RegisterQuery[bool, group_exists.GroupExistsQuery](
		cq,
		group_exists.GroupExistsQuery{},
		group_exists.Handle,
	)
	cqrs.RegisterQuery[bool, sensor_exists.SensorExistsQuery](
		cq,
		sensor_exists.SensorExistsQuery{},
		sensor_exists.Handle,
	)
	cqrs.RegisterQuery[
		average_temperature_by_sensor.AverageTemperatureBySensorResult,
		average_temperature_by_sensor.AverageTemperatureBySensorQuery,
	](
		cq,
		average_temperature_by_sensor.AverageTemperatureBySensorQuery{},
		average_temperature_by_sensor.Handle,
	)

	dsn := os.Getenv("DB_DNS")
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
	db := bun.NewDB(sqldb, pgdialect.New())
	db.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(true)))

	return di_container.ApiContainer{
		CQRS: cq,
		DB:   db,
	}
}
