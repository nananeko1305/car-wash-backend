package database

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func PGClient(ctx context.Context) (*pgxpool.Pool, error) {
	client, err := pgxpool.New(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		return nil, err
	}

	return client, nil
}
