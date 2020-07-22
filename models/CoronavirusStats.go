package models

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/bwmarrin/discordgo"
)

type Corona []struct {
	Country   string `json:"Country"`
	Confirmed int    `json:"Confirmed"`
	Deaths    int    `json:"Deaths"`
	Recovered int    `json:"Recovered"`
	Active    int    `json:"Active"`
}

func CoronavirusStats(session *discordgo.Session, message *discordgo.MessageCreate, country string) {
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
