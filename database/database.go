package database

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	_ "github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"
)

type Character struct {
	bun.BaseModel `bun:"table:characters"`

	User   string    `bun:"user,type:bigint,pk"`
	ID     uuid.UUID `bun:"id,type:uuid,default:uuid_generate_v4()"`
	Name   string    `bun:"name,type:varchar(32),notnull,unique"`
	Specks int64     `bun:"specks,type:bigint,default:0"`
}

var Db *bun.DB

func Connect(connectionUri string) error {
	db, err := sql.Open("pg", connectionUri)

	if err != nil {
		return err
	}

	Db = bun.NewDB(db, pgdialect.New())

	// TODO: We need to remove this query hook when we are confident
	Db.AddQueryHook(bundebug.NewQueryHook(
		bundebug.WithVerbose(true),
		bundebug.FromEnv("BUNDEBUG"),
	))

	return err
}

func GetCharacter(user string) (Character, error) {
	character := Character{}
	err := Db.NewSelect().Model(&character).Where("\"user\" = ?", user).Scan(context.Background())
	return character, err
}

func InsertCharacter(user string, name string, specks int64) error {
	character := Character{
		User:   user,
		Name:   name,
		Specks: specks,
	}
	_, err := Db.NewInsert().Model(&character).ExcludeColumn("id").Exec(context.Background())
	return err
}
