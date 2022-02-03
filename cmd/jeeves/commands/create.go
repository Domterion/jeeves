package commands

import (
	"database/sql"

	"github.com/bwmarrin/discordgo"
	"github.com/domterion/jeeves/commander"
	"github.com/domterion/jeeves/database"
)

var CreateCommand commander.Command = commander.Command{
	BaseCommand: commander.BaseCommand{
		Name:        "create",
		Description: "Start your space exploration adventure!",
		Type:        discordgo.ChatApplicationCommand,
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "name",
				Description: "The name of your character",
				Required:    true,
			},
		},
		BeforeRun: func(context *commander.CommandContext) bool {
			if _, err := database.GetCharacter(context.Member.User.ID); err != sql.ErrNoRows {
				context.RespondTextEphemeral("You already have a character!")

				return false
			}

			return true
		},
		Run: func(context *commander.CommandContext) error {
			components := commander.Components{}

			actionRow := commander.ActionRow{}
			actionRow.AddButton(commander.Button{
				BaseComponent: commander.BaseComponent{
					CustomID: context.Manager.SnowflakeNode.Generate().String(),
					Run: func(ctx *commander.ComponentContext) error {
						name := context.Event.ApplicationCommandData().Options[0].StringValue()

						err := database.InsertCharacter(context.Member.User.ID, name, 0)

						if err != nil {
							return context.ResponseEdit(&discordgo.WebhookEdit{
								Content:    "Err creating character:\n" + err.Error(),
								Components: []discordgo.MessageComponent{},
							})
						}

						msg := `Your space exploration journey starts here!

I am **Jeeves**, your captain! I'll help guide you through this journey. 

You have **50** SPC, formally known as Specks or the currency, to start.
`

						return context.ResponseEdit(&discordgo.WebhookEdit{
							Content:    msg,
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

			msg := `**Bot Rules**
			
Following bot rules is a must and failure to do so will result in punishment. You are responsible for keeping up with the rules and following them.

**1**) No alting or multi-accounting. You are allowed **ONE** character. 

**2**) No using macros or otherwise that'd give you an unfair advantage. This includes input automation such as automatic typers.

**3**) You may not abuse any exploits or bugs and are to report them immediately.

*Last updated 1/31/2022*
`

			return context.Respond(&discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content:    msg,
					Flags:      1 << 6,
					Components: components.ToMessageComponent(),
				},
			})
		},
	},
}
