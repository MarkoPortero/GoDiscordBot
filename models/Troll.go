package models

import (
	"github.com/bwmarrin/discordgo"
	"io"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func Goose(session *discordgo.Session, message *discordgo.MessageCreate) {
	f, err := os.Open("./resources/TOTALLYSAFEZIP.zip")
	check(err)

	session.ChannelFileSend(message.ChannelID, "NotVirus.zip", io.Reader(f))
}
