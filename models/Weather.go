package models

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/bwmarrin/discordgo"
)

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

func WeatherApiConsumer(session *discordgo.Session, message *discordgo.MessageCreate, placeName string) {
	var record Weather
	var WeatherDescriptor string

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
