package handler

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/christiansoetanto/servant-of-servus-dei/src/config"
	"github.com/christiansoetanto/servant-of-servus-dei/src/util"
	"log"
	"regexp"
	"strings"
)

//TODO list, later:
//1. move rq to rd
//2. is X a sin

func MessageReactionAddHandler(s *discordgo.Session, m *discordgo.MessageReactionAdd) {

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example, but it's a good practice.
	if m.Member.User.ID == s.State.User.ID {
		return
	}

	guildId := m.GuildID

	if m.Emoji.ID != config.Config[guildId].Reaction.Upvote && m.Emoji.ID != config.Config[guildId].Reaction.Sin {
		return
	}

	if _, ok := config.Moderator[config.ModeratorUserId(m.UserID)]; !ok {
		return
	}
	ChannelId := m.ChannelID
	if ChannelId != config.Config[guildId].Channel.ReligiousQuestions {
		return
	}
	messageId := m.MessageID
	message, err := s.ChannelMessage(config.Config[guildId].Channel.ReligiousQuestions, messageId)
	if err != nil {
		log.Println(err)
		return
	}
	question, questionAsker := message.Content, message.Author.ID

	zzz, err := getMessageReactions(s, guildId, m.ChannelID, messageId)
	if err != nil {
		log.Println(err)
		return
	}
	_ = zzz
	zzz, err = crawlReligiousDiscussionChannel(s, zzz, guildId, question)
	if err != nil {
		log.Println(err)
		return
	}

	msgUrl, err := sendAnswerEmbed(s, zzz, guildId, question, questionAsker)
	if err != nil {
		log.Println(err)
		return
	}

	err = s.ChannelMessageDelete(m.ChannelID, messageId)
	if err != nil {
		log.Println(err)
		return
	}

	log.Printf("[%s] : [%s] | [%s]", "Answered Question", config.Moderator[config.ModeratorUserId(m.UserID)], msgUrl)

}

func getMessageReactions(s *discordgo.Session, guildId, ChannelId, messageId string) ([]z, error) {
	rd1Users, err := s.MessageReactions(ChannelId, messageId, config.ReligiousDiscussions1WhiteCheckMarkEmojiName, 0, "", "")
	if err != nil {
		return nil, err
	}
	rd2Users, err := s.MessageReactions(ChannelId, messageId, config.ReligiousDiscussions2BallotBoxWithCheckEmojiName, 0, "", "")
	if err != nil {
		return nil, err
	}

	//init the answer map. do this because i want to attach the member even if the answer url is not found...
	rd1Answers := make(map[userId]answerUrl)
	rd2Answers := make(map[userId]answerUrl)
	for _, user := range rd1Users {
		rd1Answers[userId(user.ID)] = ""
	}
	for _, user := range rd2Users {
		rd2Answers[userId(user.ID)] = ""
	}

	res := []z{
		{
			religiousDiscussionChannelId: config.Config[guildId].Channel.ReligiousDiscussions1,
			religiousDiscussionEmoji:     config.ReligiousDiscussions1WhiteCheckMarkEmojiName,
			users:                        rd1Users,
			answer:                       rd1Answers,
		},
		{
			religiousDiscussionChannelId: config.Config[guildId].Channel.ReligiousDiscussions2,
			religiousDiscussionEmoji:     config.ReligiousDiscussions2BallotBoxWithCheckEmojiName,
			users:                        rd2Users,
			answer:                       rd2Answers,
		},
	}
	return res, nil
}

func crawlReligiousDiscussionChannel(s *discordgo.Session, zzz []z, guildId, question string) ([]z, error) {
	for _, z2 := range zzz {
		if len(z2.users) == 0 {
			continue
		}
		lastMessageId := ""
		totalAnswerToBeFound := len(z2.users)
		for i := 0; i < MaxMessageAmount/LimitPerRequest && totalAnswerToBeFound > 0; i++ {
			fmt.Printf("current iter: %d, max iter: %d, answer left: %d\n", i, MaxMessageAmount/LimitPerRequest, totalAnswerToBeFound)
			messages, err := s.ChannelMessages(z2.religiousDiscussionChannelId, LimitPerRequest, lastMessageId, "", "")
			if err != nil {
				return nil, err
			}
			lastMessageId = messages[len(messages)-1].ID

			for _, answerToBe := range messages {
				isUserValid := false
				for _, user := range z2.users {
					if answerToBe.Author.ID == user.ID {
						isUserValid = true
						break
					}
				}
				if !isUserValid {
					continue
				}
				if !strings.Contains(sanitize(answerToBe.Content), sanitize(question)) {
					continue
				}
				answerLink := fmt.Sprintf("https://discord.com/channels/%s/%s/%s", guildId, z2.religiousDiscussionChannelId, answerToBe.ID)
				z2.answer[userId(answerToBe.Author.ID)] = answerUrl(answerLink)
				totalAnswerToBeFound -= 1
				if totalAnswerToBeFound <= 0 {
					break
				}
			}

		}
	}
	return zzz, nil
}

func sendAnswerEmbed(s *discordgo.Session, zzz []z, guildId, question, asker string) (string, error) {
	description := fmt.Sprintf("Question by <@%s>:\n %s", asker, question)

	var fields []*discordgo.MessageEmbedField
	value := ""
	for _, z2 := range zzz {

		if len(z2.answer) == 0 {
			continue
		}

		value += fmt.Sprintf("\nin <#%s> by:\n", z2.religiousDiscussionChannelId)
		for id, url := range z2.answer {
			value += fmt.Sprintf("‣ <@%s>", id)
			if url != "" {
				value += fmt.Sprintf(" [jump to answer!](%s)", url)
			}
			value += "\n"
		}
	}
	fields = append(fields, &discordgo.MessageEmbedField{
		Name:   "Answer(s): ",
		Value:  value,
		Inline: false,
	})
	embed := &discordgo.MessageEmbed{
		URL:         util.ServusDeiWebsiteURL,
		Type:        discordgo.EmbedTypeRich,
		Description: description,
		Color:       util.GoldenYellowColor,
		Footer: &discordgo.MessageEmbedFooter{
			Text:    util.FooterText,
			IconURL: util.LogoURL,
		},
		Fields: fields,
	}
	sentAnsweredQuestion, err := s.ChannelMessageSendEmbed(config.Config[guildId].Channel.AnsweredQuestions, embed)
	if err != nil {
		return "", err
	}
	answerLink := fmt.Sprintf("https://discord.com/channels/%s/%s/%s", guildId, config.Config[guildId].Channel.AnsweredQuestions, sentAnsweredQuestion.ID)
	return answerLink, nil
}

const (
	LimitPerRequest  = 100
	MaxMessageAmount = 1000
)

func sanitize(input string) string {
	return strings.ToLower(regexp.MustCompile(`/[\W_]+/g`).ReplaceAllString(input, ""))
}

type userId string
type answerUrl string
type z struct {
	religiousDiscussionChannelId string
	religiousDiscussionEmoji     string
	users                        []*discordgo.User
	answer                       map[userId]answerUrl
}
