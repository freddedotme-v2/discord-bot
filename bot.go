package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func main() {
	discord, err := discordgo.New("Bot " + os.Getenv("DISCORD_TOKEN"))
	if err != nil {
		fmt.Println(err)
		return
	}

	discord.AddHandler(ready)
	discord.AddHandler(messageCreate)

	err = discord.Open()
	if err != nil {
		fmt.Println(err)
	}
}

func ready(s *discordgo.Session, event *discordgo.Ready) {
	s.UpdateGameStatus(0, "!8ball")
}

func messageCreate(s *discordgo.Session, event *discordgo.MessageCreate) {
	// Ignore if the author is the bot, to avoid infinite loop.
	if event.Author.ID == s.State.User.ID {
		return
	}

	// Discard if it doesn't have the correct prefix.
	if !strings.HasPrefix(event.Content, "!8ball ") {
		return
	}

	answers := []string{
		"It is certain",
		"It is decidedly so",
		"Without a doubt",
		"Yes definitely",
		"You may rely on it",
		"As I see it, yes",
		"Most likely",
		"Outlook good",
		"Yes",
		"Signs point to yes",
		"Reply hazy, try again",
		"Ask again later",
		"Better not tell you now",
		"Cannot predict now",
		"Concentrate and ask again",
		"Don't count on it",
		"My reply is no",
		"My sources say no",
		"Outlook not so good",
		"Very doubtful",
	}

	_, err := s.ChannelMessageSend(event.ChannelID, answers[rand.Intn(len(answers))])
	if err != nil {
		fmt.Printf("Error while sending: %s\n", err)
		return
	}
}
