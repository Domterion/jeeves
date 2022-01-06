package commands

import (
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/domterion/jeeves/handler"
)

var UserInfoCommand handler.Command = handler.Command{
	BaseCommand: handler.BaseCommand{
		Name:    "userinfo",
		Type:    discordgo.UserApplicationCommand,
		Options: []*discordgo.ApplicationCommandOption{},
		Run:     RunUserInfo,
	},
	SubCommands:      []*handler.SubCommand{&OwnerSayCommand},
	SubCommandGroups: []*handler.SubCommandGroup{},
}

func RunUserInfo(context *handler.Context) error {
	log.Printf("userinfo command!")

	return nil
}
