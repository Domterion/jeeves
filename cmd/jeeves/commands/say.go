package commands

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/domterion/jeeves/commander"
)

var SayCommand commander.Command = commander.Command{
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
		Run: func(context *commander.CommandContext) error {
			message := fmt.Sprintf("%s says: %s", context.Member.User.Mention(), context.Options[0].Value)
			return context.RespondText(message)
		},
	},
}
