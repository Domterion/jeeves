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
			return context.Respond(&discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Flags:   1 << 6,
					Content: "You just got stinkbugged!",
					Components: []discordgo.MessageComponent{
						discordgo.ActionsRow{
							Components: []discordgo.MessageComponent{
								discordgo.Button{
									Emoji: discordgo.ComponentEmoji{
										Name: "🔨",
									},
									Label: "ponk",
									CustomID: context.Member.User.ID + ":" + "ponk",
									Style: discordgo.PrimaryButton,
								},
							},
						},
					},
				},
			})
		},
	},
}
