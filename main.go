package main

import (
	"fmt"
)

var MainBotConfig BotConfig
var MainBot Bot

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
