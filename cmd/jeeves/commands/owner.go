package commands

import (
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/domterion/jeeves/handler"
)

var OwnerCommand handler.Command = handler.Command{
	BaseCommand: handler.BaseCommand{
		Name:        "owner",
		Description: ".",
		Type:        discordgo.ChatApplicationCommand,
		Options:     []*discordgo.ApplicationCommandOption{},
		Run:         RunOwner,
	},
	SubCommands:      []*handler.SubCommand{&OwnerSayCommand, &OwnerStatsCommand},
	SubCommandGroups: []*handler.SubCommandGroup{&OwnerTestCommand},
}

func RunOwner(context *handler.Context) error {
	log.Printf("owner command!")

	return nil
}
