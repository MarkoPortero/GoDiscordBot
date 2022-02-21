package BotSetup

import (
	"GoDiscordBot/GlobalVariables"
	"GoDiscordBot/handlers"
	"fmt"
	"github.com/bwmarrin/discordgo"
)

const token string = "NzM0MDE2NjQ4NzcxNTM0ODg5.XxLkRg.gMqoU7uapRlz6Ix2UFmDtcWqBVM"

func FullDiscordSetup() {
	discGo, error, done := SetupBot()
	if done {
		return
	}

	discGo.AddHandler(handlers.MessageHandler)
	//open connection
	error = discGo.Open()
	if error != nil {
		fmt.Println(error.Error())
		return
	}
	fmt.Println("ITS ALIVE")

	//new channel to keep bot running
	<-make(chan struct{})
	return
}

func SetupBot() (*discordgo.Session, error, bool) {
	discGo, error := discordgo.New("Bot " + token)
	if error != nil {
		fmt.Println(error.Error())
		return nil, nil, true
	}
	User, error := discGo.User("@me")
	if error != nil {
		fmt.Println(error.Error())
	}
	GlobalVariables.BotID = User.ID
	fmt.Println(GlobalVariables.BotID)
	return discGo, error, false
}
