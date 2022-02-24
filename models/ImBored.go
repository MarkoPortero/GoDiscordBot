package models

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/bwmarrin/discordgo"
)

type BoredActivity struct {
	Activity      string  `json:"activity"`
	Type          string  `json:"type"`
	Participants  int     `json:"participants"`
	Price         int     `json:"price"`
	Link          string  `json:"link"`
	Key           string  `json:"key"`
	Accessibility float64 `json:"accessibility"`
}

func Bored(session *discordgo.Session, message *discordgo.MessageCreate) {
	var record BoredActivity
	response, err := http.Get("https://www.boredapi.com/api/activity/")
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		if err := json.NewDecoder(response.Body).Decode(&record); err != nil {
			log.Println(err)
			return
		}
		fmt.Println(record)
		send, err := session.ChannelMessageSend(message.ChannelID, "You should "+record.Activity)
		if err != nil {
			log.Fatal(err)
			return
		}
		log.Println("Correctly sent: ", send)

		if len(record.Link) > 0 {
			messageSend, err := session.ChannelMessageSend(message.ChannelID, "Here's a link "+record.Link)
			if err != nil {
				log.Fatal(err)
				return
			}
			log.Println("Correctly sent: ", messageSend)
		}
	}
}
