package database

import (
	"context"

	"github.com/jackc/pgx/v4"
)

func Connect(connectionUri string) (*pgx.Conn, error) {
	conn, err := pgx.Connect(context.Background(), connectionUri)

	return conn, err
}