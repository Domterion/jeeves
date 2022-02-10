package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
	_ "github.com/uptrace/bun/driver/pgdriver"
)

type Character struct {
	bun.BaseModel `bun:"table:characters"`

	User   string    `bun:"user,type:bigint,pk"`
	ID     uuid.UUID `bun:"id,type:uuid,default:uuid_generate_v4()"`
	Name   string    `bun:"name,type:varchar(32),notnull,unique"`
	Specks int64     `bun:"specks,type:bigint,default:0"`
	Planet string    `bun:"planet,type:varchar(32),notnull"`
}
