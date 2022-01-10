package commands

import (
	"github.com/bwmarrin/discordgo"
	"github.com/domterion/jeeves/commander"
)

var OwnerSayCommand commander.SubCommand = commander.SubCommand{
	BaseCommand: commander.BaseCommand{
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
		Run: func(context *commander.Context) error {
			context.RespondText("template")

			return nil
		},
	},
}
