package handler

import (
	"errors"
	"log"

	"github.com/bwmarrin/discordgo"
)

/*
Base functions for slash commands
*/
type BaseCommand interface {
	Name() string                                   // A function to get the name of the command to register with Discord
	Description() string                            // A function to get the description of the slash command
	Type() discordgo.ApplicationCommandType         // A function to get the type of the command
	Options() []*discordgo.ApplicationCommandOption // A function to get all the command options

	Run(context *Context) error // The handler function for the command
}

/*
A slash command
*/
type Command interface {
	BaseCommand

	SubCommands() *[]SubCommand           // A function to get the subcommands for this command
	SubCommandGroups() *[]SubCommandGroup // A function to get the subcommand groups for this command
}

/*
A subcommand group slash command
*/
type SubCommandGroup interface {
	BaseCommand

	SubCommands() *[]SubCommand // A function to get the subcommands for this subcommand group
}

/*
A subcommand slash command
*/
type SubCommand interface {
	BaseCommand
}

func GetType(c interface{}) (interface{}, error) {
	var t interface{}

	switch t := c.(type) {
	case Command:
		log.Printf("command type")
	case SubCommandGroup:
		log.Printf("subcommandgroup type")
	case SubCommand:
		log.Printf("subcommand type")
	default:
		log.Printf("unknown type: %v", t)
		return "", errors.New("unknown command type")
	}

	return t, nil
}
