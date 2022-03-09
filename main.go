package main

import (
	"GoDiscordBot/BotSetup"
	"GoDiscordBot/GlobalVariables"
)

func main() {
	GlobalVariables.LoadEnvVariables()
	BotSetup.FullDiscordSetup()
}
