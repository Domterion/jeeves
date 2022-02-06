package utils

import (
	"context"

	"github.com/domterion/jeeves/internal/models"
	"github.com/uptrace/bun"
)

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