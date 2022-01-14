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
	},
	SubCommands: []*commander.SubCommand{&UserAvatarCommand},
}
