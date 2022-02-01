package database

import (
	"context"

	"github.com/jackc/pgtype/ext/gofrs-uuid"
)

var (
	INSERT_INTO_CHARACTERS = "INSERT INTO characters (\"user\", \"name\") VALUES ($1, $2);"
	SELECT_CHARACTER = "SELECT * FROM characters WHERE \"user\" = $1"
)

func InsertCharacter(user string, name string) error {
	_, err := Pool.Exec(context.Background(), INSERT_INTO_CHARACTERS, user, name)
	return err
}

func SelectCharacter(user string, u *int64, i *uuid.UUID, n *string, s *int64) error {
	// TODO: Figuring out a nice way to serialize the rows to a struct would be good

	err := Pool.QueryRow(context.Background(), SELECT_CHARACTER, user).Scan(u, i, n, s)
	return err
}