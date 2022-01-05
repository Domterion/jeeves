package commands

import (
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/domterion/jeeves/handler"
)

type OwnerSayCommand struct{}

func (c *OwnerSayCommand) Name() string {
	return "owner"
}

func (c *OwnerSayCommand) Description() string {
	return "Make the bot repeat a phrase"
}

func (c *OwnerSayCommand) Type() discordgo.ApplicationCommandType {
	return discordgo.ChatApplicationCommand
}

func (c *OwnerSayCommand) Options() []*discordgo.ApplicationCommandOption {
	return []*discordgo.ApplicationCommandOption{
		{
			Type:        discordgo.ApplicationCommandOptionString,
			Name:        "phrase",
			Description: "The phrase to repeat",
			Required:    true,
		},
	}
}

func (c *OwnerSayCommand) Run(context *handler.Context) error {
	log.Printf("owner say command!")

	return nil
}
