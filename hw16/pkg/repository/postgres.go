package repository

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
)

func NewPostgresDB(ctx context.Context, connString string) (*pgxpool.Pool, error) {
	db, err := pgxpool.Connect(ctx, connString)
	if err != nil {
		return nil, err
	}

	return db, db.Ping(ctx)
}
