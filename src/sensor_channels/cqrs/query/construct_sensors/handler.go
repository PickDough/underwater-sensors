package construct_sensors

import (
	"underwaterSensors/src/sensor_channels/domain/faker"
)

func ConstructSensors(query ConstructQuery) (*ConstructResult, error) {
	lst := make([]*faker.SensorFaker, 0)

	groupMap := make(map[string]int)

	for _, s := range query.Config.Sensors {
		_, ok := groupMap[s.Group]
		if ok {
			groupMap[s.Group] += 1
		} else {
			groupMap[s.Group] = 1
		}
		s.Index = groupMap[s.Group]
		lst = append(lst, faker.New(s, query.Config.Fishes[:]))
	}

	return &ConstructResult{Sensors: lst}, nil
}
