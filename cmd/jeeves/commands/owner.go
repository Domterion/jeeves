package commands

import (
	"github.com/bwmarrin/discordgo"
	"github.com/domterion/jeeves/commander"
)

var OwnerCommand commander.Command = commander.Command{
	BaseCommand: commander.BaseCommand{
		Name:        "owner",
		Description: ".",
		Type:        discordgo.ChatApplicationCommand,
		Options:     []*discordgo.ApplicationCommandOption{},
		// This command cant be called since it has subcommands
		Run: nil,
	},
	SubCommands:      []*commander.SubCommand{&OwnerStatsCommand},
	SubCommandGroups: []*commander.SubCommandGroup{&OwnerTestCommand},
}
