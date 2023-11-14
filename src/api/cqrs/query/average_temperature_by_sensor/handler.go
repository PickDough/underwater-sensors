package average_temperature_by_sensor

import (
	"context"
	"underwaterSensors/src/common/domain/sensor/model"
)

func Handle(query AverageTemperatureBySensorQuery) (*AverageTemperatureBySensorResult, error) {
	db := query.Container.DB

	var tempC float32
	err := db.NewSelect().
		Model((*model.SensedDataModel)(nil)).
		Join("join sensors s on s.id = sensor_id").
		Join("join sensor_groups sg on s.sensor_group_id = sg.id").
		Where("sg.group_name = ?", query.Group).
		Where("s.sensor_index = ?", query.Index).
		Where(
			"created_at between to_timestamp(?) :: timestamp and to_timestamp(?) :: timestamp",
			query.FromDateTime,
			query.UntilDateTime,
		).
		ColumnExpr("avg(temperature_c)").
		Scan(context.Background(), &tempC)
	if err != nil {
		return nil, err
	}

	return &AverageTemperatureBySensorResult{
		TemperatureC: tempC,
	}, nil
}
