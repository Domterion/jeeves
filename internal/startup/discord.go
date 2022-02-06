package startup

import (
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/domterion/jeeves/internal/models"
	"github.com/sarulabs/di/v2"
)

func InitDiscord(container di.Container) (*discordgo.Session, error) {
	config := container.Get("config").(*models.Config)

	discord, err := discordgo.New("Bot " + config.Token)
	if err != nil {
		return nil, err
	}

	discord.AddHandler(func(s *discordgo.Session, e *discordgo.Ready) {
		log.Println("Bot is ready!")
	})

	err = discord.Open()
	if err != nil {
		return nil, err
	}

	return discord, err
}
