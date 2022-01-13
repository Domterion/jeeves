package checks

import (
	"github.com/bwmarrin/discordgo"
	"github.com/domterion/jeeves/commander"
)

// BeforeRun is always called before the Run function
// If you return true in BeforeRun then Run will be called, otherwise it run
// BeforeRun is inheritied if a subcommand doesnt have its own BeforeRun defined

var OwnerCommand commander.Command = commander.Command{
	BaseCommand: commander.BaseCommand{
		Name:        "owner",
		Description: ".",
		Type:        discordgo.ChatApplicationCommand,
		Options:     []*discordgo.ApplicationCommandOption{},
		BeforeRun: func(context *commander.Context) bool {
			// Check if the command caller ID matches
			return context.Member.User.ID == "300088143422685185"
		},
		// This command cant be called since it has subcommands
		Run: nil,
	},
	SubCommands:      []*commander.SubCommand{&OwnerCoolCommand},
	SubCommandGroups: []*commander.SubCommandGroup{},
}

// The owner cool command will inherit the check from its base command so we dont need to redefine it
var OwnerCoolCommand commander.SubCommand = commander.SubCommand{
	BaseCommand: commander.BaseCommand{
		Name:        "cool",
		Description: "Are you cool?",
		Type:        discordgo.ChatApplicationCommand,
		Options:     []*discordgo.ApplicationCommandOption{},
		BeforeRun:   nil,
		Run: func(context *commander.Context) error {
			context.RespondText("youre so cool!")

			return nil
		},
	},
}
