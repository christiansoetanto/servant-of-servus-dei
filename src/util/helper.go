package util

import (
	"github.com/bwmarrin/discordgo"
	"log"
	"regexp"
	"strings"
)

func Sanitize(input string) string {
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		log.Println(err)
		return input
	}
	return strings.ToLower(reg.ReplaceAllString(input, ""))
}

func ReportError(s *discordgo.Session, msg string) {
	channel, err := s.UserChannelCreate("255514888041005057")
	if err != nil {
		log.Print(err.Error())
		return
	}
	_, err = s.ChannelMessageSend(channel.ID, msg)
	if err != nil {
		log.Print(err.Error())
		return
	}
	log.Print(msg)
}
