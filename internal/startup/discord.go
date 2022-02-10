package startup

import (
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/domterion/jeeves/internal/models"
	"github.com/domterion/jeeves/internal/utils"
	"github.com/sarulabs/di/v2"
)

func InitDiscord(container di.Container) (*discordgo.Session, error) {
	config := container.Get(utils.DIConfig).(*models.Config)

	discord := container.Get(utils.DIDiscord).(*discordgo.Session)

	discord.Token = "Bot " + config.Token

	discord.AddHandler(func(s *discordgo.Session, e *discordgo.Ready) {
		log.Println("Bot is ready!")
	})

	err := discord.Open()
	if err != nil {
		return nil, err
	}

	return discord, err
}
