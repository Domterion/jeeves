package commands

import (
	"github.com/bwmarrin/discordgo"
	"github.com/domterion/jeeves/handler"
)

var OwnerTestCommand handler.SubCommandGroup = handler.SubCommandGroup{
	BaseCommand: handler.BaseCommand{
		Name:        "test",
		Description: ".",
		Type:        discordgo.ChatApplicationCommand,
		Options:     []*discordgo.ApplicationCommandOption{},
		// This also has subcommands so wont ever be called
		Run: nil,
	},

	SubCommands: []*handler.SubCommand{&OwnerTestIclesCommand, &OwnerSayCommand},
}
