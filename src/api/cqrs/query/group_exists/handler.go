package group_exists

import (
	"context"
	"underwaterSensors/src/common/domain/sensor/model"
)

func Handle(query GroupExistsQuery) (*bool, error) {
	db := query.Container.DB

	exists, err := db.NewSelect().Model((*model.GroupModel)(nil)).Where("group_name = ?", query.Group).Exists(context.Background())
	if err != nil {
		return nil, err
	}

	return &exists, err
}
