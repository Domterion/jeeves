package commands

import (
	"github.com/bwmarrin/discordgo"
	"github.com/domterion/jeeves/commander"
)

var CreateCommand commander.Command = commander.Command{
	BaseCommand: commander.BaseCommand{
		Name:        "create",
		Description: "Start your space exploration adventure!",
		Type:        discordgo.ChatApplicationCommand,
		BeforeRun: func(context *commander.CommandContext) bool {
			// TODO: Once we get this database connected we will do the character check here
			return true
		},
		Run: func(context *commander.CommandContext) error {
			components := commander.Components{}

			actionRow := commander.ActionRow{}
			actionRow.AddButton(commander.Button{
				BaseComponent: commander.BaseComponent{
					CustomID: context.Manager.SnowflakeNode.Generate().String(),
					Run: func(ctx *commander.ComponentContext) error {
						return context.ResponseEdit(&discordgo.WebhookEdit{
							Content:    "Creating character!",
							Components: []discordgo.MessageComponent{},
						})
					},
				},
				Emoji: &discordgo.ComponentEmoji{
					Name: "greencheck",
					ID:   "758380151544217670",
				},
				Label: "Accept",
				Style: discordgo.PrimaryButton,
			})
			actionRow.AddButton(commander.Button{
				BaseComponent: commander.BaseComponent{
					CustomID: context.Manager.SnowflakeNode.Generate().String(),
					Run: func(ctx *commander.ComponentContext) error {
						return context.ResponseEdit(&discordgo.WebhookEdit{
							Content:    "You must accept the rules to continue, aborting.",
							Components: []discordgo.MessageComponent{},
						})
					},
				},
				Emoji: &discordgo.ComponentEmoji{
					Name: "redcross",
					ID:   "758380151238033419",
				},
				Label: "Deny",
				Style: discordgo.DangerButton,
			})

			components.AddActionRow(actionRow)
			context.Manager.AddComponents(components)

			rules := `**Bot Rules**
Following bot rules is a must and failure to do so will result in punishment. You are responsible for keeping up with the rules and following them.

**1**) No alting or multi-accounting. You are allowed **ONE** character. 

**2**) No using macros or otherwise that'd give you an unfair advantage. This includes input automation such as automatic typers.

**3**) You may not abuse any exploits or bugs and are to report them immediately.

*Last updated 1/31/2022*
`

			return context.Respond(&discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content:    rules,
					Flags:      1 << 6,
					Components: components.ToMessageComponent(),
				},
			})
		},
	},
}
