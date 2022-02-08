package models

import (
	"github.com/uptrace/bun"
	_ "github.com/uptrace/bun/driver/pgdriver"
)

/*
CREATE TABLE items (
	"id" BIGSERIAL PRIMARY KEY,
	"owner" BIGINT NOT NULL,
	"name" VARCHAR(256) NOT NULL,
	-- Value is either the damage or defense for the item
	"value" NUMERIC(5, 2) NOT NULL,
	-- The category can either be shield, saber, helmet, chestplate, leggings or boots
	"category" VARCHAR(32) NOT NULL,
	-- The slot can be head, torso, legs, feet or equipment
	-- equipment slot is for shields, sabers and the like
	-- head, torso, legs and feet are for their respective coverings
	"slot" VARCHAR(32) NOT NULL,
	-- The rarity can be common, uncommon, rare or mythic
	"rarity" VARCHAR(32) NOT NULL
);
*/

type Item struct {
	bun.BaseModel `bun:"table:characters"`

	ID       int64   `bun:"id,type:bigserial,pk"`
	Owner    string  `bun:"owner,type:bigint"`
	Equipped bool    `bun:"equipped,type:boolean,notnull"`
	Name     string  `bun:"name,type:varchar(256),notnull"`
	Value    float64 `bun:"value,type:numeric(5,2),notnull"`
	Category string  `bun:"category,type:varchar(32),notnull"`
	Slot     string  `bun:"slot,type:varchar(32),notnull"`
	Rarity   string  `bun:"rarity,type:varchar(32),notnull"`
}
