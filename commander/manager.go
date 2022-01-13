package commander

// TODO: We need better logging, it should be configurable too

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
)

// A struct to hold all data for the manager
type Manager struct {
	Session  *discordgo.Session     // The discordgo session to use with registering commands and handling  events
	commands map[string]interface{} // All commands registered to the manager, it can be a Command or SubCommand
	options  *Options               // Registered option to act as a configuration
}

// The options, or configuration, for the manager
type Options struct {
	TestGuild      string                            // The ID of a guild to register commands to or empty for global
	OnCommandError func(err error, context *Context) // The function that is fired when there is an error returned from a command run
}

// Construct a new command manager
func New(session *discordgo.Session, options ...Options) (*Manager, error) {
	manager := &Manager{
		Session:  session,
		commands: make(map[string]interface{}),
	}

	manager.options = &Options{
		TestGuild: "",
		OnCommandError: func(err error, context *Context) {
			log.Printf("%v command error: %v", context.CommandName, err)
		},
	}

	if len(options) > 0 {
		if options[0].TestGuild != "" {
			manager.options.TestGuild = options[0].TestGuild
		}

		if options[0].OnCommandError != nil {
			manager.options.OnCommandError = options[0].OnCommandError
		}
	}

	manager.Session.AddHandler(manager.onReady)
	manager.Session.AddHandler(manager.onInteractionCreate)

	return manager, nil
}

// Adds a command to the manager and registers all subcommands
func (m *Manager) AddCommand(command Command) {
	baseCommandName := command.Name

	m.commands[baseCommandName] = command
	// log.Printf("command %s added", baseCommandName)

	for _, subcommand := range command.SubCommands {
		subCommandName := fmt.Sprintf("%s %s", baseCommandName, subcommand.Name)

		// If the base command has a BeforeRun defined and subcommand doesnt then it will use the base commands BefordRun
		if command.BeforeRun != nil && subcommand.BeforeRun == nil {
			subcommand.BeforeRun = command.BeforeRun
		}

		m.commands[subCommandName] = subcommand
		// log.Printf("subcommand %s added", subCommandName)
	}

	for _, subcommandgroup := range command.SubCommandGroups {
		subCommandGroupName := fmt.Sprintf("%s %s", baseCommandName, subcommandgroup.Name)
		// log.Printf("on subcommandGroup %s", subCommandGroupName)

		// If the base command has a BeforeRun defined but the group doesnt then it will use the base commmands
		if command.BeforeRun != nil && subcommandgroup.BeforeRun == nil {
			subcommandgroup.BeforeRun = command.BeforeRun
		}

		for _, subcommand := range subcommandgroup.SubCommands {

			// If the subcommandgroup has a BeforeRun defined but the subcommand doesnt then it will use the groups
			if subcommandgroup.BeforeRun != nil && subcommand.BeforeRun == nil {
				subcommand.BeforeRun = subcommandgroup.BeforeRun
			}

			subCommandName := fmt.Sprintf("%s %s", subCommandGroupName, subcommand.Name)
			m.commands[subCommandName] = subcommand
			// log.Printf("subcommand %s added in group", subCommandName)
		}

		// log.Println("end group")
	}
}

func (m *Manager) onReady(s *discordgo.Session, e *discordgo.Ready) {
	for _, command := range m.commands {
		switch c := command.(type) {
		case Command:
			_, err := m.Session.ApplicationCommandCreate(m.Session.State.User.ID, m.options.TestGuild, c.ToApplicationCommand())

			if err != nil {
				log.Printf("failed to register %v command: %v", c.Name, err)
			}
		}

		continue
	}
}

func (m *Manager) onInteractionCreate(s *discordgo.Session, e *discordgo.InteractionCreate) {
	switch e.Type {
	case discordgo.InteractionMessageComponent:
		m.handleMessageComponent(s, e)
	case discordgo.InteractionApplicationCommand:
		m.handleApplicationCommand(s, e)
	default:
		log.Printf("unsupported interaction type: %v", e.Type)
	}
}

func (m *Manager) handleApplicationCommand(s *discordgo.Session, e *discordgo.InteractionCreate) {
	name := recurseCommandOptions(e.ApplicationCommandData().Options, e.ApplicationCommandData().Name)

	command, exists := m.commands[name]

	if !exists {
		return
	}

	var commandObject BaseCommand

	switch c := command.(type) {
	case Command:
		commandObject = c.BaseCommand
	case *SubCommandGroup:
		commandObject = c.BaseCommand
	case *SubCommand:
		commandObject = c.BaseCommand
	}

	context := Context{
		Session:         m.Session,
		Event:           e,
		Options:         e.ApplicationCommandData().Options,
		CommandName:     name,
		ResolvedOptions: e.ApplicationCommandData().Resolved,
		Member:          e.Member,
	}

	if commandObject.BeforeRun != nil {
		before := commandObject.BeforeRun(&context)

		if !before {
			return
		}
	}

	err := commandObject.Run(&context)

	if err != nil {
		m.options.OnCommandError(err, &context)

		return
	}
}

func (m *Manager) handleMessageComponent(s *discordgo.Session, e *discordgo.InteractionCreate) {
	log.Println("message component")
}

func recurseCommandOptions(options []*discordgo.ApplicationCommandInteractionDataOption, name string) string {
	for _, option := range options {
		if option.Type == discordgo.ApplicationCommandOptionSubCommand || option.Type == discordgo.ApplicationCommandOptionSubCommandGroup {
			newName := fmt.Sprintf("%s %s", name, option.Name)
			return recurseCommandOptions(option.Options, newName)
		}

		continue
	}

	return name
}
