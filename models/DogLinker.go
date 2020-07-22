package models

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/bwmarrin/discordgo"
)

type doggyJson struct {
	Url string `json:"url"`
}

func DoggyStyle(session *discordgo.Session, message *discordgo.MessageCreate) {
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
