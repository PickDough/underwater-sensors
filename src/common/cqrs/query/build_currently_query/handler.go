package build_currently_query

import "underwaterSensors/src/common/domain/sensor/model"

func Handle(query CurrentlyQuery) (*CurrentlyQueryResult, error) {
	subq := query.DB.NewSelect().
		Model((*model.SensedDataModel)(nil)).
		ColumnExpr("max(sensed_data_model.id)").
		Join("join public.sensors s on s.id = sensed_data_model.sensor_id").
		Join("join public.sensor_groups sg on sg.id = s.sensor_group_id").
		Group("sensor_index")

	return &CurrentlyQueryResult{Query: subq}, nil
}
