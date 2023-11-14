package species

import (
	"context"
	"underwaterSensors/src/common/cqrs"
	"underwaterSensors/src/common/cqrs/query/build_currently_query"
	"underwaterSensors/src/common/domain/sensor/dto"
	"underwaterSensors/src/common/domain/sensor/model"
)

func Handle(query SpeciesQuery) (*SpeciesResult, error) {
	db := query.Container.DB

	subq, _ := cqrs.ExecuteQuery[build_currently_query.CurrentlyQueryResult, build_currently_query.CurrentlyQuery](
		query.Container.CQRS,
		build_currently_query.CurrentlyQuery{
			DB: db,
		})

	subq.Query.Where("group_name = ?", query.Group)

	sensedFishes := make([]map[string]any, 0)
	q := db.NewSelect().
		Model((*model.SensedDataModel)(nil)).
		Where("sensed_data_model.id in (?)", subq.Query).
		Join("join fish_readings fr on fr.sensor_readings_id = sensed_data_model.id").
		ColumnExpr("fr.fish, sum(count)").
		Group("fr.fish")

	if query.Limit > 0 {
		q.Order("sum desc").
			Limit(query.Limit)
	}
	err := q.Scan(context.Background(), &sensedFishes)
	if err != nil {
		return nil, err
	}
	sensedFishesDto := make([]*dto.SensedFish, 0)
	for _, m := range sensedFishes {
		sensedFishesDto = append(sensedFishesDto, &dto.SensedFish{
			Name:  m["fish"].(string),
			Count: int(m["sum"].(int64)),
		})
	}

	return &SpeciesResult{SensedFishes: sensedFishesDto}, nil
}
