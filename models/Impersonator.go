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
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	scanner := bufio.NewScanner(file)

	randSource := rand.NewSource(time.Now().UnixNano())
	randGenerator := rand.New(randSource)

	lineNum := 1
	var pick string
	for scanner.Scan() {
		line := scanner.Text()

		roll := randGenerator.Intn(lineNum)
		if roll == 0 {
			pick = line
		}

		lineNum += 1
	}
	send, err := session.ChannelMessageSend(message.ChannelID, pick)
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Println("Correctly sent: ", send)
}

func PersonalityTransfer(session *discordgo.Session, message *discordgo.MessageCreate, personality string) {
	if isValidPersonality(personality) {
		GlobalVariables.BotPersonality = personality + "Personality.txt"
		send, err := session.ChannelMessageSend(message.ChannelID, "Using stored personality for: "+personality)
		if err != nil {
			log.Fatal(err)
			return
		}
		log.Println("Correctly sent: ", send)
	} else {
		GlobalVariables.BotPersonality = "dinoPersonality.txt"
		send, err := session.ChannelMessageSend(message.ChannelID, "Could not find stored personality for: "+personality+". Reverting to Dinofault.")
		if err != nil {
			log.Fatal(err)
			return
		}
		log.Println("Correctly sent: ", send)
	}
}

func isValidPersonality(personality string) bool {
	return strings.Contains(strings.ToLower(personality), "mark") || strings.Contains(strings.ToLower(personality), "matt") || strings.Contains(strings.ToLower(personality), "chris") || strings.Contains(strings.ToLower(personality), "tommy") || strings.Contains(strings.ToLower(personality), "gerrit") || strings.Contains(strings.ToLower(personality), "dino")
}
