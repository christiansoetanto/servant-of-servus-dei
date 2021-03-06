package handler

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/christiansoetanto/servant-of-servus-dei/src/config"
	"github.com/christiansoetanto/servant-of-servus-dei/src/util"
	"log"
	"strings"
)

var questionOneString = util.Sanitize("and give us the code")

var INRI = util.Sanitize("INRI")

func MessageCreateHandlerQuestionOne(s *discordgo.Session, m *discordgo.MessageCreate) {

	if m.Author.ID == s.State.User.ID {
		return
	}
	guildId := m.GuildID
	guildConfig, ok := config.Config[guildId]
	if !ok {
		return
	}
	if m.ChannelID != guildConfig.Channel.Responses {
		return
	}

	sanitizedContent := util.Sanitize(m.Content)
	//lol apparently Latin Rite contains INRI so i have to remove it first before detecting the real "INRI"
	sanitizedContent = strings.ReplaceAll(sanitizedContent, "latinrite", "")
	if strings.Contains(sanitizedContent, questionOneString) && !strings.Contains(sanitizedContent, INRI) {
		userId := m.Author.ID

		_, err := s.ChannelMessageSend(guildConfig.Channel.Responses, fmt.Sprintf("Hey <@%s>! It looks like you missed question 1. Please re-read the <#%s> again, we assure you that the code is in there. Thank you for your understanding.\nPS: if you are sure you got it right, please ignore this message.", userId, guildConfig.Channel.RulesVetting))
		if err != nil {
			log.Println(err)
			return
		}
	}

}
