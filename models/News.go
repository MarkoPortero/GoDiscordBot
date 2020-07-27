package models

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

var (
	currentIterator int
	news            newsArticle
)

type newsArticle struct {
	Status       string `json:"status"`
	TotalResults int    `json:"totalResults"`
	Articles     []struct {
		Source struct {
			ID   interface{} `json:"id"`
			Name string      `json:"name"`
		} `json:"source"`
		Author      string      `json:"author"`
		Title       string      `json:"title"`
		Description string      `json:"description"`
		URL         string      `json:"url"`
		URLToImage  string      `json:"urlToImage"`
		PublishedAt time.Time   `json:"publishedAt"`
		Content     interface{} `json:"content"`
	} `json:"articles"`
}

// GetNews gets news article for location
func GetNews(session *discordgo.Session, message *discordgo.MessageCreate) {
	searchCriteria := strings.Split(strings.ToLower(message.Content), "!news")
	if len(searchCriteria) > 1 {
		response, err := http.Get("http://newsapi.org/v2/top-headlines?country=" + strings.Trim(searchCriteria[1], " ") + "&apiKey=323f3e846e014d03a7fe84bcce50016a")
		if err != nil {
			fmt.Printf("The HTTP request failed with error %s\n", err)
		} else {
			if err := json.NewDecoder(response.Body).Decode(&news); err != nil {
				log.Println(err)
			}
			fmt.Println(news)

			if (news.TotalResults) == 0 {
				session.ChannelMessageSend(message.ChannelID, "Sorry, i'm having trouble finding news articles for that criteria.")
				return
			}

			currentIterator = 0
			printNewsArticle(session, message)
		}
	}
}

// NextNewsArticle prints next news article
func NextNewsArticle(session *discordgo.Session, message *discordgo.MessageCreate) {
	currentIterator++
	if news.TotalResults == 0 || currentIterator == news.TotalResults || currentIterator > news.TotalResults {
		session.ChannelMessageSend(message.ChannelID, "Sorry, i'm out of news.")
		return
	}
	printNewsArticle(session, message)
}

func printNewsArticle(session *discordgo.Session, message *discordgo.MessageCreate) {
	session.ChannelMessageSend(message.ChannelID,
		"Showing article: "+fmt.Sprint(currentIterator+1)+" of "+fmt.Sprint(news.TotalResults)+"\n\n"+news.Articles[currentIterator].Title+"\n\n"+news.Articles[currentIterator].Description+"\n\n"+news.Articles[currentIterator].URL)
}
