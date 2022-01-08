package commands

import (
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/domterion/jeeves/handler"
)

var OwnerStatsCommand handler.SubCommand = handler.SubCommand{
	BaseCommand: handler.BaseCommand{
		Name:        "stats",
		Description: "Get bot stats... or something",
		Type:        discordgo.ChatApplicationCommand,
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "phrase",
				Description: "The phrase to repeat",
				Required:    true,
			},
			{
				Type:        discordgo.ApplicationCommandOptionChannel,
				Name:        "channel",
				Description: "Yes",
				Required:    false,
			},
		},
		Run: RunOwnerStats,
	},
}

func RunOwnerStats(context *handler.Context) error {
	log.Printf("owner stats command!")

	return nil
}
