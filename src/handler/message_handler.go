package handler

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/christiansoetanto/servant-of-servus-dei/src/util"
	"log"
	"strings"
)

func MessageCreateHandler(s *discordgo.Session, m *discordgo.MessageCreate) {

	if m.Author.ID == s.State.User.ID {
		return
	}
	if m.Content == "ping" {
		s.ChannelMessageSend(m.ChannelID, "Pong!")
	}

	if m.Content == "pong" {
		s.ChannelMessageSend(m.ChannelID, "Ping!")
	}
}

var questionOneString = util.Sanitize("and give us the code")

var INRI = util.Sanitize("INRI")

func MessageCreateHandlerQuestionOne(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}
	if m.ChannelID != util.ResponsesChannelId {
		return
	}

	sanitizedContent := util.Sanitize(m.Content)
	//lol apparently Latin Rite contains INRI so i have to remove it first before detecting the real "INRI"
	sanitizedContent = strings.ReplaceAll(sanitizedContent, "latinrite", "")
	if strings.Contains(sanitizedContent, questionOneString) && !strings.Contains(sanitizedContent, INRI) {
		userId := m.Author.ID
		err := s.GuildMemberRoleAdd(util.GuildID, userId, util.VettingQuestioningRoleId)
		if err != nil {
			log.Println(err)
			return
		}
		_, err = s.ChannelMessageSend(util.VettingQuestioningChannelId, fmt.Sprintf("Hey <@%s>! It looks like you missed question 1. Please re-read the <#%s> again, we assure you that the code is in there. Thank you for your understanding.\nPS: if you are sure you got it right, please ignore this message.", userId, util.RulesVettingChannelId))
		if err != nil {
			log.Println(err)
			return
		}
	}

}
