package commands

import (
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/domterion/jeeves/handler"
)

type PingCommand struct{}

func (c *PingCommand) Name() string {
	return "ping"
}

func (c *PingCommand) Description() string {
	return "pong!"
}

func (c *PingCommand) SubCommands() *[]handler.SubCommand {
	return &[]handler.SubCommand{}
}

func (c *PingCommand) SubCommandGroups() *[]handler.SubCommandGroup {
	return &[]handler.SubCommandGroup{}
}

func (c *PingCommand) Type() discordgo.ApplicationCommandType {
	return discordgo.ChatApplicationCommand
}

func (c *PingCommand) Options() []*discordgo.ApplicationCommandOption {
	return []*discordgo.ApplicationCommandOption{}
}

func (c *PingCommand) Run(context *handler.Context) error {
	log.Printf("ping command!")

	return nil
}
