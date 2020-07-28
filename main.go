package main

import (
	"GoDiscordBot/handlers"
	"flag"
	"fmt"

	"github.com/bwmarrin/discordgo"
)

const token string = "NzM0MDE2NjQ4NzcxNTM0ODg5.XxLkpQ.t8OCxa0STBQK5jWJc95FveAUXhU"

var (
	Token      string
	AvatarFile string
	AvatarURL  string
	BotID      string
)

func init() {
	flag.StringVar(&AvatarURL, "u", "", "https://img.sunset02.com/sites/default/files/styles/4_3_horizontal_inbody_900x506/public/image/2016/09/main/dungeness-crab.jpg")
	flag.Parse()
}

func main() {
	discordgo, error := discordgo.New("Bot " + token)
	if error != nil {
		fmt.Println(error.Error())
		return
	}
	User, error := discordgo.User("@me")
	if error != nil {
		fmt.Println(error.Error())
	}
	BotID = User.ID

	discordgo.AddHandler(handlers.MessageHandler)
	//open connection
	error = discordgo.Open()
	if error != nil {
		fmt.Println(error.Error())
		return
	}
	fmt.Println("ITS ALIVE")

	//new channel to keep bot running
	<-make(chan struct{})
	return
}
