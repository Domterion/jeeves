package commands

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/domterion/jeeves/commander"
)

var UserAvatarCommand commander.SubCommand = commander.SubCommand{
	BaseCommand: commander.BaseCommand{
		Name:        "avatar",
		Description: "Get the avatar of a user",
		Type:        discordgo.ChatApplicationCommand,
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionUser,
				Name:        "user",
				Description: "The users avatar to get",
				Required:    true,
			},
		},
		BeforeRun: nil,
		Run: func(context *commander.Context) error {
			// userId := context.Options[0].Value
			// log.Println(userId)
			// log.Println(context.Event.ApplicationCommandData().Options[0].Options[0].UserValue(nil).AvatarURL(""))

			for _, user := range context.ResolvedOptions.Users {
				message := fmt.Sprintf("%s's avatar:\n\n%s", user.Mention(), user.AvatarURL(""))
				context.RespondText(message)
			}
			SHEnTiveriSERmONsfIatEPicIa
			return nil
		},
	},
}
