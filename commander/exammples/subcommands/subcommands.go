package subcommands

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/domterion/jeeves/commander"
)

var UserCommand commander.Command = commander.Command{
	BaseCommand: commander.BaseCommand{
		Name:        "user",
		Description: ".",
		Type:        discordgo.ChatApplicationCommand,
		Options:     []*discordgo.ApplicationCommandOption{},
		BeforeRun:   nil,
		// This command cant be called since it has subcommands
		Run: nil,
	},
	SubCommands:      []*commander.SubCommand{&UserAvatarCommand},
	SubCommandGroups: []*commander.SubCommandGroup{},
}

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
			for _, user := range context.ResolvedOptions.Users {
				message := fmt.Sprintf("%s's avatar:\n\n%s", user.Mention(), user.AvatarURL(""))
				context.RespondText(message)
			}

			return nil
		},
	},
}
