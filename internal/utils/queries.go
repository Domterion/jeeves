package utils

import (
	"context"
	"fmt"

	"github.com/domterion/jeeves/internal/models"
	"github.com/uptrace/bun"
)

// Character Queries

func GetCharacter(db *bun.DB, user string) (models.Character, error) {
	character := models.Character{}
	err := db.NewSelect().Model(&character).Where("\"user\" = ?", user).Scan(context.Background())

	return character, err
}

func InsertCharacter(db *bun.DB, user string, name string, specks int64) error {
	character := models.Character{
		User:   user,
		Name:   name,
		Specks: specks,
	}
	_, err := db.NewInsert().Model(&character).ExcludeColumn("id").Exec(context.Background())

	return err
}

// Item Queries

func InsertItem(db *bun.DB, owner string, equipped bool, name string, value float64, category CategoryType, slot SlotType, rarity RarityType) error {
	item := models.Item{
		Owner:    owner,
		Equipped: equipped,
		Name:     name,
		Value:    value,
		Category: string(category),
		Slot:     string(slot),
		Rarity:   string(rarity),
	}

	res, err := db.NewInsert().Model(&item).ExcludeColumn("id").Returning("*").Exec(context.Background())

	fmt.Printf("res: %v\n", res)

	return err
}
