package model

import "github.com/uptrace/bun"

type SensedDataModel struct {
	bun.BaseModel `bun:"table:sensor_readings"`
	ID            int64                `bun:"id,pk,autoincrement"`
	TemperatureC  float32              `bun:"temperature_c"`
	SensorId      int64                `bun:"sensor_id"`
	FishesModel   []*SensedFishesModel `bun:"rel:has-many,join:id=sensor_readings_id"`
}

type SensedFishesModel struct {
	bun.BaseModel    `bun:"table:fish_readings"`
	ID               int64  `bun:"id,pk,autoincrement"`
	SensorReadingsID int64  `bun:"sensor_readings_id"`
	Fish             string `bun:"fish"`
	Count            int    `bun:"count"`
}

type SensorModel struct {
	bun.BaseModel `bun:"table:sensors"`
	ID            int64              `bun:"id,pk,autoincrement"`
	Index         int64              `bun:"sensor_index"`
	SensorGroupId int64              `bun:"sensor_group_id"`
	Readings      []*SensedDataModel `bun:"rel:has-many,join:id=sensor_id"`
}

type GroupModel struct {
	bun.BaseModel `bun:"table:sensor_groups"`
	ID            int64          `bun:"id,pk,autoincrement"`
	Name          string         `bun:"group_name"`
	SensorModel   []*SensorModel `bun:"rel:has-many,join:id=sensor_group_id"`
}
