package commands

import (
	"github.com/bwmarrin/discordgo"
	"github.com/domterion/jeeves/commander"
	"github.com/domterion/jeeves/database"
	"github.com/jackc/pgtype/ext/gofrs-uuid"
	"github.com/jackc/pgx/v4"
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
		Run: func(context *commander.CommandContext) error {
			var (
				u int64
				i uuid.UUID
				n string
				s int64
			)

			err := database.SelectCharacter(context.Member.User.ID, &u, &i, &n, &s)

			if err != pgx.ErrNoRows {
				return context.RespondText("You already have a character!")
			}

			components := commander.Components{}

			actionRow := commander.ActionRow{}
			actionRow.AddButton(commander.Button{
				BaseComponent: commander.BaseComponent{
					CustomID: context.Manager.SnowflakeNode.Generate().String(),
					Run: func(ctx *commander.ComponentContext) error {
						name := context.Event.ApplicationCommandData().Options[0].StringValue()
						err := database.InsertCharacter(ctx.Member.User.ID, name)

						if err != nil {
							return context.ResponseEdit(&discordgo.WebhookEdit{
								Content:    "Err creating character:\n" + err.Error(),
								Components: []discordgo.MessageComponent{},
							})
						}

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
