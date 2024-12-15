package main

import (
	"errors"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type BotConfig struct {
	ClientId string
	Token    string
}

func ReadEnvConfig() BotConfig {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Unable to load .env file:\n%s", err)

		return BotConfig{}
	}

	var newConfig BotConfig

	newConfig.Token = os.Getenv("BOT_TOKEN")
	newConfig.ClientId = os.Getenv("BOT_CLIENT_ID")

	return newConfig
}

func (config *BotConfig) CheckConfig() error {
	if len(config.ClientId) <= 0 {
		return errors.New("client ID was not specified in the .env file, this is necessary for slash commands")
	}

	if len(config.Token) <= 0 {
		return errors.New("token was not specified in the .env file, this is necessary to start the bot")
	}

	return nil
}
