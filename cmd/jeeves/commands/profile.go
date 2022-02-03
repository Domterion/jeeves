package commands

import (
	"database/sql"
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/domterion/jeeves/commander"
	"github.com/domterion/jeeves/database"
)

var ProfileCommand commander.Command = commander.Command{
	BaseCommand: commander.BaseCommand{
		Name:        "profile",
		Description: "View your profile",
		Type:        discordgo.ChatApplicationCommand,
		BeforeRun: func(context *commander.CommandContext) bool {
			if _, err := database.GetCharacter(context.Member.User.ID); err == sql.ErrNoRows {
				context.RespondTextEphemeral("You need a character for this command..")

				return false
			}

			return true
		},
		Run: func(context *commander.CommandContext) error {
			character, _ := database.GetCharacter(context.Member.User.ID)

			description := fmt.Sprintf(`**Name**: %s
**Specks** (**SPC**): %d

**ID**: %s
`, character.Name, character.Specks, character.ID)

			embed := discordgo.MessageEmbed{
				Title:       "Profile",
				Description: description,
			}

			return context.Respond(&discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Embeds: []*discordgo.MessageEmbed{
						&embed,
					},
				},
			})
		},
	},
}
