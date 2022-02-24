package models

import (
	"context"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/vartanbeno/go-reddit/v2/reddit"
	"log"
	"strings"
)

func GetRedditPost(session *discordgo.Session, message *discordgo.MessageCreate) {
	breakDown := strings.Split(strings.ToLower(message.Content), "reddit")
	fmt.Println("searching for: https://www.reddit.com/r/" + strings.TrimSpace(breakDown[1]) + "/random.json")

	client, _ := reddit.NewReadonlyClient()
	posts, _, err := client.Post.RandomFromSubreddits(context.Background(), strings.TrimSpace(breakDown[1]))
	if err != nil {
		fmt.Println(err)
	}
	if posts == nil {
		send, err := session.ChannelMessageSend(message.ChannelID, "Subreddit not found or reddits being a wanker.")
		if err != nil {
			log.Fatal(err)
			return
		}
		log.Println("Correctly sent: ", send)
		return
	}
	if len(posts.Post.Permalink) > 0 {
		send, err := session.ChannelMessageSend(message.ChannelID, "https://www.reddit.com"+posts.Post.Permalink)
		if err != nil {
			log.Fatal(err)
		}
		log.Println("Correctly sent: ", send)
	}
	if len(posts.Post.Title) > 0 {
		send, err := session.ChannelMessageSend(message.ChannelID, posts.Post.Title)
		if err != nil {
			log.Fatal(err)
		}
		log.Println("Correctly sent: ", send)
	}
	if len(posts.Post.Body) > 0 {
		send, err := session.ChannelMessageSend(message.ChannelID, posts.Post.Body)
		if err != nil {
			log.Fatal(err)
		}
		log.Println("Correctly sent: ", send)
	}
	if len(posts.Post.URL) > 0 {
		send, err := session.ChannelMessageSend(message.ChannelID, posts.Post.URL)
		if err != nil {
			log.Fatal(err)
		}
		log.Println("Correctly sent: ", send)
	}
}
