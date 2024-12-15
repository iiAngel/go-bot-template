package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"
)

type Command struct {
	RegisteredCommand *discordgo.ApplicationCommand
	Data              *discordgo.ApplicationCommand
	Func              CommandFunc
}

type Bot struct {
	Session  *discordgo.Session
	Commands map[string]Command
}

func NewBot() Bot {
	var err error
	var b Bot

	b.Session, err = discordgo.New("Bot " + MainBotConfig.Token)

	if err != nil {
		log.Fatal(err)
	}

	b.Commands = make(map[string]Command)

	return b
}

func (b *Bot) RegisterCommands(commands map[string]BuildCommand) {
	for name, command := range commands {
		var commandData Command

		commandData.Data = command.Data
		commandData.Func = command.Func

		// Create command so it works on discords end

		registeredCommand, err := b.Session.ApplicationCommandCreate(MainBotConfig.ClientId, "", command.Data)
		if err != nil {
			log.Println(err)
			continue
		}

		commandData.RegisteredCommand = registeredCommand
		b.Commands[name] = commandData

		fmt.Printf("Registered command: %s\n", name)
	}
}

func (b *Bot) Stop() {
	for _, command := range b.Commands {
		if err := b.Session.ApplicationCommandDelete(MainBotConfig.ClientId, "", command.RegisteredCommand.ID); err != nil {
			log.Println(err)
			continue
		}

		fmt.Printf("Removed command: %s\n", command.Data.Name)
	}

	if err := b.Session.Close(); err != nil {
		log.Fatal(err)
	}
}

func (b *Bot) Start() {
	// Know when bot is logged in
	b.Session.AddHandler(func(_ *discordgo.Session, _ *discordgo.Ready) {
		fmt.Printf("Logged in as: %s\n", b.Session.State.User.Username)
	})

	// Interaction handler (Slash commands)
	b.Session.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if command, ok := b.Commands[i.ApplicationCommandData().Name]; ok {
			command.Func(s, i)
		}
	})

	if err := b.Session.Open(); err != nil {
		log.Fatal(err)
	}

	// Not stop the bot until the signal channel gets updated
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	fmt.Println("Press Ctrl+C to stop the bot!")

	<-stop

	b.Stop()
}
