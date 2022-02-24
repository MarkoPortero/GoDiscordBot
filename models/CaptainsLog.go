package models

import (
	"GoDiscordBot/DataAccess"
	"github.com/bwmarrin/discordgo"
	"log"
)

func StoreCaptainsLogs(session *discordgo.Session, message *discordgo.MessageCreate) {
	DataAccess.MongoDbStoreCaptainsLogInDatabase(message)
	send, err := session.ChannelMessageSend(message.ChannelID, "Aye aye Captain, log received.")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Correctly sent: ", send)
}

func ReadCaptainsLogs(session *discordgo.Session, message *discordgo.MessageCreate) {
	captainsLog := DataAccess.MongoDbReadCaptainsLogInDatabase(message)
	send, err := session.ChannelMessageSend(message.ChannelID, captainsLog)
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Println("Correctly sent: ", send)
}
