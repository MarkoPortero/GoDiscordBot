package BotSetup

import (
	"GoDiscordBot/GlobalVariables"
	"GoDiscordBot/handlers"
	"fmt"
	"github.com/bwmarrin/discordgo"
)

const token string = "NzM0MDE2NjQ4NzcxNTM0ODg5.XxLkRg.gMqoU7uapRlz6Ix2UFmDtcWqBVM"

func FullDiscordSetup() {
	discGo, err, done := SetupBot()
	if done {
		return
	}

	discGo.AddHandler(handlers.MessageHandler)
	//open connection
	err = discGo.Open()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("ITS ALIVE")

	//new channel to keep bot running
	<-make(chan struct{})
	return
}

func SetupBot() (*discordgo.Session, error, bool) {
	discGo, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println(err.Error())
		return nil, nil, true
	}
	User, err := discGo.User("@me")
	if err != nil {
		panic(err)
	}
	GlobalVariables.BotID = User.ID
	fmt.Println(GlobalVariables.BotID)
	return discGo, err, false
}
