package average_temperature

import (
	"context"
	"underwaterSensors/src/common/cqrs"
	"underwaterSensors/src/common/cqrs/query/build_currently_query"
	"underwaterSensors/src/common/domain/sensor/model"
)

func Handle(query AverageTemperatureQuery) (*AverageTemperatureResult, error) {
	db := query.Container.DB

	subq, _ := cqrs.ExecuteQuery[build_currently_query.CurrentlyQueryResult, build_currently_query.CurrentlyQuery](
		query.Container.CQRS,
		build_currently_query.CurrentlyQuery{
			DB: query.Container.DB,
		})

	subq.Query.Where("group_name = ?", query.Group)

	var tempC float32
	err := db.NewSelect().
		Model((*model.SensedDataModel)(nil)).
		Where("id in (?)", subq.Query).
		ColumnExpr("avg(temperature_c)").
		Scan(context.Background(), &tempC)
	if err != nil {
		return nil, err
	}

	return &AverageTemperatureResult{TemperatureC: tempC}, nil
}
