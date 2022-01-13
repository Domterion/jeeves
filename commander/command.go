package commander

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

// A base command that all commands must implement
type BaseCommand struct {
	Name        string                                // The name of the command to register with Discord
	Description string                                // The description for the command
	Type        discordgo.ApplicationCommandType      // The type of the command, User, Message or Chat
	Options     []*discordgo.ApplicationCommandOption // Options for the command
	BeforeRun    func(context *Context) bool          // The function called before Run, typically used for checks
	Run         func(context *Context) error          // The handler function for the command
}

// A root command
type Command struct {
	BaseCommand

	SubCommands      []*SubCommand      // A slice of subcommands for this root command
	SubCommandGroups []*SubCommandGroup // A slice of subcommand groups for this command
}

// A subcommand
type SubCommand struct {
	BaseCommand
}

// A subcommand group
type SubCommandGroup struct {
	BaseCommand
	SubCommands []*SubCommand // A slice of subcommands for this subcommand group
}

// Converts a commander command object to a discordgo application command
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

// Converts a subcommand group and its subcommands to one discordgo option type
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

// Converts a subcommand to a discordgo option
func (c SubCommand) ToOption() *discordgo.ApplicationCommandOption {
	return &discordgo.ApplicationCommandOption{
		Name:        c.Name,
		Description: c.Description,
		Options:     c.Options,
		Type:        discordgo.ApplicationCommandOptionSubCommand,
	}
}
