package utils

import (
	"context"

	"github.com/domterion/jeeves/internal/models"
	"github.com/uptrace/bun"
)

/*
	database := context.Get("database").(*bun.DB)

	if err := database.NewSelect().Model(&models.Character{}).Where("\"user\" = ?", context.Member.User.ID).Scan(context_.Background()); err != sql.ErrNoRows {
		context.RespondTextEphemeral("You already have a character!")

		return false
	}
*/

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