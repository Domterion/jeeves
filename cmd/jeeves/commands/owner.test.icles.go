package commands

import (
	"github.com/bwmarrin/discordgo"
	"github.com/domterion/jeeves/handler"
)

var OwnerTestIclesCommand handler.SubCommand = handler.SubCommand{
	BaseCommand: handler.BaseCommand{
		Name:        "icles",
		Description: "icles...",
		Type:        discordgo.ChatApplicationCommand,
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionBoolean,
				Name:        "like",
				Description: "like testicles?",
				Required:    true,
			},
		},
		Run: func(context *handler.Context) error {
			context.RespondText("hehe lewd!")

			return nil
		},
	},
}
