package models

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

var (
	WhatImReminding string
	TimeToRemind    string
)

func RemindMe(session *discordgo.Session, message *discordgo.MessageCreate, remind string, ping string) {

	messageContent := message.Content
	breakDown := strings.Split(strings.ToLower(messageContent), "remind me to ")
	if len(breakDown) > 1 {
		TimeToRemind = breakDown[0]
		WhatImReminding = breakDown[1]
		fmt.Println(remind)
		send, err := session.ChannelMessageSend(message.ChannelID, "Certainly. I'll remind you to "+remind+", young Padawan.")
		if err != nil {
			log.Fatal(err)
		}
		log.Println("Correctly sent: ", send)
	}
	if strings.Contains(strings.ToLower(TimeToRemind), "minute") {
		timer1 := time.NewTimer(time.Minute * 1)
		<-timer1.C
		send, err := session.ChannelMessageSend(message.ChannelID, "Oi, "+ping+", i'm reminding you to "+remind)
		if err != nil {
			log.Fatal(err)
		}
		log.Println("Correctly sent: ", send)
	}
	if strings.Contains(strings.ToLower(TimeToRemind), "hour") {
		timer1 := time.NewTimer(time.Hour * 1)
		<-timer1.C
		send, err := session.ChannelMessageSend(message.ChannelID, "Oi, "+ping+", i'm reminding you to "+remind)
		if err != nil {
			log.Fatal(err)
		}
		log.Println("Correctly sent: ", send)
	}
}
