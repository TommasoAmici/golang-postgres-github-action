package postgres

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
)

// NewDB opens a postgres connection to dbURL
func NewDB(dbURL *string) (pool *pgxpool.Pool) {
	config, _ := pgxpool.ParseConfig(*dbURL)

	pool, err := pgxpool.ConnectConfig(context.Background(), config)
	if err != nil {
		log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile).Fatal(err)
	}
	return
}
