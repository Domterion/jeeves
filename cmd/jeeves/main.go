package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"
	"github.com/domterion/jeeves/cmd/jeeves/commands"
	"github.com/domterion/jeeves/common/config"
	"github.com/domterion/jeeves/commander"
)

func main() {
	if err := config.Load(); err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	discord, err := discordgo.New("Bot " + config.Config.Token)

	if err != nil {
		log.Fatalf("Failed to create discordgo client: %v", err)
	}

	discord.AddHandler(func(s *discordgo.Session, e *discordgo.Ready) {
		log.Println("Bot is ready!")
	})

	commandManager, err := commander.New(discord, commander.Options{
		TestGuild: "897619857187676210",
	})
	if err != nil {
		log.Fatalf("Failed to create command manager: %v", err)
	}

	commandManager.AddCommand(commands.OwnerCommand)
	commandManager.AddCommand(commands.PingCommand)
	commandManager.AddCommand(commands.SayCommand)
	commandManager.AddCommand(commands.UserInfoCommand)

	err = discord.Open()
	if err != nil {
		log.Fatalf("Failed to open session: %v", err)
	}

	defer discord.Close()

	stop := make(chan os.Signal)
	signal.Notify(stop, os.Interrupt)
	<-stop
	log.Println("Gracefully shutting down..")
}
