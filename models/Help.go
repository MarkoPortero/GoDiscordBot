package models

import (
	"github.com/bwmarrin/discordgo"
	"io/ioutil"
	"log"
)

func HelpCommand(session *discordgo.Session, message *discordgo.MessageCreate) {
	content, err := ioutil.ReadFile("./resources/help.txt")

	if err != nil {
		log.Fatal(err)
	}

	send, err := session.ChannelMessageSend(message.ChannelID, string(content))
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Correctly sent: ", send)
}

func NsfwHelp(session *discordgo.Session, message *discordgo.MessageCreate) {
	content, err := ioutil.ReadFile("./resources/nsfwhelp.txt")

	if err != nil {
		log.Fatal(err)
	}

	send, err := session.ChannelMessageSend(message.ChannelID, "```"+string(content)+"```")
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Println("Correctly sent: ", send)

	content, err = ioutil.ReadFile("./resources/nsfwhelp2.txt")
	if err != nil {
		log.Fatal(err)
		return
	}

	messageSend, err := session.ChannelMessageSend(message.ChannelID, "```"+string(content)+"```")
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Println("Correctly sent: ", messageSend)
}

func NewsHelp(session *discordgo.Session, message *discordgo.MessageCreate) {
	content, err := ioutil.ReadFile("./resources/newshelp.txt")

	if err != nil {
		log.Fatal(err)
	}

	send, err := session.ChannelMessageSend(message.ChannelID, string(content))
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Correctly sent: ", send)
}
