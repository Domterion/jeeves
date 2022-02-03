package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"
	"github.com/domterion/jeeves/cmd/jeeves/commands"
	"github.com/domterion/jeeves/commander"
	"github.com/domterion/jeeves/common/config"
	"github.com/domterion/jeeves/database"
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

	commander, err := commander.New(discord, commander.Options{
		GuildID: "897619857187676210",
	})

	if err != nil {
		log.Fatalf("Failed to create command manager: %v", err)
	}

	commander.AddCommand(commands.CreateCommand)
	commander.AddCommand(commands.ProfileCommand)

	err = discord.Open()
	if err != nil {
		log.Fatalf("Failed to open session: %v", err)
	}
	defer discord.Close()

	err = database.Connect(config.Config.DatabaseUri)

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	stop := make(chan os.Signal)
	signal.Notify(stop, os.Interrupt)
	<-stop

	// Do our cleanup stuff here
	log.Println("Gracefully shutting down..")
	database.Db.Close()
}
