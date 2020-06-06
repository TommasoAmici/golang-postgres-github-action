package postgres

import (
	"context"
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

	stmtDropTable := "DROP TABLE users;"
	stmtCreateTable := "CREATE TABLE IF NOT EXISTS users (id SERIAL PRIMARY KEY, username VARCHAR(128));"
	stmtInsert := "INSERT INTO users (username) VALUES ('User name');"
	_, _ = pool.Exec(context.Background(), stmtDropTable)
	_, _ = pool.Exec(context.Background(), stmtCreateTable)
	_, _ = pool.Exec(context.Background(), stmtInsert)

	return pool, func() {
		pool.Close()
	}
}
