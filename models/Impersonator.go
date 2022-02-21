package models

import (
	"GoDiscordBot/GlobalVariables"
	"bufio"
	"github.com/bwmarrin/discordgo"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

func Impersonate(session *discordgo.Session, message *discordgo.MessageCreate) {
	personality := "./resources/" + strings.TrimSpace(GlobalVariables.BotPersonality)
	file, err := os.Open(personality)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	randsource := rand.NewSource(time.Now().UnixNano())
	randgenerator := rand.New(randsource)

	lineNum := 1
	var pick string
	for scanner.Scan() {
		line := scanner.Text()

		roll := randgenerator.Intn(lineNum)
		if roll == 0 {
			pick = line
		}

		lineNum += 1
	}
	session.ChannelMessageSend(message.ChannelID, pick)
}

func PersonalityTransfer(session *discordgo.Session, message *discordgo.MessageCreate, personality string) {
	if strings.Contains(strings.ToLower(personality), "mark") || strings.Contains(strings.ToLower(personality), "matt") || strings.Contains(strings.ToLower(personality), "chris") || strings.Contains(strings.ToLower(personality), "tommy") || strings.Contains(strings.ToLower(personality), "gerrit") {
		GlobalVariables.BotPersonality = personality + "Personality.txt"
		session.ChannelMessageSend(message.ChannelID, "Using stored personality for: "+personality)
	} else {
		GlobalVariables.BotPersonality = "dinoPersonality.txt"
		session.ChannelMessageSend(message.ChannelID, "Could not find stored personality for: "+personality+". Reverting to Dinofault.")
	}
}
