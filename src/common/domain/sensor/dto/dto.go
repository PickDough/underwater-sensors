package dto

type SensedData struct {
	Sensor       *Sensor       `json:"sensor"`
	TemperatureC float32       `json:"temperature_C"`
	Fishes       []*SensedFish `json:"fishes"`
}

type Sensor struct {
	Group       string `yaml:"group"`
	Coordinates struct {
		X float32 `yaml:"x"`
		Y float32 `yaml:"y"`
		Z float32 `yaml:"z"`
	} `yaml:"coordinates"`
	Index int
	Rate  int `yaml:"rate"`
}

type Fish struct {
	Name string `yaml:"name"`
}

type SensedFish struct {
	Name  string
	Count int
}
