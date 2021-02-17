package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/rs/zerolog/log"
)

const (
	Token = "write token here"
)

func main() {
	fmt.Println("I was born to be reminding 勉強会!")

	discord, err := discordgo.New()
	discord.Token = Token
	if err != nil {
		log.Fatal().Msgf("Error logging in: %v", err)
	}
}
