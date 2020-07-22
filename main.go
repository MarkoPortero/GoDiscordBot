package main

import (
	"GoDiscordBot/models"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/bwmarrin/discordgo"
)

const token string = "NzM0MDE2NjQ4NzcxNTM0ODg5.XxLkpQ.t8OCxa0STBQK5jWJc95FveAUXhU"

var (
	BotID      string
	Token      string
	AvatarFile string
	AvatarURL  string
	Place      string
	Country    string
)

func init() {
	flag.StringVar(&AvatarURL, "u", "", "https://img.sunset02.com/sites/default/files/styles/4_3_horizontal_inbody_900x506/public/image/2016/09/main/dungeness-crab.jpg")
	flag.Parse()
}

func main() {
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	u, err := dg.User("@me")
	if err != nil {
		fmt.Println(err.Error())
	}
	BotID = u.ID
	dg.AddHandler(messageHandler)
	//open connection
	err = dg.Open()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("ITS ALIVE")

	//new channel to keep bot running
	<-make(chan struct{})
	return
}

func messageHandler(session *discordgo.Session, message *discordgo.MessageCreate) {
	if message.Author.ID == BotID {
		return
	}

	fmt.Println(message.Author)
	fmt.Println(message.Content)

	if message.Content == "ping" {
		if message.Author.String() == "cactusapple#2171" {
			session.ChannelMessageSend(message.ChannelID, "fuck you Matt. Go pick up your package.")
			return
		}
		_, _ = session.ChannelMessageSend(message.ChannelID, "hello "+message.Author.Mention())
	}

	if strings.Contains(strings.ToLower(message.Content), "hello") {
		session.ChannelMessageSend(message.ChannelID, "https://giphy.com/gifs/mrw-top-escalator-Nx0rz3jtxtEre")
		return
	}

	if strings.Contains(strings.ToLower(message.Content), "remind me") {
		fmt.Println(message.Author.ID)
		ToRemind := message.Author.ID
		ToRemindPing := message.Author.Mention()
		models.RemindMe(session, message, ToRemind, ToRemindPing)
		return
	}

	if strings.Contains(strings.ToLower(message.Content), "weather in") {
		breakDown := strings.Split(strings.ToLower(message.Content), "weather in ")
		if len(breakDown) > 1 {
			Place = breakDown[1]
			models.WeatherApiConsumer(session, message, Place)
		}
		return
	}

	if strings.Contains(strings.ToLower(message.Content), "dog") {
		models.DoggyStyle(session, message)
		return
	}

	if strings.Contains(strings.ToLower(message.Content), "coronavirus in") {
		breakDown := strings.Split(strings.ToLower(message.Content), "coronavirus in")
		if len(breakDown) > 1 {
			Country = breakDown[1]
			models.CoronavirusStats(session, message, Country)
		}
		return
	}

	if strings.Contains(strings.ToLower(message.Content), "bored") {
		models.Bored(session, message)
		return
	}
}

func watchYoProfanity(session *discordgo.Session, message *discordgo.MessageCreate) {
	response, err := http.Get("https://www.purgomalum.com/service/containsprofanity?text=" + message.Content)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		bodyBytes, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}
		body := string(bodyBytes)
		if strings.Contains(body, "true") {
			session.ChannelMessageSend(message.ChannelID, "https://tenor.com/view/watch-your-profanity-funny-gif-5600117")
		}
		fmt.Println(body)
	}
}
