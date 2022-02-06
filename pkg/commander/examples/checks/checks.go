package checks

import (
	"github.com/bwmarrin/discordgo"
	"github.com/domterion/jeeves/pkg/commander"
)

// BeforeRun is always called before the Run function
// If you return true in BeforeRun then Run will be called, otherwise it run
// BeforeRun is inheritied if a subcommand doesnt have its own BeforeRun defined

var OwnerCommand commander.Command = commander.Command{
	BaseCommand: commander.BaseCommand{
		Name:        "owner",
		Description: ".",
		Type:        discordgo.ChatApplicationCommand,
		BeforeRun: func(context *commander.CommandContext) bool {
			// Check if the command caller ID matches
			owner := context.Member.User.ID == "300088143422685185"

			if !owner {
				context.RespondText("You need to be the owner for this command.")
			}

			return owner
		},
	},
	SubCommands: []*commander.SubCommand{&OwnerCoolCommand},
}

// The owner cool command will inherit the check from its base command so we dont need to redefine it
var OwnerCoolCommand commander.SubCommand = commander.SubCommand{
	BaseCommand: commander.BaseCommand{
		Name:        "cool",
		Description: "Are you cool?",
		Type:        discordgo.ChatApplicationCommand,
		Run: func(context *commander.CommandContext) error {
			return context.RespondText("Youre so cool!")
		},
	},
}
