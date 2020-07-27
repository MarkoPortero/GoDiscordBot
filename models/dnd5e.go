package models

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/bwmarrin/discordgo"
)

type spells struct {
	Slug               string `json:"slug"`
	Name               string `json:"name"`
	Desc               string `json:"desc"`
	HigherLevel        string `json:"higher_level"`
	Page               string `json:"page"`
	Range              string `json:"range"`
	Components         string `json:"components"`
	Material           string `json:"material"`
	Ritual             string `json:"ritual"`
	Duration           string `json:"duration"`
	Concentration      string `json:"concentration"`
	CastingTime        string `json:"casting_time"`
	Level              string `json:"level"`
	LevelInt           int    `json:"level_int"`
	School             string `json:"school"`
	DndClass           string `json:"dnd_class"`
	Archetype          string `json:"archetype"`
	Circles            string `json:"circles"`
	DocumentSlug       string `json:"document__slug"`
	DocumentTitle      string `json:"document__title"`
	DocumentLicenseURL string `json:"document__license_url"`
}

func GetSpells(session *discordgo.Session, message *discordgo.MessageCreate) {
	var Spell spells
	searchSpell := strings.Split(strings.ToLower(message.Content), "!spells")
	if len(searchSpell) > 1 {
		response, err := http.Get("https://api.open5e.com/spells/" + strings.Trim(searchSpell[1], " "))
		if err != nil {
			fmt.Printf("The HTTP request failed with error %s\n", err)
		} else {
			if err := json.NewDecoder(response.Body).Decode(&Spell); err != nil {
				log.Println(err)
			}
			fmt.Println(Spell)

			if (spells{}) == Spell {
				session.ChannelMessageSend(message.ChannelID, "Sorry, i'm having trouble finding that spell.")
				return
			}

			session.ChannelMessageSend(message.ChannelID,
				"```Spell name: "+Spell.Name+"\n\nDescription: "+Spell.Desc+"\n\nRange: "+Spell.Range+"\n\nBook Page: "+Spell.Page+"```")
		}
	}
}
