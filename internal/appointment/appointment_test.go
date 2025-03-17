package appointment_test

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/pgx"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"

	"github.com/TimKotowski/future-take-home/internal/database"
	"github.com/TimKotowski/future-take-home/internal/entities"
	"github.com/TimKotowski/future-take-home/internal/routes"
	"github.com/TimKotowski/future-take-home/internal/utils"
)

func TestGetAppointmentsByTrainerIdController(t *testing.T) {
	ctx := context.Background()
	pgContainer, err := postgres.Run(ctx,
		"postgres:16-alpine",
		postgres.WithDatabase("future"),
		postgres.WithUsername("user"),
		postgres.WithPassword("12345"),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).WithStartupTimeout(5*time.Second)),
	)
	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() {
		if err := pgContainer.Terminate(ctx); err != nil {
			t.Fatalf("failed to terminate pgContainer: %s", err)
		}
	})

	dsn, err := pgContainer.ConnectionString(ctx, "sslmode=disable")
	assert.NoError(t, err)
	db, err := database.GetDatabase(dsn)
	assert.NoError(t, err)

	_, err = db.Exec("CREATE SCHEMA IF NOT EXISTS future_schema")
	if err != nil {
		log.Fatalf("failed to create schema: %v", err)
	}

	p := &pgx.Config{
		DatabaseName:    "future",
		MigrationsTable: "schema_migrations",
		SchemaName:      "future_schema",
	}
	driver, err := pgx.WithInstance(db.DB, p)
	if err != nil {
		log.Fatalf("Error creating postgres driver: %v", err)
	}

	// For now just hard code path, since not enough time on take home to work on better test set up.
	m, err := migrate.NewWithDatabaseInstance("file://../../migrations", "future", driver)
	if err != nil {
		log.Fatalf("FAILED %v", err)
	}
	if err := m.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			println("no migration changes found")
		} else {
			println("issue with migrations", err.Error())
		}
	}

	r := chi.NewRouter()
	controllers := []routes.RouteRegister{
		routes.NewAppointmentRouteRegister(db),
	}

	for _, route := range controllers {
		route.RegisterRoutes(r)
	}
	app := httptest.NewServer(r)
	t.Cleanup(func() {
		app.Close()
	})

	url := fmt.Sprintf("%s/appointments/v1/1", app.URL)
	res, err := app.Client().Get(url)
	assert.NoError(t, err)

	data, err := io.ReadAll(res.Body)
	assert.NoError(t, err)

	var appointments []entities.Appointment
	err = json.Unmarshal(data, &appointments)
	assert.NoError(t, err)
	// Not good test validation, but don't have much time on test due to time constraints apologies.
	assert.Len(t, appointments, 5)
	for _, appointment := range appointments {
		assert.Equal(t, appointment.TrainerId, 1)
	}
}

func TestCreateAppointmentsController(t *testing.T) {
	ctx := context.Background()
	pgContainer, err := postgres.Run(ctx,
		"postgres:16-alpine",
		postgres.WithDatabase("future"),
		postgres.WithUsername("user"),
		postgres.WithPassword("12345"),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).WithStartupTimeout(5*time.Second)),
	)
	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() {
		if err := pgContainer.Terminate(ctx); err != nil {
			t.Fatalf("failed to terminate pgContainer: %s", err)
		}
	})

	dsn, err := pgContainer.ConnectionString(ctx, "sslmode=disable")
	assert.NoError(t, err)
	db, err := database.GetDatabase(dsn)
	assert.NoError(t, err)

	_, err = db.Exec("CREATE SCHEMA IF NOT EXISTS future_schema")
	if err != nil {
		log.Fatalf("failed to create schema: %v", err)
	}

	p := &pgx.Config{
		DatabaseName:    "future",
		MigrationsTable: "schema_migrations",
		SchemaName:      "future_schema",
	}
	driver, err := pgx.WithInstance(db.DB, p)
	if err != nil {
		log.Fatalf("Error creating postgres driver: %v", err)
	}

	// For now just hard code path, since not enough time on take home to work on better test set up.
	m, err := migrate.NewWithDatabaseInstance("file://../../migrations", "future", driver)
	if err != nil {
		log.Fatalf("FAILED %v", err)
	}
	if err := m.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			println("no migration changes found")
		} else {
			println("issue with migrations", err.Error())
		}
	}

	r := chi.NewRouter()
	controllers := []routes.RouteRegister{
		routes.NewAppointmentRouteRegister(db),
	}

	for _, route := range controllers {
		route.RegisterRoutes(r)
	}
	app := httptest.NewServer(r)
	t.Cleanup(func() {
		app.Close()
	})
	reqBody := struct {
		StartSlot string `json:"start_slot"`
		EndSlot   string `json:"end_slot"`
		Status    string `json:"status"`
	}{
		StartSlot: "2019-01-24T12:00:00-08:00",
		EndSlot:   "2019-01-24T12:30:00-08:00",
		Status:    "ACTIVE",
	}
	assert.NoError(t, err)
	// Marshal the request body to JSON.
	bodyBytes, err := json.Marshal(reqBody)
	if err != nil {
		t.Fatalf("failed to marshal request body: %v", err)
	}
	assert.NoError(t, err)

	bodyReader := bytes.NewReader(bodyBytes)

	trainerId := 1
	userId := 8
	url := fmt.Sprintf("%s/appointments/v1/slots/%d/%d", app.URL, trainerId, userId)
	res, err := app.Client().Post(url, "application/json", bodyReader)
	assert.NoError(t, err)

	data, err := io.ReadAll(res.Body)
	assert.NoError(t, err)

	var appointment entities.Appointment
	err = json.Unmarshal(data, &appointment)
	assert.NoError(t, err)
	expectedStart, expectedEnd, err := utils.ParseStartEndSlots("2019-01-24T12:00:00-08:00", "2019-01-24T12:30:00-08:00")
	assert.Equal(t, appointment.TrainerId, 1)
	assert.Equal(t, appointment.UserId, 8)
	assert.Equal(t, appointment.StartSlot.Time.UTC().String(), expectedStart.UTC().String())
	assert.Equal(t, appointment.EndSlot.Time.UTC().String(), expectedEnd.UTC().String())
}

func TestCreateAppointmentsByTimeRangeController(t *testing.T) {
	ctx := context.Background()
	pgContainer, err := postgres.Run(ctx,
		"postgres:16-alpine",
		postgres.WithDatabase("future"),
		postgres.WithUsername("user"),
		postgres.WithPassword("12345"),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).WithStartupTimeout(5*time.Second)),
	)
	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() {
		if err := pgContainer.Terminate(ctx); err != nil {
			t.Fatalf("failed to terminate pgContainer: %s", err)
		}
	})

	dsn, err := pgContainer.ConnectionString(ctx, "sslmode=disable")
	assert.NoError(t, err)
	db, err := database.GetDatabase(dsn)
	assert.NoError(t, err)

	_, err = db.Exec("CREATE SCHEMA IF NOT EXISTS future_schema")
	if err != nil {
		log.Fatalf("failed to create schema: %v", err)
	}

	p := &pgx.Config{
		DatabaseName:    "future",
		MigrationsTable: "schema_migrations",
		SchemaName:      "future_schema",
	}
	driver, err := pgx.WithInstance(db.DB, p)
	if err != nil {
		log.Fatalf("Error creating postgres driver: %v", err)
	}

	// For now just hard code path, since not enough time on take home to work on better test set up.
	m, err := migrate.NewWithDatabaseInstance("file://../../migrations", "future", driver)
	if err != nil {
		log.Fatalf("FAILED %v", err)
	}
	if err := m.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			println("no migration changes found")
		} else {
			println("issue with migrations", err.Error())
		}
	}

	r := chi.NewRouter()
	controllers := []routes.RouteRegister{
		routes.NewAppointmentRouteRegister(db),
	}

	for _, route := range controllers {
		route.RegisterRoutes(r)
	}
	app := httptest.NewServer(r)
	t.Cleanup(func() {
		app.Close()
	})
	assert.NoError(t, err)

	trainerId := 1
	url := fmt.Sprintf("%s/appointments/v1/slots/%d/%s/%s", app.URL, trainerId, "2019-01-24T09:00:00-08:00", "2019-01-24T09:30:00-08:00")
	res, err := app.Client().Get(url)
	assert.NoError(t, err)

	expectedStart, expectedEnd, err := utils.ParseStartEndSlots("2019-01-24T09:00:00-08:00", "2019-01-24T09:30:00-08:00")
	data, err := io.ReadAll(res.Body)
	assert.NoError(t, err)
	var appointments []entities.Appointment
	err = json.Unmarshal(data, &appointments)
	assert.NoError(t, err)
	assert.Len(t, appointments, 1)
	appointment := appointments[0]
	assert.Equal(t, appointment.TrainerId, 1)
	assert.Equal(t, appointment.StartSlot.Time.UTC().String(), expectedStart.UTC().String())
	assert.Equal(t, appointment.EndSlot.Time.UTC().String(), expectedEnd.UTC().String())
}
