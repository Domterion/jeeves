package commands

import (
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/domterion/jeeves/handler"
)

var OwnerTestCommand handler.SubCommandGroup = handler.SubCommandGroup{
	BaseCommand: handler.BaseCommand{
		Name:        "test",
		Description: "...test...",
		Type:        discordgo.ChatApplicationCommand,
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "phrase",
				Description: "The phrase to repeat",
				Required:    true,
			},
			{
				Type:        discordgo.ApplicationCommandOptionChannel,
				Name:        "channel",
				Description: "Yes",
				Required:    false,
			},
		},
		Run: func(context *handler.Context) error {
			log.Printf("owner stats command!")

			return nil
		},
	},

	SubCommands: []*handler.SubCommand{&OwnerTestIclesCommand, &OwnerSayCommand},
}
