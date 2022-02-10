CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE characters (
    "user" BIGINT PRIMARY KEY,
    "id" UUID DEFAULT uuid_generate_v4(),
    "name" VARCHAR(32) UNIQUE NOT NULL,
	"planet" VARCHAR(32) UNIQUE NOT NULL,
    "specks" BIGINT DEFAULT 50
);

CREATE TABLE items (
	"id" BIGSERIAL PRIMARY KEY,
	"owner" BIGINT NOT NULL REFERENCES characters("user") ON DELETE CASCADE,
	"equipped" BOOLEAN NOT NULL,
	"name" VARCHAR(256) NOT NULL,
	-- Value is either the damage or defense for the item
	"value" NUMERIC(5, 2) NOT NULL,
	-- The category can either be shield, saber, helmet, chestplate, gloves, leggings or boots
	"category" VARCHAR(32) NOT NULL,
	-- The slot can be head, torso, legs, feet or equipment
	-- equipment slot is for shields, sabers and the like
	-- head, torso, hands, legs and feet are for their respective coverings
	"slot" VARCHAR(32) NOT NULL,
	-- The rarity can be common, uncommon, rare, legendary and mythic
	"rarity" VARCHAR(32) NOT NULL
);

CREATE INDEX items_owner_idx ON "items" USING btree ("owner");
