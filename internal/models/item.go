package models

import (
	"github.com/uptrace/bun"
	_ "github.com/uptrace/bun/driver/pgdriver"
)

type Item struct {
	bun.BaseModel `bun:"table:items"`

	ID       int64   `bun:"id,type:bigserial,pk"`
	Owner    string  `bun:"owner,type:bigint"`
	Equipped bool    `bun:"equipped,type:boolean,notnull"`
	Name     string  `bun:"name,type:varchar(256),notnull"`
	Value    float64 `bun:"value,type:numeric(5,2),notnull"`
	Category string  `bun:"category,type:varchar(32),notnull"`
	Slot     string  `bun:"slot,type:varchar(32),notnull"`
	Tier     string  `bun:"tier,type:varchar(32),notnull"`
}
