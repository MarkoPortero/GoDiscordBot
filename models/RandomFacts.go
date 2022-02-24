package models

import (
	"encoding/json"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"net/http"
)

type Facts struct {
	ID        string `json:"id"`
	Text      string `json:"text"`
	Source    string `json:"source"`
	SourceURL string `json:"source_url"`
	Language  string `json:"language"`
	Permalink string `json:"permalink"`
}

// GetNews gets news article for location
func GetFact(session *discordgo.Session, message *discordgo.MessageCreate) {
	var fact Facts
	response, err := http.Get("https://uselessfacts.jsph.pl/random.json?language=en")
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		if err := json.NewDecoder(response.Body).Decode(&fact); err != nil {
			log.Println(err)
		}
		fmt.Println(fact)

		currentIterator = 0
		send, err := session.ChannelMessageSend(message.ChannelID, fact.Text)
		if err != nil {
			log.Fatal(err)
		}
		log.Println("Correctly sent: ", send)
	}
}
