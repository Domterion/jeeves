package commander

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

type Runnable func(context *Context) error

type BaseCommand struct {
	Name        string                                // The name of the command to register with Discord
	Description string                                // The description for the command
	Type        discordgo.ApplicationCommandType      // The type of the command, User, Message or Chat
	Options     []*discordgo.ApplicationCommandOption // Options for the command
	Run         Runnable                              // The handler function for the command
}

type Command struct {
	BaseCommand

	SubCommands      []*SubCommand
	SubCommandGroups []*SubCommandGroup
}

type SubCommand struct {
	BaseCommand
}

type SubCommandGroup struct {
	BaseCommand
	SubCommands []*SubCommand
}

func (c Command) ToApplicationCommand() *discordgo.ApplicationCommand {
	switch c.Type {
	case discordgo.ChatApplicationCommand:

		var options []*discordgo.ApplicationCommandOption

		for _, subcommandgroup := range c.SubCommandGroups {
			options = append(options, subcommandgroup.ToOption())
		}

		for _, subcommand := range c.SubCommands {
			options = append(options, subcommand.ToOption())
		}

		options = append(options, c.Options...)

		return &discordgo.ApplicationCommand{
			Name:        c.Name,
			Description: c.Description,
			Options:     options,
			Type:        discordgo.ChatApplicationCommand,
		}
	case discordgo.UserApplicationCommand:
		return &discordgo.ApplicationCommand{
			Name: c.Name,
			Type: discordgo.UserApplicationCommand,
		}
	case discordgo.MessageApplicationCommand:
		return &discordgo.ApplicationCommand{
			Name: c.Name,
			Type: discordgo.MessageApplicationCommand,
		}
	default:
		log.Fatal("unknown command type")
		return nil
	}
}

func (c SubCommandGroup) ToOption() *discordgo.ApplicationCommandOption {
	var subcommands []*discordgo.ApplicationCommandOption

	for _, subcommand := range c.SubCommands {
		subcommands = append(subcommands, subcommand.ToOption())
	}

	return &discordgo.ApplicationCommandOption{
		Name:        c.Name,
		Description: c.Description,
		Options:     subcommands,
		Type:        discordgo.ApplicationCommandOptionSubCommandGroup,
	}
}

func (c SubCommand) ToOption() *discordgo.ApplicationCommandOption {
	return &discordgo.ApplicationCommandOption{
		Name:        c.Name,
		Description: c.Description,
		Options:     c.Options,
		Type:        discordgo.ApplicationCommandOptionSubCommand,
	}
}
