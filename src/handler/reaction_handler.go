package handler

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/christiansoetanto/servant-of-servus-dei/src/config"
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

	channelId := m.ChannelID
	if channelId != config.Config[guildId].Channel.ReligiousQuestions {
		return
	}
	//what i have done: do all validation, already able to catch the add upvote reaction handler in RQ
	//todo next:
	//1. get apakah dia di rd1 atau rd2, simpan map user yang menjawab.
	//2. crawl rd1 and/or rd2, di setiap message check message.user (yang kirim messgagenya),
	//apakah ada di map user yang menjawab tsb.
	//kalau ada baru cek contentnya apakah sama (jangan lupa sanitize, cek code lama),
	//if all match, get the message link. simpan balik ke map tadi aja,
	//jadi map tadi isinya map[userid]interface_answer yang isinya adalah link jawaban, user, dan channel.
	//tp channel ini lebih enak kalo map aslinya dipisah per channel sih.
	//3. setelah selesai crawl, send embed ke answered question

	messageId := m.MessageID

	rd1Users, err := s.MessageReactions(m.ChannelID, messageId, config.ReligiousDiscussions1WhiteCheckMarkEmojiName, 0, "", "")
	//rd2Users, err := s.MessageReactions(m.ChannelID, messageId, util.ReligiousDiscussions2BallotBoxWithCheckEmojiName, 0, "", "")

	message, err := s.ChannelMessage(config.Config[guildId].Channel.ReligiousQuestions, messageId)
	if err != nil {
		return
	}
	content, questionAsker, reactions := message.Content, message.Member, message.Reactions

	rd1 := make(aaa)
	rd2 := make(aaa)

	if len(rd1Users) > 0 {

		//di sini bikin algs utk crawl channelnya. cek logic di code lama.
		messages, err := s.ChannelMessages(config.Config[guildId].Channel.ReligiousDiscussions1, 100, "", "", "")
		_, _ = messages, err
	}

	for _, reaction := range reactions {
		emoji := reaction.Emoji

		if emoji.Name != config.ReligiousDiscussions1WhiteCheckMarkEmojiName && emoji.Name != config.ReligiousDiscussions2BallotBoxWithCheckEmojiName {
			continue
		}

		//add to rd1 map
		if emoji.Name == config.ReligiousDiscussions1WhiteCheckMarkEmojiName {

			rd1[emoji.User.ID] = answerData{
				user: emoji.User,
			}

		}
		if emoji.Name == config.ReligiousDiscussions2BallotBoxWithCheckEmojiName {
			rd2[emoji.User.ID] = answerData{
				user: emoji.User,
			}
		}

	}

	fmt.Println(m, content, questionAsker, reactions, rd1, rd2)
}

type aaa map[string]answerData

type answerData struct {
	channelId string
	answerUrl string
	user      *discordgo.User
}
