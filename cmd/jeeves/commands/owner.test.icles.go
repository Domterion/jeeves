package commands

import (
	"github.com/bwmarrin/discordgo"
	"github.com/domterion/jeeves/commander"
)

var OwnerTestIclesCommand commander.SubCommand = commander.SubCommand{
	BaseCommand: commander.BaseCommand{
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
		Run: func(context *commander.Context) error {
			context.RespondText("hehe lewd!")

			return nil
		},
	},
}
