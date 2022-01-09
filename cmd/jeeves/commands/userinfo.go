package commands

import (
	"github.com/bwmarrin/discordgo"
	"github.com/domterion/jeeves/handler"
)

var UserInfoCommand handler.Command = handler.Command{
	BaseCommand: handler.BaseCommand{
		Name:    "userinfo",
		Type:    discordgo.UserApplicationCommand,
		Options: []*discordgo.ApplicationCommandOption{},
		Run: func(context *handler.Context) error {
			context.Respond(&discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Flags:   1 << 6,
					Content: "You just got stinkbugged!",
				},
			})

			return nil
		},
	},
	SubCommands:      []*handler.SubCommand{},
	SubCommandGroups: []*handler.SubCommandGroup{},
}
