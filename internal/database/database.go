package database

import (
	"database/sql"
	"fmt"
	"log"
	"time"

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

	if err := healthCheck(db.DB, 10); err != nil {
		return nil, err
	}

	return db, nil
}

// healthCheck ensures that the database is up and running before allowing application to continue.
func healthCheck(db *sql.DB, maxRetries int) error {
	for i := 0; i < maxRetries; i++ {
		if err := db.Ping(); err == nil {
			return nil
		}
		time.Sleep(5 * time.Second)
	}
	return fmt.Errorf("postgres not ready after %d retries", maxRetries)
}
