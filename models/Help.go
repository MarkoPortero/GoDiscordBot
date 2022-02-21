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

	session.ChannelMessageSend(message.ChannelID, string(content))
}

func NsfwHelp(session *discordgo.Session, message *discordgo.MessageCreate) {
	content, err := ioutil.ReadFile("./resources/nsfwhelp.txt")

	if err != nil {
		log.Fatal(err)
	}

	session.ChannelMessageSend(message.ChannelID, "```"+string(content)+"```")

	content, err = ioutil.ReadFile("./resources/nsfwhelp2.txt")

	if err != nil {
		log.Fatal(err)
	}

	session.ChannelMessageSend(message.ChannelID, "```"+string(content)+"```")
}
