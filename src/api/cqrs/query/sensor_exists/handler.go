package sensor_exists

import (
	"context"
	"underwaterSensors/src/common/domain/sensor/model"
)

func Handle(query SensorExistsQuery) (*bool, error) {
	db := query.Container.DB

	exists, err := db.NewSelect().
		Model((*model.SensorModel)(nil)).
		Join("join sensor_groups sg on sg.id = sensor_group_id").
		Where("sg.group_name = ?", query.Group).
		Where("sensor_index = ?", query.Index).
		Exists(context.Background())
	if err != nil {
		return nil, err
	}

	return &exists, err
}
