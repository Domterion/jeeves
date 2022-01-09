package commands

import (
	"github.com/bwmarrin/discordgo"
	"github.com/domterion/jeeves/handler"
)

var PingCommand handler.Command = handler.Command{
	BaseCommand: handler.BaseCommand{
		Name:        "ping",
		Description: "pong",
		Type:        discordgo.ChatApplicationCommand,
		Options: []*discordgo.ApplicationCommandOption{},
		Run: func(context *handler.Context) error {
			context.RespondText("pong!")

			return nil
		},
	},
	SubCommands:      []*handler.SubCommand{},
	SubCommandGroups: []*handler.SubCommandGroup{},
}
