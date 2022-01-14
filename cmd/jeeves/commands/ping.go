package commands

import (
	"github.com/bwmarrin/discordgo"
	"github.com/domterion/jeeves/commander"
)

var PingCommand commander.Command = commander.Command{
	BaseCommand: commander.BaseCommand{
		Name:        "ping",
		Description: "pong",
		Type:        discordgo.ChatApplicationCommand,
		Run: func(context *commander.Context) error {
			return context.RespondText("pong!")
		},
	},
}
