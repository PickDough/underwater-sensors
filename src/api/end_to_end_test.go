package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"
	"underwaterSensors/src/api/di"
	"underwaterSensors/src/api/request_handlers"
)

func TestAverageTemperature(t *testing.T) {
	ctx := context.Background()

	pgContaienr := setupContainer(ctx)
	defer func() {
		if err := pgContaienr.Terminate(ctx); err != nil {
			panic(err)
		}
	}()

	diContainer := di.BuildApiContainer()
	r := gin.Default()
	temperatureHandler := request_handlers.NewAverageTemperatureHandler(diContainer)
	r.GET("/group/:groupName/temperature/average", temperatureHandler.Handle)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/group/gamma/temperature/average", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.JSONEq(t, `{"temperature_c": 15.0}`, w.Body.String())
}

func TestAverageTemperatureExists(t *testing.T) {
	ctx := context.Background()

	pgContaienr := setupContainer(ctx)
	defer func() {
		if err := pgContaienr.Terminate(ctx); err != nil {
			panic(err)
		}
	}()

	diContainer := di.BuildApiContainer()
	r := gin.Default()
	temperatureHandler := request_handlers.NewAverageTemperatureHandler(diContainer)
	r.GET("/group/:groupName/temperature/average", temperatureHandler.Handle)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/group/JOB_OFFERS/temperature/average", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)
	assert.Equal(t, `group with name "JOB_OFFERS" doesn't exist`, w.Body.String())
}

func setupContainer(ctx context.Context) *postgres.PostgresContainer {
	dbName := "test"
	dbUser := "test"
	dbPassword := "test"

	pgContainer, err := postgres.RunContainer(ctx,
		testcontainers.WithImage("docker.io/postgres:15.2-alpine"),
		postgres.WithInitScripts("db_test/init_test.sql"),
		postgres.WithDatabase(dbName),
		postgres.WithUsername(dbUser),
		postgres.WithPassword(dbPassword),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).
				WithStartupTimeout(5*time.Second),
		),
	)
	if err != nil {
		panic(err)
	}
	conn, _ := pgContainer.ConnectionString(ctx)
	fmt.Printf("CONNECTION: %s\n", conn)

	os.Setenv("DB_DNS", fmt.Sprintf("%ssslmode=disable", conn))
	return pgContainer
}
