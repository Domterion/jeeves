package utils

import (
	"context"

	"github.com/domterion/jeeves/internal/models"
	"github.com/uptrace/bun"
)

// Character Queries

func GetCharacter(db *bun.DB, user string) (models.Character, error) {
	var character models.Character
	err := db.NewSelect().Model(&character).Where("\"user\" = ?", user).Scan(context.Background())

	return character, err
}

func InsertCharacter(db *bun.DB, user string, name string, specks int64, planet PlanetType) error {
	character := models.Character{
		User:   user,
		Name:   name,
		Specks: specks,
		Planet: string(planet),
	}
	_, err := db.NewInsert().Model(&character).ExcludeColumn("id").Exec(context.Background())

	return err
}

// Item Queries

func InsertItem(db *bun.DB, owner string, equipped bool, name string, value float64, category CategoryType, slot SlotType, rarity RarityType) (*models.Item, error) {
	item := models.Item{
		Owner:    owner,
		Equipped: equipped,
		Name:     name,
		Value:    value,
		Category: string(category),
		Slot:     string(slot),
		Rarity:   string(rarity),
	}

	var returned models.Item
	_, err := db.NewInsert().Model(&item).ExcludeColumn("id").Returning("*").Exec(context.Background(), &returned)

	return &returned, err
}

func GetEquippedItems(db *bun.DB, owner string) ([]models.Item, error) {
	var items []models.Item

	err := db.NewSelect().Model(&items).Where("\"owner\" = ?", owner).Where("\"equipped\" = true").Scan(context.Background())

	return items, err
}
