package commands

import (
	"github.com/bwmarrin/discordgo"
	"github.com/domterion/jeeves/commander"
)

var UserCommand commander.Command = commander.Command{
	BaseCommand: commander.BaseCommand{
		Name:        "user",
		Description: ".",
		Type:        discordgo.ChatApplicationCommand,
		Options:     []*discordgo.ApplicationCommandOption{},
		// BeforeRun if set here will set for all child command
		BeforeRun: nil,
		// This command cant be called since it has subcommands
		Run: nil,
	},
	SubCommands:      []*commander.SubCommand{&UserAvatarCommand},
	SubCommandGroups: []*commander.SubCommandGroup{},
}
