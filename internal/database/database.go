package database

import (
	"log"

	"github.com/jmoiron/sqlx"
)

const (
	ExclusionViolation = "23P01"
)

func GetDatabase(dsn string) (*sqlx.DB, error) {
	// Just use static values for testing, no need to add configs in this test for right DSN, etc.
	db, err := sqlx.Open("pgx", dsn)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
		return nil, err
	}

	return db, nil
}
