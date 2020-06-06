package postgres

import (
	"testing"

	"github.com/jackc/pgx/v4/pgxpool"
)

// newTestDB creates a db connection to the test database
func newTestDB(t *testing.T) (*pgxpool.Pool, func()) {
	if testing.Short() {
		t.Skip("postgres: skipping integration test")
	}

	// in a real application you'll want to pass the db url as a flag or environment variable
	dbURL := "postgresql://test_golang:test_golang@localhost:5432/test_golang"
	pool := NewDB(&dbURL)
	return pool, func() {
		pool.Close()
	}
}
