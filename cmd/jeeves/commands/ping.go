package commands

import (
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/domterion/jeeves/handler"
)

var PingCommand handler.Command = handler.Command{
	BaseCommand: handler.BaseCommand{
		Name:        "ping",
		Description: ".",
		Type:        discordgo.ChatApplicationCommand,
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionBoolean,
				Name:        "pong",
				Description: "pong?",
				Required:    true,
			},
		},
		Run: func(context *handler.Context) error {
			log.Printf("ping command!")

			return nil
		},
	},
	SubCommands:      []*handler.SubCommand{},
	SubCommandGroups: []*handler.SubCommandGroup{},
}
