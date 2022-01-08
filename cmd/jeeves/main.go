package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"
	"github.com/domterion/jeeves/cmd/jeeves/commands"
	"github.com/domterion/jeeves/common/config"
	"github.com/domterion/jeeves/handler"
)

func main() {
	if err := config.Load(); err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	discord, err := discordgo.New("Bot " + config.Config.Token)

	if err != nil {
		log.Fatalf("Failed to create discordgo client: %v", err)
	}

	discord.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Println("Bot is ready!")
	})

	err = discord.Open()
	if err != nil {
		log.Fatalf("Failed to open session: %v", err)
	}

	defer discord.Close()

	commandManager, err := handler.New(discord)
	if err != nil {
		log.Fatalf("Failed to create command manager: %v", err)
	}

	commandManager.AddCommand(commands.OwnerCommand)
	commandManager.AddCommand(commands.PingCommand)
	commandManager.AddCommand(commands.UserInfoCommand)

	ownercommand := commands.OwnerCommand.ToApplicationCommand()
	pingcommand := commands.PingCommand.ToApplicationCommand()
	userinfocommand := commands.UserInfoCommand.ToApplicationCommand()

	_, err = discord.ApplicationCommandCreate(discord.State.User.ID, "897619857187676210", ownercommand)
	if err != nil {
		log.Fatalf("Failed to register owner command: %v", err)
	}

	_, err = discord.ApplicationCommandCreate(discord.State.User.ID, "897619857187676210", pingcommand)
	if err != nil {
		log.Fatalf("Failed to register ping command: %v", err)
	}

	_, err = discord.ApplicationCommandCreate(discord.State.User.ID, "897619857187676210", userinfocommand)
	if err != nil {
		log.Fatalf("Failed to register userinfo command: %v", err)
	}

	stop := make(chan os.Signal)
	signal.Notify(stop, os.Interrupt)
	<-stop
	log.Println("Gracefully shutting down..")
}
