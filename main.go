package main

import (
	"fmt"
	"net/http"
)

var MainBotConfig BotConfig
var MainBot Bot
var httpClient = &http.Client{}

func main() {
	MainBotConfig = ReadEnvConfig()
	if err := MainBotConfig.CheckConfig(); err != nil {
		fmt.Println(err)
		return
	}

	MainBot = NewBot()

	MainBot.RegisterCommands(BuildedCommands)
	MainBot.Start()
}
