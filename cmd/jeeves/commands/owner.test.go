package commands

import (
	"github.com/bwmarrin/discordgo"
	"github.com/domterion/jeeves/commander"
)

var OwnerTestCommand commander.SubCommandGroup = commander.SubCommandGroup{
	BaseCommand: commander.BaseCommand{
		Name:        "test",
		Description: ".",
		Type:        discordgo.ChatApplicationCommand,
		Options:     []*discordgo.ApplicationCommandOption{},
		// This also has subcommands so wont ever be called
		Run: nil,
	},

	SubCommands: []*commander.SubCommand{&OwnerTestIclesCommand},
}
