package db

import (
	"context"
	"fmt"
	"os"

	"github.com/mbosa/search-pokemon-descriptions/config"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewDbPool() *pgxpool.Pool {
	dbpool, err := pgxpool.New(context.Background(), config.DATABASE_URL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	return dbpool
}
