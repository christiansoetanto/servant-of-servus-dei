package handler

import (
	"errors"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/christiansoetanto/servant-of-servus-dei/src/calendar"
	"github.com/christiansoetanto/servant-of-servus-dei/src/config"
	"github.com/christiansoetanto/servant-of-servus-dei/src/util"
	"log"
)

const (
	Ping          = "ping"
	SDVerify      = "sdverify"
	SDQuestionOne = "sdquestionone"
	Calendar      = "calendar"
	NiceTryBro    = "Nice try, bro! You are not allowed to use this command... <@255514888041005057>"
)

var (
	commandHandlers    = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate) error{}
	commands           []*discordgo.ApplicationCommand
	registeredCommands []*discordgo.ApplicationCommand
)

func InitCommandHandler() {
	commandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate) error{
		Ping: func(s *discordgo.Session, i *discordgo.InteractionCreate) error {
			if !isMod(i.Member.User.ID) {
				err := alertNonMod(s, i)
				if err != nil {
					return err
				}
				return nil
			}
			err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "Pong!",
				},
			})
			if err != nil {
				return err
			}
			return nil
		},
		SDVerify: func(s *discordgo.Session, i *discordgo.InteractionCreate) error {
			if !isMod(i.Member.User.ID) {
				err := alertNonMod(s, i)
				if err != nil {
					return err
				}
				return nil

			}
			err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "Processing... please wait...",
				},
			})
			if err != nil {
				return err
			}
			// Access options in the order provided by the user.
			options := i.ApplicationCommandData().Options
			guildId := i.GuildID
			guildConfig, ok := config.Config[guildId]
			if !ok {
				return errors.New("config not found")
			}

			// Or convert the slice into a map
			optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
			for _, opt := range options {
				optionMap[opt.Name] = opt
			}

			acknowledgementMessageArgs := make([]interface{}, 0, len(options))
			acknowledgementMessageFormat := guildConfig.Wording.AcknowledgementMessageFormat
			welcomeMessageArgs := make([]interface{}, 0, 1)

			var user *discordgo.User

			userOpt, userOptOk := optionMap["user-option"]
			roleOpt, roleOptOk := optionMap["role-option"]
			var roleType string
			if userOptOk && roleOptOk {
				user = userOpt.UserValue(s)
				acknowledgementMessageArgs = append(acknowledgementMessageArgs, user.ID)
				welcomeMessageArgs = append(welcomeMessageArgs, user.ID)
				welcomeMessageArgs = append(welcomeMessageArgs, guildConfig.Channel.ReactionRoles)
				welcomeMessageArgs = append(welcomeMessageArgs, guildConfig.Channel.ServerInformation)

				//actually i dont need to put this in here, because user is required anyway. but just to be safe haha
				roleType = roleOpt.StringValue()
				roleId := guildConfig.ReligionRoleMappingMap[config.ReligionRoleType(roleType)]
				acknowledgementMessageArgs = append(acknowledgementMessageArgs, roleId)

				err := s.GuildMemberRoleAdd(guildId, user.ID, string(roleId))
				if err != nil {
					fmt.Println(err)
					return err
				}

				err = s.GuildMemberRoleAdd(guildId, user.ID, guildConfig.Role.ApprovedUser)
				if err != nil {
					return err
				}
				err = s.GuildMemberRoleRemove(guildId, user.ID, guildConfig.Role.Vetting)
				if err != nil {
					return err
				}
				err = s.GuildMemberRoleRemove(guildId, user.ID, guildConfig.Role.VettingQuestioning)
				if err != nil {
					return err
				}

			} else {
				_, err := s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
					Content: "Please choose user and role.",
				})
				if err != nil {
					return err
				}
				return nil
			}

			mod := i.Member
			content := fmt.Sprintf(guildConfig.Wording.WelcomeMessageFormat, user.Mention(), mod.Mention())
			_, err = s.ChannelMessageSend(guildConfig.Channel.GeneralDiscussion, content)
			if err != nil {
				return err
			}

			_, err = s.ChannelMessageSendEmbed(guildConfig.Channel.GeneralDiscussion, util.EmbedBuilder(guildConfig.Wording.WelcomeTitle, fmt.Sprintf(guildConfig.Wording.WelcomeMessageEmbedFormat, welcomeMessageArgs...), util.RandomWelcomeImage()))
			if err != nil {
				return err
			}
			_, err = s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
				Content: fmt.Sprintf(
					acknowledgementMessageFormat,
					acknowledgementMessageArgs...,
				),
			})

			if err != nil {
				return err
			}

			log.Printf("[%s] : [%s] | [%s] | [%s]", SDVerify, mod.User.Username, user.Username, roleType)
			return nil
		},
		SDQuestionOne: func(s *discordgo.Session, i *discordgo.InteractionCreate) error {

			if !isMod(i.Member.User.ID) {
				err := alertNonMod(s, i)
				if err != nil {
					return err
				}
				return nil
			}

			err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "Processing... please wait...",
				},
			})
			if err != nil {
				return err
			}
			// Access options in the order provided by the user.
			options := i.ApplicationCommandData().Options
			guildId := i.GuildID
			guildConfig, ok := config.Config[guildId]
			if !ok {
				return errors.New("config not found")
			}

			// Or convert the slice into a map
			optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
			for _, opt := range options {
				optionMap[opt.Name] = opt
			}

			missedQuestionOneMessageFormatArgs := make([]interface{}, 0)
			missedQuestionOneMessageFormat := guildConfig.Wording.MissedQuestionOneFormatNoPS

			var user *discordgo.User

			userOpt, userOptOk := optionMap["user-option"]
			if userOptOk {
				user = userOpt.UserValue(s)
				missedQuestionOneMessageFormatArgs = append(missedQuestionOneMessageFormatArgs, user.ID)
				missedQuestionOneMessageFormatArgs = append(missedQuestionOneMessageFormatArgs, guildConfig.Channel.RulesVetting)
				err := s.GuildMemberRoleAdd(guildId, user.ID, guildConfig.Role.VettingQuestioning)
				if err != nil {
					fmt.Println(err)
					return err
				}
			} else {
				_, err := s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
					Content: "Please choose user.",
				})
				if err != nil {
					return err
				}
				return nil
			}

			mod := i.Member

			_, err = s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
				Content: fmt.Sprintf(
					missedQuestionOneMessageFormat,
					missedQuestionOneMessageFormatArgs...,
				),
			})

			if err != nil {
				return err
			}

			log.Printf("[%s] : [%s] | [%s]", SDQuestionOne, mod.User.Username, user.Username)
			return nil
		},
		Calendar: func(s *discordgo.Session, i *discordgo.InteractionCreate) error {
			err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "Processing... please wait...",
				},
			})
			if err != nil {
				return err
			}

			calendarText, msg := calendar.GetCalendarText()
			if msg != "" {
				util.ReportError(s, msg)
				return errors.New(msg)
			}
			_, err = s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
				Content: calendarText,
			})

			if err != nil {
				util.ReportError(s, err.Error())
				return err
			}

			log.Printf("[%s] : [%s] ", Calendar, i.Interaction.Member.User.Username)
			return nil
		},
	}
	commands = []*discordgo.ApplicationCommand{
		{
			Name:        Ping,
			Description: "Ping",
		},
		{
			Name:        SDVerify,
			Description: "Command for verifying new peeps and welcoming them",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionUser,
					Name:        "user-option",
					Description: "User to verify",
					Required:    true,
				},
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "role-option",
					Description: "Religion role to give",
					Required:    true,
					Choices:     buildReligionRoleOptionChoices(),
				},
			},
		},
		{
			Name:        SDQuestionOne,
			Description: "Command for alerting peeps that they missed question one code",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionUser,
					Name:        "user-option",
					Description: "User to alert",
					Required:    true,
				},
			},
		},
		{
			Name:        Calendar,
			Description: "Get today's liturgical calendar",
		},
	}
}

func InteractionCreateHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if h, ok := commandHandlers[i.ApplicationCommandData().Name]; ok {
		err := h(s, i)
		if err != nil {
			util.ReportError(s, err.Error())
		}
	}
}

func RegisterCommand(s *discordgo.Session) (*discordgo.Session, error) {
	log.Println("Adding commands...")
	for _, v := range commands {
		for guildId, _ := range config.Config {
			cmd, err := s.ApplicationCommandCreate(s.State.User.ID, guildId, v)
			if err != nil {
				log.Fatalf("Cannot create '%v' command: %v", v.Name, err)
				return s, err
			}
			registeredCommands = append(registeredCommands, cmd)
		}

	}

	return s, nil

}

func RemoveCommand(s *discordgo.Session) error {
	log.Println("Removing commands...")

	registeredCommandsToDelete, err := s.ApplicationCommands(s.State.User.ID, "")
	registeredCommandsToDelete1, err := s.ApplicationCommands(s.State.User.ID, config.ServusDeiConfigGuildID)
	registeredCommandsToDelete2, err := s.ApplicationCommands(s.State.User.ID, config.LocalServerConfigGuildID)

	if err != nil {
		return err
	}
	for _, v := range registeredCommandsToDelete {
		err := s.ApplicationCommandDelete(s.State.User.ID, "", v.ID)
		if err != nil {
			log.Fatalf("Cannot delete '%v' command: %v", v.Name, err)
			return err
		}
	}
	for _, v := range registeredCommandsToDelete1 {
		err := s.ApplicationCommandDelete(s.State.User.ID, config.ServusDeiConfigGuildID, v.ID)
		if err != nil {
			log.Fatalf("Cannot delete '%v' command: %v", v.Name, err)
			return err
		}
	}
	for _, v := range registeredCommandsToDelete2 {
		err := s.ApplicationCommandDelete(s.State.User.ID, config.LocalServerConfigGuildID, v.ID)
		if err != nil {
			log.Fatalf("Cannot delete '%v' command: %v", v.Name, err)
			return err
		}
	}
	return nil
}

func buildReligionRoleOptionChoices() []*discordgo.ApplicationCommandOptionChoice {

	c := []*discordgo.ApplicationCommandOptionChoice{
		{
			Name:  string(config.LatinCatholic),
			Value: config.LatinCatholic,
		},
		{
			Name:  string(config.EasternCatholic),
			Value: config.EasternCatholic,
		},
		{
			Name:  string(config.OrthodoxChristian),
			Value: config.OrthodoxChristian,
		},
		{
			Name:  string(config.RCIACatechumen),
			Value: config.RCIACatechumen,
		},
		{
			Name:  string(config.Protestant),
			Value: config.Protestant,
		},
		{
			Name:  string(config.NonCatholic),
			Value: config.NonCatholic,
		},
		{
			Name:  string(config.Atheist),
			Value: config.Atheist,
		},
	}

	return c
}

func isMod(userId string) bool {
	_, ok := config.Moderator[config.ModeratorUserId(userId)]
	return ok
}

func alertNonMod(s *discordgo.Session, i *discordgo.InteractionCreate) error {
	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: NiceTryBro,
		},
	})
	if err != nil {
		return err
	}
	return nil
}
