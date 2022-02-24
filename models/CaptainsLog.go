package models

import (
	"GoDiscordBot/DataAccess"
	"github.com/bwmarrin/discordgo"
)

func StoreCaptainsLogs(session *discordgo.Session, message *discordgo.MessageCreate) {
	DataAccess.MongoDbStoreCaptainsLogInDatabase(message)
	session.ChannelMessageSend(message.ChannelID, "Aye aye Captain, log received.")
}

func ReadCaptainsLogs(session *discordgo.Session, message *discordgo.MessageCreate) {
	captainsLog := DataAccess.MongoDbReadCaptainsLogInDatabase(message)
	session.ChannelMessageSend(message.ChannelID, captainsLog)
}
