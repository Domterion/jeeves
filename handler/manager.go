package handler

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
)

type Manager struct {
	Session *discordgo.Session
	// This has to be an interface{} because we have multiple command types ie Command, SubCommandGroup and SubCommand
	commands  map[string]interface{}
	testGuild string
}

func New(session *discordgo.Session, testGuild string) (*Manager, error) {
	manager := &Manager{
		Session:   session,
		commands:  make(map[string]interface{}),
		testGuild: testGuild,
	}

	manager.Session.AddHandler(manager.onReady)
	manager.Session.AddHandler(manager.onInteractionCreate)

	return manager, nil
}

func (m *Manager) AddCommand(command Command) {
	baseCommandName := command.Name

	m.commands[baseCommandName] = command
	fmt.Printf("Command %s added \n", baseCommandName)

	for _, subcommand := range command.SubCommands {
		subCommandName := fmt.Sprintf("%s %s", baseCommandName, subcommand.Name)
		m.commands[subCommandName] = subcommand
		fmt.Printf("Subcommand %s added \n", subCommandName)
	}

	for _, subcommandgroup := range command.SubCommandGroups {
		subCommandGroupName := fmt.Sprintf("%s %s", baseCommandName, subcommandgroup.Name)
		fmt.Printf("On subcommandGroup %s \n", subCommandGroupName)

		for _, subcommand := range command.SubCommands {
			subCommandName := fmt.Sprintf("%s %s", subCommandGroupName, subcommand.Name)
			m.commands[subCommandName] = subcommand
			fmt.Printf("Subcommand %s added in group\n", subCommandName)
		}

		fmt.Println("END GROUP")
	}
}

func (m *Manager) onReady(s *discordgo.Session, e *discordgo.Ready) {
	for _, command := range m.commands {
		switch c := command.(type) {
		case Command:
			_, err := m.Session.ApplicationCommandCreate(m.Session.State.User.ID, m.testGuild, c.ToApplicationCommand())

			if err != nil {
				log.Fatalf("Failed to register %v command: %v", c.Name, err)
			}
		}

		continue
	}
}

func (m *Manager) onInteractionCreate(s *discordgo.Session, e *discordgo.InteractionCreate) {
	fmt.Println(e.Interaction.ApplicationCommandData().Name)
}