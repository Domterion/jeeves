package commands

import (
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/domterion/jeeves/commander"
)

var TestCommand commander.Command = commander.Command{
	BaseCommand: commander.BaseCommand{
		Name:        "test",
		Description: "command for... testing",
		Type:        discordgo.ChatApplicationCommand,
		Run: func(context *commander.CommandContext) error {
			components := commander.Components{}
			actionRow := commander.ActionRow{}
			actionRow.AddButton(commander.Button{
				BaseComponent: commander.BaseComponent{
					CustomID: context.Manager.SnowflakeNode.Generate().String(),
					Disabled: true,
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
					CustomID: context.Manager.SnowflakeNode.Generate().String(),
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

			components_ := commander.Components{}
			actionRow_ := commander.ActionRow{}
			actionRow_.AddButton(commander.Button{
				BaseComponent: commander.BaseComponent{
					CustomID: context.Manager.SnowflakeNode.Generate().String(),
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
			components_.AddActionRow(actionRow_)

			context.Respond(&discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content:    "Dont do it...",
					Components: components_.ToMessageComponent(),
				},
			})

			_ = context.DeferResponse()

			time.Sleep(3 * time.Second)

			return context.ResponseEdit(&discordgo.WebhookEdit{
				Content:    "e",
				Components: components.ToMessageComponent(),
			})
		},
	},
}
