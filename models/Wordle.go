package models

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"io/ioutil"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

var (
	word          string
	guess         string
	attempts      map[string]int
	attemptResult string
)

func getWord() string {
	var (
		word string
		i    int
	)
	words, err := ioutil.ReadFile("./resources/words")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	// split the file into words
	buff := strings.Split(string(words), "\n")
	// get a random word
	// seed the random number generator
	rand.Seed(time.Now().UnixNano())
	i = rand.Intn(len(buff))
	word = buff[i]
	// check if the word is 5 letters long
	if len(word) != 5 {
		fmt.Println("Error: the file must contain 5 letters words")
		os.Exit(1)
	}
	return word
}

func Wordle(session *discordgo.Session, message *discordgo.MessageCreate) {
	word = getWord()
	fmt.Println("word is: " + word)
	attempts = make(map[string]int)
	session.ChannelMessageSend(message.ChannelID, "Game begins. Use !guess to make a guess.")
}

func Guess(session *discordgo.Session, message *discordgo.MessageCreate) {
	if attempts[message.Author.ID] >= 5 {
		session.ChannelMessageSend(message.ChannelID, "Sorry, you're out of guesses.")
	}

	breakDown := strings.Split(strings.ToLower(message.Content), "!guess")
	if len(breakDown) > 1 {
		guess = breakDown[1]
	}

	if len(guess) != 5 {
		session.ChannelMessageSend(message.ChannelID, "Must be a 5 letter word.")
		return
	}
	attemptResult = ""
	for i := 0; i < len(word); i++ {
		if word[i] == guess[i] { // correct at the right position
			attemptResult = ":green_square: " + string(guess[i])
		} else if strings.ContainsAny(word, string(guess[i])) { // correct at the wrong position
			attemptResult = ":yellow_square: " + string(guess[i])
		} else { // incorrect
			attemptResult = ":black_large_square:" + string(guess[i])
		}
	}
	session.ChannelMessageSend(message.ChannelID, attemptResult)
	attempts[message.Author.ID] += 1
	session.ChannelMessageSend(message.ChannelID, "You have used "+strconv.Itoa(attempts[message.Author.ID]))
}

func Game() {
	var (
		word  string = getWord()
		guess string
		count int  = 0
		end   bool = true
	)
	for end {
		fmt.Print("> ")
		fmt.Scanf("%s", &guess)
		if len(guess) != 5 {
			fmt.Println("Please enter 5 letters")
			continue
		}
		// check if the guess is correct showing colors per letter
		for i := 0; i < len(word); i++ {
			if word[i] == guess[i] { // correct at the right position
				fmt.Print("\033[32m", string(guess[i]), "\033[0m")
			} else if strings.ContainsAny(word, string(guess[i])) { // correct at the wrong position
				fmt.Print("\033[34m", string(guess[i]), "\033[0m")
			} else { // incorrect
				fmt.Print("\033[31m", string(guess[i]), "\033[0m")
			}
		}
		fmt.Println()
		// check if the whole word is guessed
		if word == guess {
			fmt.Println("\033[32m" + "You won!" + "\033[0m")
			end = false
		}
		// check if the user has 5 guesses
		if count == 5 {
			fmt.Println("\033[31m" + "You lost!" + "\033[0m")
			fmt.Println("The word was:", word)
			end = false
		}
		count++
	}
}
