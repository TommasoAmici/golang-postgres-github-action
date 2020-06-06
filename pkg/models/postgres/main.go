package postgres

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
)

// NewDB opens a postgres connection to dbURL
func NewDB(dbURL *string) (pool *pgxpool.Pool) {
	config, _ := pgxpool.ParseConfig(*dbURL)

	// in a real application you'll want to handle errors
	pool, _ = pgxpool.ConnectConfig(context.Background(), config)
	return
}
