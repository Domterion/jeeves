package commands

import (
	"github.com/bwmarrin/discordgo"
	"github.com/domterion/jeeves/pkg/commander"
)

var PingCommand commander.Command = commander.Command{
	BaseCommand: commander.BaseCommand{
		Name:        "ping",
		Description: "pong",
		Type:        discordgo.ChatApplicationCommand,
		Run: func(context *commander.CommandContext) error {
			return context.RespondText("pong!")
		},
	},
}
