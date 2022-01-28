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
		Run: func(context *commander.Context) error {
			components := commander.Components{}

			actionRow := commander.ActionRow{}
			actionRow.AddButton(commander.Button{
				BaseComponent: commander.BaseComponent{
					CustomID: context.Event.GuildID + ":" + context.Member.User.ID + ":" + "test:button",
					Run: func(ctx *commander.ComponentContext) error {
						return ctx.RespondText("pwessed! uwu")
					},
				},
				Emoji: &discordgo.ComponentEmoji{
					Name: "🚀",
				},
				Label: "wocket",
				Style: discordgo.PrimaryButton,
			})
			components.AddActionRow(actionRow)

			actionRow = commander.ActionRow{}
			actionRow.AddSelectMenu(commander.SelectMenu{
				BaseComponent: commander.BaseComponent{
					CustomID: context.Event.GuildID + ":" + context.Member.User.ID + ":" + "test:select",
					Run: func(ctx *commander.ComponentContext) error {
						return ctx.RespondText("selected")
					},
				},
				Options: &[]discordgo.SelectMenuOption{
					{
						Label:       "Justin",
						Value:       "is gay",
						Description: "truth~",
					},
				},
			})
			components.AddActionRow(actionRow)

			context.Manager.AddComponents(components)

			return context.Respond(&discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Flags:      1 << 6,
					Content:    "butt",
					Components: components.ToMessageComponent(),
				},
			})
		},
	},
}
