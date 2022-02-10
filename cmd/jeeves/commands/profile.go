package commands

import (
	"database/sql"
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/domterion/jeeves/internal/utils"
	"github.com/domterion/jeeves/pkg/commander"
	"github.com/uptrace/bun"
)

var ProfileCommand commander.Command = commander.Command{
	BaseCommand: commander.BaseCommand{
		Name:        "profile",
		Description: "View your profile",
		Type:        discordgo.ChatApplicationCommand,
		BeforeRun: func(context *commander.CommandContext) bool {
			database := context.Get(utils.DIDatabase).(*bun.DB)
			if _, err := utils.GetCharacter(database, context.Member.User.ID); err == sql.ErrNoRows {
				context.RespondTextEphemeral("You need a character for this command..")

				return false
			}

			return true
		},
		Run: func(context *commander.CommandContext) error {
			database := context.Get(utils.DIDatabase).(*bun.DB)
			character, _ := utils.GetCharacter(database, context.Member.User.ID)

			// utils.InsertItem(database, context.Member.User.ID, true, utils.GetRandomItemName(utils.GlovesCategory, utils.MythicRarity), 100.0, utils.GlovesCategory, utils.HandsSlot, utils.MythicRarity)

			items, _ := utils.GetEquippedItems(database, context.Member.User.ID)

			var (
				helmet     = "None"
				chestplate = "None"
				leggings   = "None"
				boots      = "None"
				gloves     = "None"
				shield     = "None"
				saber      = "None"
			)

			for _, item := range items {
				switch item.Category {
				case string(utils.HelmetCategory):
					helmet = item.Name
				case string(utils.ChestplateCategory):
					chestplate = item.Name
				case string(utils.LeggingsCategory):
					leggings = item.Name
				case string(utils.BootsCategory):
					boots = item.Name
				case string(utils.GlovesCategory):
					gloves = item.Name
				case string(utils.ShieldCategory):
					shield = item.Name
				case string(utils.SaberCategory):
					saber = item.Name
				}
			}

			description := fmt.Sprintf(`**Name**: %s
**Specks** (**SPC**): %d
**Planet**: %s

**ID**: %s
`, character.Name, character.Specks, character.Planet, character.ID)

			embed := discordgo.MessageEmbed{
				Title:       "Profile",
				Description: description,
			}

			description = fmt.Sprintf(`**Helmet**: %s
**Chestplate**: %s
**Leggings**: %s
**Boots**: %s
**Gloves**: %s

**Shield**: %s
**Saber**: %s
`, helmet, chestplate, leggings, boots, gloves, shield, saber)

			equipped := discordgo.MessageEmbed{
				Title:       "Equipped",
				Description: description,
			}

			return context.Respond(&discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Embeds: []*discordgo.MessageEmbed{
						&embed,
						&equipped,
					},
				},
			})
		},
	},
}
