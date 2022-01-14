package commands

import (
	"github.com/bwmarrin/discordgo"
	"github.com/domterion/jeeves/commander"
)

var TestCommand commander.Command = commander.Command{
	BaseCommand: commander.BaseCommand{
		Name:        "test",
		Description: "command for... testing",
		Type:        discordgo.ChatApplicationCommand,
		Options:     []*discordgo.ApplicationCommandOption{},
		BeforeRun:   nil,
		Run: func(context *commander.Context) error {
			return context.Respond(&discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Flags:   1 << 6,
					Content: "butt",
					Components: []discordgo.MessageComponent{
						discordgo.ActionsRow{
							Components: []discordgo.MessageComponent{
								discordgo.Button{
									Emoji: discordgo.ComponentEmoji{
										Name: "🚀",
									},
									Label:    "wocket",
									CustomID: context.Member.User.ID + ":" + "test",
									Style:    discordgo.PrimaryButton,
								},
							},
						},
					},
				},
			})
		},
	},
	SubCommands:      []*commander.SubCommand{},
	SubCommandGroups: []*commander.SubCommandGroup{},
}
