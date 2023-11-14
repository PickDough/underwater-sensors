package extract_config

import (
	"gopkg.in/yaml.v3"
	"os"
	"underwaterSensors/src/sensor_channels/domain/config"
)

func Handle(query ConfigQuery) (*config.Config, error) {
	var f, err = os.Open(query.File)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var cfg config.Config
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
