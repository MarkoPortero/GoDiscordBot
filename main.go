package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

const token string = "NzM0MDE2NjQ4NzcxNTM0ODg5.XxLkpQ.t8OCxa0STBQK5jWJc95FveAUXhU"

var (
	BotID             string
	Token             string
	AvatarFile        string
	AvatarURL         string
	ToRemind          string
	WhatImReminding   string
	TimeToRemind      string
	ToRemindPing      string
	Place             string
	WeatherDescriptor string
	Country           string
)

type Bored struct {
	Activity      string  `json:"activity"`
	Type          string  `json:"type"`
	Participants  int     `json:"participants"`
	Price         int     `json:"price"`
	Link          string  `json:"link"`
	Key           string  `json:"key"`
	Accessibility float64 `json:"accessibility"`
}

type Corona []struct {
	Country   string `json:"Country"`
	Confirmed int    `json:"Confirmed"`
	Deaths    int    `json:"Deaths"`
	Recovered int    `json:"Recovered"`
	Active    int    `json:"Active"`
}
type doggyJson struct {
	Url string `json:"url"`
}
type Weather struct {
	Weather []DetailWeather        `json:"weather"`
	Main    MainWeatherInformation `json:"main"`
}
type DetailWeather struct {
	Description string `json:"description"`
}
type MainWeatherInformation struct {
	Temperature float64 `json:"temp"`
}

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
	err = dg.Open()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("ITS ALIVE")
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
		ToRemind = message.Author.ID
		ToRemindPing = message.Author.Mention()
		remindMe(session, message)
		return
	}
	if strings.Contains(strings.ToLower(message.Content), "weather in") {
		breakDown := strings.Split(strings.ToLower(message.Content), "weather in ")
		if len(breakDown) > 1 {
			Place = breakDown[1]
			weatherApiConsumer(session, message, Place)
		}
		return
	}
	if strings.Contains(strings.ToLower(message.Content), "dog") {
		doggyStyle(session, message)
		return
	}
	if strings.Contains(strings.ToLower(message.Content), "coronavirus in") {
		breakDown := strings.Split(strings.ToLower(message.Content), "coronavirus in")
		if len(breakDown) > 1 {
			Country = breakDown[1]
			coronavirusStats(session, message, Country)
		}
		return
	}
	if strings.Contains(strings.ToLower(message.Content), "bored") {
		bored(session, message)
		return
	}
	watchYoProfanity(session, message)
}

func remindMe(session *discordgo.Session, message *discordgo.MessageCreate) {
	messageContent := message.Content
	breakDown := strings.Split(strings.ToLower(messageContent), "remind me to ")
	if len(breakDown) > 1 {
		TimeToRemind = breakDown[0]
		WhatImReminding = breakDown[1]
		fmt.Println(ToRemind)
		session.ChannelMessageSend(message.ChannelID, "Certainly. I'll remind you to "+WhatImReminding+", young Padawan.")
	}
	if strings.Contains(strings.ToLower(TimeToRemind), "minute") {
		timer1 := time.NewTimer(time.Minute * 1)
		<-timer1.C
		session.ChannelMessageSend(message.ChannelID, "Oi, "+ToRemindPing+", i'm reminding you to "+WhatImReminding)
	}
	if strings.Contains(strings.ToLower(TimeToRemind), "hour") {
		timer1 := time.NewTimer(time.Hour * 1)
		<-timer1.C
		session.ChannelMessageSend(message.ChannelID, "Oi, "+ToRemindPing+", i'm reminding you to "+WhatImReminding)
	}
}
func weatherApiConsumer(session *discordgo.Session, message *discordgo.MessageCreate, placeName string) {
	var record Weather
	fmt.Println("Looking up weather")
	if placeName == "dublin" {
		placeName = "dublin, ie"
	}
	response, err := http.Get("http://api.openweathermap.org/data/2.5/weather?q=" + placeName + "&appid=bbabb7e7587c5280e2e68130d651f4d9&units=metric")
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		if err := json.NewDecoder(response.Body).Decode(&record); err != nil {
			log.Println(err)
		}
		if record.Main.Temperature > 0 {
			if record.Main.Temperature < 10 {
				WeatherDescriptor = "chilly"
			} else if record.Main.Temperature >= 10 && record.Main.Temperature < 20 {
				WeatherDescriptor = "mild"
			} else if record.Main.Temperature > 20 {
				WeatherDescriptor = "warm"
			}
			session.ChannelMessageSend(message.ChannelID, "Currently in "+placeName+" it's "+
				record.Weather[0].Description+", at a "+WeatherDescriptor+" "+fmt.Sprint(record.Main.Temperature)+" degrees.")
		}
		fmt.Println(record.Weather)
		fmt.Println(record.Main)
	}
}
func doggyStyle(session *discordgo.Session, message *discordgo.MessageCreate) {
	var record doggyJson
	response, err := http.Get("https://random.dog/woof.json")
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		if err := json.NewDecoder(response.Body).Decode(&record); err != nil {
			log.Println(err)
		}
		session.ChannelMessageSend(message.ChannelID, record.Url)
	}
}

func coronavirusStats(session *discordgo.Session, message *discordgo.MessageCreate, country string) {
	var record Corona
	response, err := http.Get("https://api.covid19api.com/total/country/" + strings.Trim(country, " "))
	fmt.Println("https://api.covid19api.com/total/country/" + strings.Trim(country, " "))
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		if err := json.NewDecoder(response.Body).Decode(&record); err != nil {
			log.Println(err)
		}
		fmt.Println(record)

		if len(record) == 0 {
			session.ChannelMessageSend(message.ChannelID, "Sorry, i'm having trouble finding that location. Please try again.")
			return
		}
		currentRecord := record[len(record)-1]
		session.ChannelMessageSend(message.ChannelID, "Currently in "+
			currentRecord.Country+" there are "+fmt.Sprint(currentRecord.Confirmed)+
			" confirmed cases. There have been "+fmt.Sprint(currentRecord.Deaths)+
			" deaths so far. However, a total of "+fmt.Sprint(currentRecord.Recovered)+" have recovered.")
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

func bored(session *discordgo.Session, message *discordgo.MessageCreate) {
	var record Bored
	response, err := http.Get("https://www.boredapi.com/api/activity/")
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		if err := json.NewDecoder(response.Body).Decode(&record); err != nil {
			log.Println(err)
		}
		fmt.Println(record)
		session.ChannelMessageSend(message.ChannelID, "You should "+record.Activity)
		if len(record.Link) > 0 {
			session.ChannelMessageSend(message.ChannelID, "Here's a link "+record.Link)
		}
	}
}
