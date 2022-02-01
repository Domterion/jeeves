package database

import (
	"context"
	"errors"

	"github.com/jackc/pgtype"
	pgtypeuuid "github.com/jackc/pgtype/ext/gofrs-uuid"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

var Pool *pgxpool.Pool

func Connect(connectionUri string) error {
	err := errors.New("this is needed so err is already defined, yay")
	Pool, err = pgxpool.Connect(context.Background(), connectionUri)

	Pool.Config().AfterConnect = func(ctx context.Context, conn *pgx.Conn) error {
		conn.ConnInfo().RegisterDataType(pgtype.DataType{
			Value: &pgtypeuuid.UUID{},
			Name:  "uuid",
			OID:   pgtype.UUIDOID,
		})
		return nil
	}

	return err
}
