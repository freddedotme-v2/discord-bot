package main

import (
	"fmt"
	"os"

	"github.com/bwmarrin/discordgo"
)

func main() {
	discord, err := discordgo.New("Bot " + os.Getenv("DISCORD_TOKEN"))
	if err != nil {
		fmt.Println(err)
		return
	}

	discord.AddHandler(ready)

	err = discord.Open()
	if err != nil {
		fmt.Println(err)
	}
}

func ready(s *discordgo.Session, event *discordgo.Ready) {
	s.UpdateGameStatus(0, "relaxing")
}
