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
		Options:     []*discordgo.ApplicationCommandOption{},
		BeforeRun: nil,
		Run: func(context *commander.Context) error {
			context.RespondText("pong!")

			return nil
		},
	},
	SubCommands:      []*commander.SubCommand{},
	SubCommandGroups: []*commander.SubCommandGroup{},
}
