package checks

import (
	"github.com/bwmarrin/discordgo"
	"github.com/domterion/jeeves/pkg/commander"
)

var PressCommand commander.Command = commander.Command{
	BaseCommand: commander.BaseCommand{
		Name:        "press",
		Description: "Dont press the button",
		Type:        discordgo.ChatApplicationCommand,
		Run: func(context *commander.CommandContext) error {
			components := commander.Components{}

			actionRow := commander.ActionRow{}
			actionRow.AddButton(commander.Button{
				BaseComponent: commander.BaseComponent{
					// We provide a snowflake generator that you can use for custom IDs
					CustomID: context.Manager.SnowflakeNode.Generate().String(),
					Run: func(ctx *commander.ComponentContext) error {
						return ctx.RespondText("You were told not to..")
					},
				},
				Emoji: &discordgo.ComponentEmoji{
					Name: "👀",
				},
				Label: "Press me!",
				Style: discordgo.PrimaryButton,
			})
			components.AddActionRow(actionRow)

			context.Manager.AddComponents(components)

			return context.Respond(&discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content:    "Dont do it...",
					Components: components.ToMessageComponent(),
				},
			})
		},
	},
}
