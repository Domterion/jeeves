package commands

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/domterion/jeeves/commander"
)

var UserInfoCommand commander.Command = commander.Command{
	BaseCommand: commander.BaseCommand{
		Name: "userinfo",
		Type: discordgo.UserApplicationCommand,
		Run: func(context *commander.CommandContext) error {
			return context.Respond(&discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Flags:   1 << 6,
					Content: fmt.Sprintf("User ID: %v", context.Member.User.ID),
				},
			})
		},
	},
}
