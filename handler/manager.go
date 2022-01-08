package handler

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

type Manager struct {
	Session *discordgo.Session
	// This has to be an interface{} because we have multiple command types ie Command, SubCommandGroup and SubCommand
	Commands map[string]interface{}
}

func New(session *discordgo.Session) (*Manager, error) {
	manager := &Manager{
		Session:  session,
		Commands: make(map[string]interface{}),
	}

	return manager, nil
}

func (m *Manager) AddCommand(command Command) {
	baseCommandName := command.Name

	m.Commands[baseCommandName] = command
	fmt.Printf("Command %s added \n", baseCommandName)

	for _, subcommand := range command.SubCommands {
		subCommandName := fmt.Sprintf("%s %s", baseCommandName, subcommand.Name)
		m.Commands[subCommandName] = subcommand
		fmt.Printf("Subcommand %s added \n", subCommandName)
	}

	for _, subcommandgroup := range command.SubCommandGroups {
		subCommandGroupName := fmt.Sprintf("%s %s", baseCommandName, subcommandgroup.Name)
		fmt.Printf("On subcommandGroup %s \n", subCommandGroupName)

		for _, subcommand := range command.SubCommands {
			subCommandName := fmt.Sprintf("%s %s", subCommandGroupName, subcommand.Name)
			m.Commands[subCommandName] = subcommand
			fmt.Printf("Subcommand %s added in group\n", subCommandName)
		}

		fmt.Println("END GROUP")
	}
}
