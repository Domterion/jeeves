package commands

import (
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/domterion/jeeves/handler"
)

var OwnerSayCommand handler.SubCommand = handler.SubCommand{
	BaseCommand: handler.BaseCommand{
		Name:        "say",
		Description: "Make me repeat a phrase",
		Type:        discordgo.ChatApplicationCommand,
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "phrase",
				Description: "The phrase to repeat",
				Required:    true,
			},
		},
		Run: RunOwnerSay,
	},
}

func RunOwnerSay(context *handler.Context) error {
	log.Printf("owner say command!")

	return nil
}