package di

import (
	"database/sql"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"
	"os"
	"underwaterSensors/src/common/cqrs"
	"underwaterSensors/src/common/di"
	"underwaterSensors/src/sensor_readers/cqrs/command/save_readings"
	"underwaterSensors/src/sensor_readers/di/di_container"
)

func BuildReaderContainer() di_container.SensorReadersContainer {
	cq := di.BuildCommonContainer().CQRS

	cqrs.RegisterCommand(cq, save_readings.SaveReadingsCommand{}, save_readings.SaveReadings)

	dsn := os.Getenv("DB_DNS")
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
	db := bun.NewDB(sqldb, pgdialect.New())
	db.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(true)))

	return di_container.SensorReadersContainer{
		CQRS: cq,
		DB:   db,
	}
}
