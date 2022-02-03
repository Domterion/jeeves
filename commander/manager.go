package commander

// TODO: We need better logging, it should be configurable too

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/bwmarrin/snowflake"
)

// A struct to hold all data for the manager
type Manager struct {
	Session       *discordgo.Session     // The discordgo session to use with registering commands and handling  events
	commands      map[string]interface{} // All commands registered to the manager, it can be a Command or SubCommand
	components    map[string]interface{} // Component listeners, can be any of Button and SelectMenu
	options       *Options               // Registered option to act as a configuration
	SnowflakeNode *snowflake.Node
}

// The options, or configuration, for the manager
type Options struct {
	GuildID         string                                   // The ID of a guild to register commands to or empty for global
	OnCommandError  func(err error, context *CommandContext) // The function that is fired when there is an error returned from a command run
	SnowflakeNodeId int64                                    // The id to use for the snowflake node
}

// Construct a new command manager
func New(session *discordgo.Session, options ...Options) (*Manager, error) {
	manager := &Manager{
		Session:       session,
		commands:      make(map[string]interface{}),
		components:    make(map[string]interface{}),
		SnowflakeNode: nil,
	}

	manager.options = &Options{
		GuildID: "",
		OnCommandError: func(err error, context *CommandContext) {
			log.Printf("%v command error: %v", context.Name, err)
		},
		SnowflakeNodeId: 1,
	}

	if len(options) > 0 {
		if options[0].GuildID != "" {
			manager.options.GuildID = options[0].GuildID
		}

		if options[0].OnCommandError != nil {
			manager.options.OnCommandError = options[0].OnCommandError
		}

		if options[0].SnowflakeNodeId != 0 {
			manager.options.SnowflakeNodeId = options[0].SnowflakeNodeId
		}
	}

	manager.Session.AddHandler(manager.onReady)
	manager.Session.AddHandler(manager.onInteractionCreate)

	node, err := snowflake.NewNode(manager.options.SnowflakeNodeId)
	if err != nil {
		return nil, err
	}
	manager.SnowflakeNode = node

	return manager, nil
}

// Adds a command to the manager and registers all subcommands
func (m *Manager) AddCommand(command Command) {
	baseCommandName := command.Name

	// Base command
	m.commands[baseCommandName] = command

	// SubCommands
	for _, subcommand := range command.SubCommands {
		subCommandName := fmt.Sprintf("%s %s", baseCommandName, subcommand.Name)
		if command.BeforeRun != nil && subcommand.BeforeRun == nil {
			subcommand.BeforeRun = command.BeforeRun
		}
		m.commands[subCommandName] = subcommand
	}

	// SubCommandGroups
	for _, subcommandgroup := range command.SubCommandGroups {
		subCommandGroupName := fmt.Sprintf("%s %s", baseCommandName, subcommandgroup.Name)

		// If the base command has a BeforeRun defined but the group doesnt then it will use the base commmands
		if command.BeforeRun != nil && subcommandgroup.BeforeRun == nil {
			subcommandgroup.BeforeRun = command.BeforeRun
		}

		// SubCommands of the SubCommandGroup
		for _, subcommand := range subcommandgroup.SubCommands {
			// If the subcommandgroup has a BeforeRun defined but the subcommand doesnt then it will use the groups
			if subcommandgroup.BeforeRun != nil && subcommand.BeforeRun == nil {
				subcommand.BeforeRun = subcommandgroup.BeforeRun
			}

			subCommandName := fmt.Sprintf("%s %s", subCommandGroupName, subcommand.Name)
			m.commands[subCommandName] = subcommand
		}
	}
}

func (m *Manager) addComponent(component interface{}) {
	switch c := component.(type) {
	case Button:
		m.components[c.CustomID] = c
	case SelectMenu:
		m.components[c.CustomID] = c
	case ActionRow:
		for _, comp := range c.components {
			switch c_ := comp.(type) {
			case Button:
				m.components[c_.CustomID] = c_
			case SelectMenu:
				m.components[c_.CustomID] = c_
			}
		}
	default:
		log.Printf("unsupported component type")
	}
}

func (m *Manager) AddComponents(components Components) {
	for _, c := range components.components {
		m.addComponent(c)
	}
}

func (m *Manager) onReady(s *discordgo.Session, e *discordgo.Ready) {
	for _, command := range m.commands {
		switch c := command.(type) {
		case Command:
			_, err := m.Session.ApplicationCommandCreate(m.Session.State.User.ID, m.options.GuildID, c.ToApplicationCommand())

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

	// The only commands that can be ran are commands and subcommands so anything else shouldnt pass this
	switch c := command.(type) {
	case Command:
		commandObject = c.BaseCommand
	case *SubCommand:
		commandObject = c.BaseCommand
	default:
		// This should never be a problem but better safe than sorry
		log.Fatalf("unsupported command type called..?")
	}

	context := CommandContext{
		Session:         m.Session,
		Event:           e,
		Manager:         m,
		Options:         e.ApplicationCommandData().Options,
		Name:            name,
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
	component, exists := m.components[e.MessageComponentData().CustomID]

	if !exists {
		return
	}

	var componentObject BaseComponent

	switch c := component.(type) {
	case Button:
		componentObject = c.BaseComponent
	case SelectMenu:
		componentObject = c.BaseComponent
	default:
		// This should never be a problem but better safe than sorry
		log.Fatalf("unsupported component type called..?")
	}

	context := ComponentContext{
		Session: m.Session,
		Event:   e,
		Manager: m,
		Name:    e.MessageComponentData().CustomID,
		Member:  e.Member,
	}

	// TODO: We should handle the error for this somehow
	componentObject.Run(&context)
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
