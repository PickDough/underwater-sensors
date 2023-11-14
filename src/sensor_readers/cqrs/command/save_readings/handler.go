package save_readings

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/uptrace/bun"
	"underwaterSensors/src/common/domain/sensor/model"
)

func SaveReadings(command SaveReadingsCommand) error {
	db := command.Container.DB

	groupModel := model.GroupModel{
		Name: command.Readings.Sensor.Group,
	}
	ctx := context.Background()
	_, err := db.NewInsert().
		Model(&groupModel).
		On("conflict (group_name) do update").
		Exec(ctx)
	if err != nil {
		return err
	}

	fmt.Printf("GroupId: %d \n", groupModel.ID)
	sensorModel := model.SensorModel{
		BaseModel:     bun.BaseModel{},
		Index:         int64(command.Readings.Sensor.Index),
		SensorGroupId: groupModel.ID,
	}
	_, err = db.NewInsert().
		Model(&sensorModel).
		On("conflict (sensor_group_id, sensor_index) do update").
		Exec(ctx)
	if err != nil {
		return err
	}

	tx, err := db.BeginTx(ctx, &sql.TxOptions{})

	sensedDataModel := model.SensedDataModel{
		TemperatureC: command.Readings.TemperatureC,
		SensorId:     sensorModel.ID,
	}
	_, err = tx.NewInsert().Model(&sensedDataModel).Ignore().Exec(ctx)
	if err != nil {
		tx.Rollback()
		return err
	}

	sensedFishesModel := make([]model.SensedFishesModel, 0)
	for _, fish := range command.Readings.Fishes {
		sensedFishesModel = append(sensedFishesModel, model.SensedFishesModel{
			SensorReadingsID: sensedDataModel.ID,
			Fish:             fish.Name,
			Count:            fish.Count,
		})
	}
	_, err = tx.NewInsert().Model(&sensedFishesModel).Exec(ctx)
	if err != nil {
		tx.Rollback()
		return err
	}
	err = tx.Commit()
	if err != nil {
		return err
	}

	return err
}
