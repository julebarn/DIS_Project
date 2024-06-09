package db

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func Conn(ctx context.Context) *pgx.Conn {
	conn, err := pgx.Connect(ctx, "postgres://postgres:admin@localhost:5432/postgres")
	if err != nil {
		panic(err)
	}
	return conn
}
