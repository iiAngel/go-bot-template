package main

import (
	"github.com/bwmarrin/discordgo"
)

type CommandFunc func(s *discordgo.Session, i *discordgo.InteractionCreate)
type BuildCommand struct {
	Data *discordgo.ApplicationCommand
	Func CommandFunc
}

var BuildedCommands = map[string]BuildCommand{
	// "Ping" command
	"ping": {
		Data: &discordgo.ApplicationCommand{
			Name:        "ping",
			Description: "Replies with \"pong\"!",
			Type:        discordgo.ChatApplicationCommand,
			Contexts: &[]discordgo.InteractionContextType{
				discordgo.InteractionContextGuild,
				discordgo.InteractionContextBotDM,
				discordgo.InteractionContextPrivateChannel,
			},
			IntegrationTypes: &[]discordgo.ApplicationIntegrationType{
				discordgo.ApplicationIntegrationGuildInstall,
				discordgo.ApplicationIntegrationUserInstall,
			},
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "message",
					Description: "Adds the message to the reply",
					Required:    true,
				},
			},
		},
		Func: func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			additionalMessage := i.Interaction.ApplicationCommandData().GetOption("message").StringValue()

			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "Pong!\n\nAdditional message: " + additionalMessage,
				},
			})
		},
	},
}
