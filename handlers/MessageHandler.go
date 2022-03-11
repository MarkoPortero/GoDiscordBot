package handlers

import (
	"GoDiscordBot/GlobalVariables"
	"GoDiscordBot/models"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"

	"strings"

	"github.com/bwmarrin/discordgo"
)

var (
	Place   string
	Country string
)

func MessageHandler(session *discordgo.Session, message *discordgo.MessageCreate) {
	if message.Author.ID == GlobalVariables.BotID {
		return
	}

	fmt.Println("AUTHOR ID " + message.Author.ID)
	fmt.Println("BOT ID " + GlobalVariables.BotID)
	fmt.Println(message.Author)
	fmt.Println(message.Content)

	if message.Content == "ping" {
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
		if doesMessageContain(message, "help") {
			models.HelpCommand(session, message)
			return
		}

		if doesMessageContain(message, "nsfwlists") {
			models.NsfwHelp(session, message)
			return
		}
		if doesMessageContain(message, "weather") {
			breakDown := strings.Split(strings.ToLower(message.Content), "weather")
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

		if doesMessageContain(message, "personality") {
			breakDown := strings.Split(strings.ToLower(message.Content), "personality")
			if len(breakDown) > 1 {
				personality := breakDown[1]
				fmt.Println("attempting to set personality: " + personality)
				models.PersonalityTransfer(session, message, personality)
			}
			return
		}

		if doesMessageContain(message, "schedulefactson") {
			breakDown := strings.Split(strings.ToLower(message.Content), "schedulefactson")
			if len(breakDown) > 1 {
				go func() {
					GlobalVariables.ScheduledFactStatus = true
					for GlobalVariables.ScheduledFactStatus {
						models.GetFact(session, message)
						time.Sleep(5 * time.Minute)
					}
				}()
			}
			return
		}

		if doesMessageContain(message, "schedulefactsoff") {
			GlobalVariables.ScheduledFactStatus = false
			return
		}

		if doesMessageContain(message, "scheduleon") {
			go func() {
				for GlobalVariables.ScheduleMessageStatus {
					GlobalVariables.ScheduleMessageStatus = true
					impersonateSomeone(session, message)
					time.Sleep(5 * time.Minute)
				}
			}()
			return
		}

		if doesMessageContain(message, "ScheduleOff") {
			GlobalVariables.ScheduleMessageStatus = false
			session.ChannelMessageSend(message.ChannelID, "alright then")
			return
		}

		if doesMessageContain(message, "fact") {
			models.GetFact(session, message)
			return
		}

		if doesMessageContain(message, "reddit") {
			models.GetRedditPost(session, message)
			return
		}

		if doesMessageContain(message, "captainslog") {
			models.StoreCaptainsLogs(session, message)
			return
		}

		if doesMessageContain(message, "readmylog") {
			models.ReadCaptainsLogs(session, message)
			return
		}

		if doesMessageContain(message, "envVariables") {
			if len(os.Environ()) <= 0 {
				session.ChannelMessageSend(message.ChannelID, "Sorry, couldn't find any environment variables.")
			}
			for _, env := range os.Environ() {
				// env is
				envPair := strings.SplitN(env, "=", 2)
				key := envPair[0]
				value := envPair[1]

				fmt.Printf("%s : %s\n", key, value)
				session.ChannelMessageSend(message.ChannelID, key+" : "+value)
			}
		}

		if doesMessageContain(message, "wordle") {
			models.Wordle(session, message)
			return
		}

		if doesMessageContain(message, "guess") {
			models.Wordle(session, message)
		}
	}

	if doesMessageContain(message, "hello") {
		session.ChannelMessageSend(message.ChannelID, "https://giphy.com/gifs/mrw-top-escalator-Nx0rz3jtxtEre")
		return
	}

	if doesMessageContain(message, "dino") || doesMessageContain(message, "mark") || doesMessageContain(message, "tommy") || doesMessageContain(message, "matt") || doesMessageContain(message, "chris") || doesMessageContain(message, "gerrit") {
		impersonateSomeone(session, message)
		return
	}
}

func impersonateSomeone(session *discordgo.Session, message *discordgo.MessageCreate) {
	min := 1
	max := 6
	linesToPrint := rand.Intn(max-min) + min
	fmt.Println("lines: " + strconv.Itoa(linesToPrint))
	for i := 0; i < linesToPrint; i++ {
		models.Impersonate(session, message)
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
