package faker

import (
	"math/rand"
	"time"
	"underwaterSensors/src/common/domain/sensor/dto"
)

type SensorFaker struct {
	sensor  *dto.Sensor
	rnd     *rand.Rand
	fishes  []*dto.Fish
	Channel chan dto.SensedData
}

func New(sensor *dto.Sensor, fishes []*dto.Fish) *SensorFaker {
	return &SensorFaker{
		sensor:  sensor,
		rnd:     rand.New(rand.NewSource(time.Now().Unix())),
		fishes:  fishes,
		Channel: make(chan dto.SensedData),
	}
}

func (f SensorFaker) Sense() {
	for {
		f.Channel <- dto.SensedData{
			TemperatureC: f.senseTemperatureC(),
			Sensor:       f.sensor,
			Fishes:       f.senseFishes(),
		}
		time.Sleep(time.Second * time.Duration(f.sensor.Rate))
	}
}

func (f SensorFaker) senseTemperatureC() float32 {
	return 20.0 - (f.sensor.Coordinates.Z / float32(f.rnd.Intn(25)+20))
}

func (f SensorFaker) senseFishes() []*dto.SensedFish {
	sensed := make(map[string]int)

	arrSize := f.rnd.Intn(8)
	for i := 0; i < arrSize; i++ {
		fish := f.fishes[f.rnd.Intn(len(f.fishes))]
		sensed[fish.Name] = f.rnd.Intn(15) + 1
	}

	lst := make([]*dto.SensedFish, 0)
	for s, i := range sensed {
		lst = append(lst, &dto.SensedFish{
			Name:  s,
			Count: i,
		})
	}

	return lst
}
