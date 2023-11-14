package main

import (
	"underwaterSensors/src/api/di"
	"underwaterSensors/src/api/request_handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	diContainer := di.BuildApiContainer()
	r := gin.Default()

	temperatureHandler := request_handlers.NewAverageTemperatureHandler(diContainer)
	r.GET("/group/:groupName/temperature/average", temperatureHandler.Handle)
	speciesHandler := request_handlers.NewSpeciesHandler(diContainer)
	r.GET("/group/:groupName/species", speciesHandler.Handle)
	speciesTopNHandler := request_handlers.NewSpeciesTopNHandler(diContainer)
	r.GET("/group/:groupName/species/top/:N", speciesTopNHandler.Handle)
	averageSensorTemperatureHandler := request_handlers.NewAverageTemperatureBySensorHandler(diContainer)
	r.GET("/sensor/:codename/temperature/average", averageSensorTemperatureHandler.Handle)

	r.Run()
}
