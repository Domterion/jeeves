package commands

import (
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/domterion/jeeves/handler"
)

type OwnerCommand struct{}

func (c *OwnerCommand) Name() string {
	return "owner"
}

func (c *OwnerCommand) Description() string {
	return "."
}

func (c *OwnerCommand) SubCommands() *[]handler.SubCommand {
	return &[]handler.SubCommand{&OwnerSayCommand{}}
}

func (c *OwnerCommand) SubCommandGroups() *[]handler.SubCommandGroup {
	return &[]handler.SubCommandGroup{}
}

func (c *OwnerCommand) Type() discordgo.ApplicationCommandType {
	return discordgo.ChatApplicationCommand
}

func (c *OwnerCommand) Options() []*discordgo.ApplicationCommandOption {
	return []*discordgo.ApplicationCommandOption{}
}

func (c *OwnerCommand) Run(context *handler.Context) error {
	log.Printf("owner command!")

	return nil
}
