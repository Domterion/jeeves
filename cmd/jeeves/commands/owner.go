package commands

import (
	"github.com/bwmarrin/discordgo"
	"github.com/domterion/jeeves/handler"
)

var OwnerCommand handler.Command = handler.Command{
	BaseCommand: handler.BaseCommand{
		Name:        "owner",
		Description: ".",
		Type:        discordgo.ChatApplicationCommand,
		Options:     []*discordgo.ApplicationCommandOption{},
		// This command cant be called since it has subcommands
		Run: nil,
	},
	SubCommands:      []*handler.SubCommand{&OwnerSayCommand, &OwnerStatsCommand},
	SubCommandGroups: []*handler.SubCommandGroup{&OwnerTestCommand},
}
