package commands

import (
	"github.com/bwmarrin/discordgo"
	"github.com/domterion/jeeves/commander"
)

var OwnerStatsCommand commander.SubCommand = commander.SubCommand{
	BaseCommand: commander.BaseCommand{
		Name:        "stats",
		Description: "Get bot stats... or something",
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
		Run: func(context *commander.Context) error {
			context.RespondText("yes")

			return nil
		},
	},
}
