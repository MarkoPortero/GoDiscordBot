package handlers

import (
	"GoDiscordBot/models"
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
)

var (
	BotID   string
	Place   string
	Country string
)

func MessageHandler(session *discordgo.Session, message *discordgo.MessageCreate) {
	if message.Author.ID == BotID {
		return
	}

	fmt.Println(message.Author)
	fmt.Println(message.Content)

	if message.Content == "ping" {
		if message.Author.String() == "cactusapple#2171" {
			session.ChannelMessageSend(message.ChannelID, "pongers.")
			return
		}
		_, _ = session.ChannelMessageSend(message.ChannelID, "hello "+message.Author.Mention())
		return
	}

	if strings.HasPrefix(getMessageContents(message), "!") {
		if doesMessageContain(message, "spells") {
			models.GetSpells(session, message)
			return
		}
		if doesMessageContain(message, "news") {
			models.GetNews(session, message)
			return
		}
		if doesMessageContain(message, "next") {
			models.NextNewsArticle(session, message)
			return
		}
		if doesMessageContain(message, "remind me") {
			fmt.Println(message.Author.ID)
			ToRemind := message.Author.ID
			ToRemindPing := message.Author.Mention()
			models.RemindMe(session, message, ToRemind, ToRemindPing)
			return
		}
	}
	if doesMessageContain(message, "hello") {
		session.ChannelMessageSend(message.ChannelID, "https://giphy.com/gifs/mrw-top-escalator-Nx0rz3jtxtEre")
		return
	}

	if doesMessageContain(message, "weather in") {
		breakDown := strings.Split(strings.ToLower(message.Content), "weather in ")
		if len(breakDown) > 1 {
			Place = breakDown[1]
			models.WeatherApiConsumer(session, message, Place)
		}
		return
	}

	if doesMessageContain(message, "dog") {
		models.DoggyStyle(session, message)
		return
	}

	if doesMessageContain(message, "coronavirus in") {
		breakDown := strings.Split(strings.ToLower(message.Content), "coronavirus in")
		if len(breakDown) > 1 {
			Country = breakDown[1]
			models.CoronavirusStats(session, message, Country)
		}
		return
	}

	if doesMessageContain(message, "bored") {
		models.Bored(session, message)
		return
	}
}

// func watchYoProfanity(session *discordgo.Session, message *discordgo.MessageCreate) {
// 	response, err := http.Get("https://www.purgomalum.com/service/containsprofanity?text=" + message.Content)
// 	if err != nil {
// 		fmt.Printf("The HTTP request failed with error %s\n", err)
// 	} else {
// 		bodyBytes, err := ioutil.ReadAll(response.Body)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		body := string(bodyBytes)
// 		if strings.Contains(body, "true") {
// 			session.ChannelMessageSend(message.ChannelID, "https://tenor.com/view/watch-your-profanity-funny-gif-5600117")
// 		}
// 		fmt.Println(body)
// 	}
// }

func doesMessageContain(message *discordgo.MessageCreate, containingWord string) (DoesContain bool) {
	DoesContain = strings.Contains(strings.ToLower(message.Content), containingWord)
	return
}

func getMessageContents(message *discordgo.MessageCreate) (Message string) {
	Message = strings.ToLower(message.Content)
	return
}
