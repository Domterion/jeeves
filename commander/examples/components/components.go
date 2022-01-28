package checks

import (
	"github.com/bwmarrin/discordgo"
	"github.com/domterion/jeeves/commander"
)

var PressCommand commander.Command = commander.Command{
	BaseCommand: commander.BaseCommand{
		Name:        "press",
		Description: "Dont press the button",
		Type:        discordgo.ChatApplicationCommand,
		Run: func(context *commander.Context) error {
			components := commander.Components{}

			actionRow := commander.ActionRow{}
			actionRow.AddButton(commander.Button{
				BaseComponent: commander.BaseComponent{
					CustomID: context.Event.GuildID + ":" + context.Member.User.ID + ":" + "press:button",
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
