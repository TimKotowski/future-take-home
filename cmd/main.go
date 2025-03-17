package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/pgx"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"github.com/TimKotowski/future-take-home/internal/api"
	"github.com/TimKotowski/future-take-home/internal/database"
	"github.com/TimKotowski/future-take-home/internal/routes"
)

func main() {
	done := make(chan struct{}, 1)
	dsn := "postgres://future:12345@pg_future:5432/future?sslmode=disable"
	db, err := database.GetDatabase(dsn)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer db.Close()

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

	m, err := migrate.NewWithDatabaseInstance("file://migrations", "future", driver)
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
	fmt.Println("Running Future Appointments Application")

	//	var pgErr *pgconn.PgError
	//	if errors.As(err, &pgErr) {
	//		fmt.Println("HMMMMMMMM", pgErr.Code, pgErr.SQLState())
	//		log.Fatal(err.Error())
	//	}
	//	return err
	//}

	r := api.NewApi()

	controllers := []routes.RouteRegister{
		routes.NewAppointmentRouteRegister(db),
	}

	for _, route := range controllers {
		route.RegisterRoutes(r)
	}

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal("Unable to listen and serve", err)
	}
	go awaitTerminated(done)
}

func awaitTerminated(done chan struct{}) {
	killSignal := make(chan os.Signal, 1)
	signal.Notify(killSignal, os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT)
	<-killSignal
	done <- struct{}{}
}
