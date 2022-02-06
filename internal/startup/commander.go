package startup

import (
	"github.com/bwmarrin/discordgo"
	"github.com/domterion/jeeves/cmd/jeeves/commands"
	"github.com/domterion/jeeves/pkg/commander"
	"github.com/sarulabs/di/v2"
)

var commandSlice = []commander.Command{commands.CreateCommand, commands.ProfileCommand}

func InitCommander(container di.Container) (*commander.Manager, error) {
	discord := container.Get("discord").(*discordgo.Session)
	commander, err := commander.New(discord, commander.Options{
		GuildID:            "897619857187676210",
		DependencyProvider: container,
	})

	if err != nil {
		return nil, err
	}

	for _, command := range commandSlice {
		commander.AddCommand(command)
	}

	return commander, err
}
